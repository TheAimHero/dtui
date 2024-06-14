package docker

import (
	"context"

	"github.com/docker/docker/api/types/volume"
)

func (m *DockerClient) FetchVolumes() error {
	volumes, err := m.client.VolumeList(context.Background(), volume.ListOptions{})
	m.Volumes = volumes.Volumes
	// for _, volume := range volumes.Volumes {
	// 	m.Volumes = append(m.Volumes, *volume)
	// }
	return err
}
