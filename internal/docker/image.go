package docker

import (
	"context"

	"github.com/docker/docker/api/types/image"
)

func (m *DockerClient) FetchImages() error {
	images, err := m.client.ImageList(context.Background(), image.ListOptions{ContainerCount: true, All: true})
	m.Images = images
	return err
}
