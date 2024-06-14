package docker

import (
	"context"
	"errors"

	"github.com/charmbracelet/lipgloss"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/api/types/volume"
	"github.com/docker/docker/client"
)

type Containers []types.Container
type Images []image.Summary

// type Volumes []volume.Volume
type Volumes []*volume.Volume

type DockerClient struct {
	client     *client.Client
	Containers Containers
	Images     Images
	Volumes    Volumes
}

func NewDockerClient() (DockerClient, error) {
	client, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return DockerClient{}, err
	}
	_, err = client.Ping(context.Background())
	if err != nil {
		return DockerClient{}, errors.New(lipgloss.NewStyle().
			Foreground(lipgloss.Color("#cb4154")).
			Render("Docker is not running...\nStart Docker and try again."))
	}
	return DockerClient{client: client}, nil
}
