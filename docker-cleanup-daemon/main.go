package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/docker/docker/client"

	"docker-cleanup-daemon/cleaner"
	"docker-cleanup-daemon/config"
	"docker-cleanup-daemon/scheduler"
)

func main() {
	// Parse command line arguments
	dryRunFlag := flag.Bool("dry-run", false, "Force dry-run mode (logs what would be deleted, overrides config.yaml)")
	onceFlag := flag.Bool("once", false, "Run a single cleanup cycle and exit immediately instead of scheduling background cron")
	flag.Parse()

	log.Println("[Daemon] Starting Docker Cleanup Daemon...")

	// 1. Locate and load configuration
	configPath := "config.yaml"
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		exePath, err := os.Executable()
		if err == nil {
			configPath = filepath.Join(filepath.Dir(exePath), "config.yaml")
		}
	}

	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("[Daemon] Critical Error: failed to load config from %s: %v", configPath, err)
	}

	// Override config settings if command-line flags are specified
	if *dryRunFlag {
		cfg.DryRun = true
	}

	log.Printf("[Daemon] Configuration loaded successfully from %s", configPath)
	log.Printf("[Daemon] Schedule: %q | Dry Run: %v", cfg.Schedule, cfg.DryRun)
	log.Printf("[Daemon] Cleanup plan - Containers: %v (threshold: %s), Images: %v, Volumes: %v",
		cfg.Cleanup.StoppedContainers, cfg.Thresholds.ContainersOlderThan, cfg.Cleanup.DanglingImages, cfg.Cleanup.UnusedVolumes)

	// 2. Initialize Docker SDK client
	cli, err := client.NewClientWithOpts(
		client.FromEnv,
		client.WithAPIVersionNegotiation(),
	)
	if err != nil {
		log.Fatalf("[Daemon] Critical Error: failed to create Docker client: %v", err)
	}
	defer cli.Close()

	// 3. Define the cleanup execution task
	runCleanup := func() {
		log.Println("[Daemon] [Cleanup Triggered] Starting cleanup cycle...")
		ctx := context.Background()

		if cfg.Cleanup.StoppedContainers {
			olderThan := cfg.GetContainersOlderThanDuration()
			if err := cleaner.CleanContainers(ctx, cli, cfg.DryRun, olderThan); err != nil {
				log.Printf("[Daemon] Error cleaning containers: %v", err)
			}
		}

		if cfg.Cleanup.DanglingImages {
			if err := cleaner.CleanImages(ctx, cli, cfg.DryRun); err != nil {
				log.Printf("[Daemon] Error cleaning images: %v", err)
			}
		}

		if cfg.Cleanup.UnusedVolumes {
			if err := cleaner.CleanVolumes(ctx, cli, cfg.DryRun); err != nil {
				log.Printf("[Daemon] Error cleaning volumes: %v", err)
			}
		}
		log.Println("[Daemon] [Cleanup Finished] Cleanup cycle complete.")
	}

	// 4. Perform startup cleanup cycle
	log.Println("[Daemon] Performing initial cleanup on startup...")
	runCleanup()

	// If run-once is requested, exit immediately without starting scheduler
	if *onceFlag {
		log.Println("[Daemon] Run once completed. Exiting daemon process.")
		return
	}

	// 5. Setup scheduler
	runner := scheduler.NewJobRunner()
	_, err = runner.AddSchedule(cfg.Schedule, runCleanup)
	if err != nil {
		log.Fatalf("[Daemon] Critical Error: failed to schedule cron job: %v", err)
	}
	runner.Start()
	log.Printf("[Daemon] Scheduler running. Scheduled next runs using pattern: %s", cfg.Schedule)

	// 6. Graceful shutdown handler
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	sig := <-sigChan
	log.Printf("[Daemon] Received signal %v. Initiating graceful shutdown...", sig)

	runner.Stop()
	log.Println("[Daemon] Scheduler stopped. Exiting daemon process.")
}
