package docker

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/pkg/jsonmessage"
)

func (m *DockerClient) FetchImages() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	images, err := m.client.ImageList(ctx, image.ListOptions{ContainerCount: true, All: true})
	m.Images = images
	if err != nil {
		return fmt.Errorf("failed to fetch images: %w", err)
	}
	return nil
}

func (m *DockerClient) DeleteImage(imageID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := m.client.ImageRemove(ctx, imageID, image.RemoveOptions{PruneChildren: true})
	if err != nil {
		return fmt.Errorf("failed to delete image %s: %w", imageID, err)
	}
	return nil
}

func (m *DockerClient) PruneImage() error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	_, err := m.client.ImagesPrune(ctx, filters.Args{})
	if err != nil {
		return fmt.Errorf("failed to prune images: %w", err)
	}
	return nil
}

type PullProgressEvent struct {
	ID       string
	Status   string
	Progress *ProgressDetail
	Error    string
}

type ProgressDetail struct {
	Current int64
	Total   int64
	Start   int64
}

type PullProgressInfo struct {
	ID       string
	Status   string
	Progress *ProgressDetail
}

func (m *DockerClient) PullImage(imageName string, progressChan chan<- PullProgressEvent) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()
	stream, err := m.client.ImagePull(ctx, imageName, image.PullOptions{})
	if err != nil {
		return fmt.Errorf("failed to pull image %s: %w", imageName, err)
	}
	defer stream.Close()

	decoder := json.NewDecoder(stream)
	for {
		var msg jsonmessage.JSONMessage
		if err := decoder.Decode(&msg); err != nil {
			if err == io.EOF {
				break
			}
			return fmt.Errorf("failed to decode image pull stream for %s: %w", imageName, err)
		}

		event := PullProgressEvent{
			ID:     msg.ID,
			Status: msg.Status,
		}
		if msg.Error != nil {
			event.Error = msg.Error.Message
		}
		if msg.Progress != nil {
			event.Progress = &ProgressDetail{
				Current: msg.Progress.Current,
				Total:   msg.Progress.Total,
				Start:   msg.Progress.Start,
			}
		}
		select {
		case progressChan <- event:
		default:
		}
	}
	return nil
}
