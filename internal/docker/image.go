package docker

import (
	"context"
	"errors"
	"io"
	"sync"
	"time"

	"github.com/docker/docker/api/types/filters"
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

func (m *DockerClient) PruneImage() error {
	_, err := m.client.ImagesPrune(context.Background(), filters.Args{})
	return err
}

func (m *DockerClient) PullImage(imageName string) (io.ReadCloser, error) {
	wg := sync.WaitGroup{}
	wg.Add(1)
	errChan := make(chan error)
	go func() {
		stream, err := m.client.ImagePull(context.Background(), imageName, image.PullOptions{})
		if err != nil {
			wg.Done()
			errChan <- errors.New("Error while pulling image: " + imageName)
			return
		}
		defer stream.Close()
		_, _ = io.Copy(io.Discard, stream)
		wg.Done()
		errChan <- nil
	}()
	wg.Wait()
	return nil, <-errChan
}
