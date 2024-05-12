package docker

import (
	"context"
	"time"

	"github.com/docker/docker/api/types/container"
)

func (m *DockerClient) FetchContainers() error {
	containers, err := m.client.ContainerList(context.Background(), container.ListOptions{All: true})
	m.Containers = containers
	return err
}

func (m *DockerClient) StopContainer(containerID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := m.client.ContainerStop(ctx, containerID, container.StopOptions{})
	return err
}

func (m *DockerClient) StartContainer(containerID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := m.client.ContainerStart(ctx, containerID, container.StartOptions{})
	return err
}

func (m *DockerClient) StartContainers(containerIDs []string) []string {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	errors := make([]string, 0)
	defer cancel()
	for _, containerID := range containerIDs {
		err := m.client.ContainerStart(ctx, containerID, container.StartOptions{})
		if err != nil {
			errors = append(errors, err.Error())
		}
	}
	return errors
}
