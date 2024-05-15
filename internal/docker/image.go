package docker

import (
	"context"
	"time"

	"github.com/docker/docker/api/types/image"
)

func (m *DockerClient) FetchImages() error {
	images, err := m.client.ImageList(context.Background(), image.ListOptions{ContainerCount: true, All: true})
	m.Images = images
	return err
}

func (m *DockerClient) DeleteImage(imageID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := m.client.ImageRemove(ctx, imageID, image.RemoveOptions{PruneChildren: true})
	return err
}
