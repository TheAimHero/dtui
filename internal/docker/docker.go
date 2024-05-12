package docker

import (
	"context"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
)

type Containers []types.Container
type Images []image.Summary

type DockerClient struct {
	client     *client.Client
	Containers Containers
	Images     Images
}

func NewDockerClient() (DockerClient, error) {
	client, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return DockerClient{}, err
	}
	return DockerClient{client: client}, nil
}

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

func (m *DockerClient) FetchImages() error {
	images, err := m.client.ImageList(context.Background(), types.ImageListOptions{})
	m.Images = images
	return err
}
