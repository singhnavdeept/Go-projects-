package cleaner

import (
	"context"
	"fmt"
	"log"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/volume"
	"github.com/docker/docker/client"
)

// CleanVolumes removes unused local volumes.
// If dryRun is true, it simulates the prune by checking which volumes
// are not attached to any container (running or stopped).
func CleanVolumes(ctx context.Context, cli *client.Client, dryRun bool) error {
	log.Printf("[Volumes] Scanning for unused volumes (dryRun=%v)...", dryRun)

	if dryRun {
		// Fetch all containers (running and stopped) to see what volumes are in use
		containers, err := cli.ContainerList(ctx, types.ContainerListOptions{All: true})
		if err != nil {
			return fmt.Errorf("failed to list containers for volume dry-run check: %w", err)
		}

		usedVolumes := make(map[string]bool)
		for _, c := range containers {
			for _, m := range c.Mounts {
				if m.Type == "volume" && m.Name != "" {
					usedVolumes[m.Name] = true
				}
			}
		}

		// Fetch all volumes
		volList, err := cli.VolumeList(ctx, volume.ListOptions{})
		if err != nil {
			return fmt.Errorf("failed to list volumes: %w", err)
		}

		simulatedCleanCount := 0
		for _, v := range volList.Volumes {
			if !usedVolumes[v.Name] {
				log.Printf("[Volumes] [Dry Run] Would prune unused volume: %s", v.Name)
				simulatedCleanCount++
			}
		}
		log.Printf("[Volumes] Done. [Dry Run] Would clean up %d volumes.", simulatedCleanCount)
		return nil
	}

	// Active Prune
	pruneReport, err := cli.VolumesPrune(ctx, filters.NewArgs())
	if err != nil {
		return fmt.Errorf("failed to prune volumes: %w", err)
	}

	cleanedCount := len(pruneReport.VolumesDeleted)
	for _, volName := range pruneReport.VolumesDeleted {
		log.Printf("[Volumes] Pruned unused volume: %s", volName)
	}

	log.Printf("[Volumes] Done. Cleaned up %d volumes. Space reclaimed: %.2f MB",
		cleanedCount, float64(pruneReport.SpaceReclaimed)/(1024*1024))
	return nil
}
