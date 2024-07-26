package docker

import (
	"context"
	"time"

	"github.com/TheAimHero/dtui/internal/ui/styles"
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
	client, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return DockerClient{}, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	_, err = client.Ping(ctx)
	if err != nil {
		return DockerClient{}, styles.ErrorMessage("Docker is not running...\nStart Docker and try again.")
	}
	return DockerClient{client: client}, nil
}
