package docker

import (
	"context"
	"errors"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/api/types/volume"
	"github.com/docker/docker/client"
)

type DockerInterface interface {
  // container actions
	FetchContainers() error
	StopContainer(containerID string) error
	StartContainer(containerID string) error
	DeleteContainer(containerID string) error

  // get the docker client from the docker interface
  GetDockerClient() *DockerClient
}

type Containers []types.Container
type Images []image.Summary
type Volumes []*volume.Volume

type DockerClient struct {
	client     *client.Client
	Containers Containers
	Images     Images
	Volumes    Volumes
}

func (d *DockerClient) GetDockerClient() *DockerClient {
	return d
}

func NewDockerClient() (DockerClient, error) {
	client, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return DockerClient{}, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	_, err = client.Ping(ctx)
	if err != nil {
		return DockerClient{}, errors.New(lipgloss.NewStyle().
			Foreground(lipgloss.Color("#cb4154")).
			Render("Docker is not running...\nStart Docker and try again."))
	}
	return DockerClient{client: client}, nil
}
