package cleaner

import (
	"context"
	"fmt"
	"log"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
)

// CleanImages searches for and removes dangling (untagged/unused parent-less) Docker images.
// If dryRun is true, it only logs the images that would be removed.
func CleanImages(ctx context.Context, cli *client.Client, dryRun bool) error {
	log.Printf("[Images] Scanning for dangling images (dryRun=%v)...", dryRun)

	// Filter for dangling images
	filterArgs := filters.NewArgs()
	filterArgs.Add("dangling", "true")

	images, err := cli.ImageList(ctx, types.ImageListOptions{
		Filters: filterArgs,
	})
	if err != nil {
		return fmt.Errorf("failed to list images: %w", err)
	}

	cleanedCount := 0
	for _, img := range images {
		// Log some representation info of the image (first repo tag if available, or just ID)
		displayID := img.ID
		if len(displayID) > 19 {
			displayID = displayID[:19] // standard sha256:abcd... display prefix
		}

		if dryRun {
			log.Printf("[Images] [Dry Run] Would remove dangling image %s (Size: %.2f MB)", displayID, float64(img.Size)/(1024*1024))
		} else {
			log.Printf("[Images] Removing dangling image %s (Size: %.2f MB)...", displayID, float64(img.Size)/(1024*1024))
			removeOpts := types.ImageRemoveOptions{
				Force:         false,
				PruneChildren: true,
			}
			_, err := cli.ImageRemove(ctx, img.ID, removeOpts)
			if err != nil {
				log.Printf("[Images] Error removing image %s: %v", displayID, err)
				continue
			}
			log.Printf("[Images] Successfully removed image %s", displayID)
		}
		cleanedCount++
	}

	log.Printf("[Images] Done. Cleaned up %d images.", cleanedCount)
	return nil
}
