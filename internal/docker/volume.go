package docker

import (
	"context"

	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/volume"
)

func (m *DockerClient) FetchVolumes() error {
	volumes, err := m.client.VolumeList(context.Background(), volume.ListOptions{})
	m.Volumes = volumes.Volumes
	return err
}

func (m *DockerClient) PruneVolume() error {
	_, err := m.client.VolumesPrune(context.Background(), filters.NewArgs())
	return err
}

func (m *DockerClient) DeleteVolume(volumeID string, force bool) error {
	err := m.client.VolumeRemove(context.Background(), volumeID, force)
	return err
}
