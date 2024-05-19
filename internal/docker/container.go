package docker

import (
	"context"
	"io"
	"time"

	"github.com/docker/docker/api/types/container"
)

func (m *DockerClient) FetchContainers() error {
	containers, err := m.client.ContainerList(context.Background(), container.ListOptions{All: true})
	m.Containers = containers
	return err
}

func (m *DockerClient) StopContainer(containerID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err := m.client.ContainerStop(ctx, containerID, container.StopOptions{})
	return err
}

func (m *DockerClient) StartContainer(containerID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err := m.client.ContainerStart(ctx, containerID, container.StartOptions{})
	return err
}

func (m *DockerClient) GetLogs(containerID string) (io.ReadCloser, error) {
	stream, err := m.client.ContainerLogs(context.Background(), containerID, container.LogsOptions{Follow: true, ShowStdout: true, Details: true, ShowStderr: true, Timestamps: true})
	if err != nil {
		return nil, err
	}
	return stream, nil
}

func (m *DockerClient) DeleteContainer(containerID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err := m.client.ContainerRemove(ctx, containerID, container.RemoveOptions{Force: true})
	return err
}
