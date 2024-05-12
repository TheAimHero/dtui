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

func (m *DockerClient) DeleteImage() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := m.client.ImageRemove(ctx, m.Images[0].ID, image.RemoveOptions{PruneChildren: true})
	return err
}

func (m *DockerClient) DeleteImages(selectedImagesIDs []string) []string {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	errors := []string{}
	for _, imageID := range selectedImagesIDs {
		_, err := m.client.ImageRemove(ctx, imageID, image.RemoveOptions{PruneChildren: true})
		if err != nil {
			errors = append(errors, err.Error())
		}
	}
	return errors
}
