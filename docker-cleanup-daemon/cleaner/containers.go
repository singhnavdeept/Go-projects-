package cleaner

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

// CleanContainers removes stopped containers (status exited or created)
// that are older than the specified duration. If dryRun is true, it only
// logs the containers that would be removed.
func CleanContainers(ctx context.Context, cli *client.Client, dryRun bool, olderThan time.Duration) error {
	log.Printf("[Containers] Scanning for stopped containers (olderThan=%v, dryRun=%v)...", olderThan, dryRun)

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{All: true})
	if err != nil {
		return fmt.Errorf("failed to list containers: %w", err)
	}

	cleanedCount := 0
	for _, c := range containers {
		// We only want to clean up stopped containers.
		// "exited" and "created" represent stopped/non-running containers.
		if c.State != "exited" && c.State != "created" {
			continue
		}

		containerName := "unknown"
		if len(c.Names) > 0 {
			containerName = c.Names[0]
		}

		// If a threshold is defined, inspect the container to check its stoppage or creation time.
		if olderThan > 0 {
			inspect, err := cli.ContainerInspect(ctx, c.ID)
			if err != nil {
				log.Printf("[Containers] Warning: failed to inspect container %s (%s): %v. Skipping.", c.ID[:12], containerName, err)
				continue
			}

			var referenceTime time.Time
			if inspect.State.Status == "exited" {
				finishedTime, err := time.Parse(time.RFC3339Nano, inspect.State.FinishedAt)
				if err != nil || finishedTime.IsZero() {
					// Fallback to Created time if FinishedAt is unavailable
					createdTime, err := time.Parse(time.RFC3339Nano, inspect.Created)
					if err == nil {
						referenceTime = createdTime
					} else {
						referenceTime = time.Unix(c.Created, 0)
					}
				} else {
					referenceTime = finishedTime
				}
			} else {
				// State "created" or other stopped state without a FinishedAt
				createdTime, err := time.Parse(time.RFC3339Nano, inspect.Created)
				if err == nil {
					referenceTime = createdTime
				} else {
					referenceTime = time.Unix(c.Created, 0)
				}
			}

			age := time.Since(referenceTime)
			if age < olderThan {
				log.Printf("[Containers] Container %s (%s) is only %s old (threshold is %s). Keeping.", c.ID[:12], containerName, age.Round(time.Second), olderThan)
				continue
			}
		}

		if dryRun {
			log.Printf("[Containers] [Dry Run] Would remove container %s (%s) [State: %s]", c.ID[:12], containerName, c.State)
		} else {
			log.Printf("[Containers] Removing container %s (%s) [State: %s]...", c.ID[:12], containerName, c.State)
			removeOpts := types.ContainerRemoveOptions{
				RemoveVolumes: false,
				RemoveLinks:   false,
				Force:         false,
			}
			if err := cli.ContainerRemove(ctx, c.ID, removeOpts); err != nil {
				log.Printf("[Containers] Error removing container %s: %v", c.ID[:12], err)
				continue
			}
			log.Printf("[Containers] Successfully removed container %s", c.ID[:12])
		}
		cleanedCount++
	}

	log.Printf("[Containers] Done. Cleaned up %d containers.", cleanedCount)
	return nil
}
