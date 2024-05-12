package docker

import (
	"github.com/docker/docker/api/types"
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
