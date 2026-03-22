package docker

import (
	"context"
	"fmt"
	"time"

	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/volume"
)

func (m *DockerClient) FetchVolumes() (Volumes, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	volumes, err := m.client.VolumeList(ctx, volume.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to fetch volumes: %w", err)
	}
	result := make(Volumes, len(volumes.Volumes))
	for i, v := range volumes.Volumes {
		result[i] = &Volume{
			Name:       v.Name,
			CreatedAt:  v.CreatedAt,
			Mountpoint: v.Mountpoint,
		}
		if v.UsageData != nil {
			result[i].UsageData = &VolumeUsageData{
				Size: v.UsageData.Size,
			}
		}
	}
	return result, nil
}

func (m *DockerClient) PruneVolume() error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	_, err := m.client.VolumesPrune(ctx, filters.NewArgs())
	if err != nil {
		return fmt.Errorf("failed to prune volumes: %w", err)
	}
	return nil
}

func (m *DockerClient) DeleteVolume(volumeID string, force bool) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := m.client.VolumeRemove(ctx, volumeID, force)
	if err != nil {
		return fmt.Errorf("failed to delete volume %s: %w", volumeID, err)
	}
	return nil
}
