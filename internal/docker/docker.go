package docker

import (
	"context"
	"fmt"
	"time"

	"github.com/docker/docker/client"
)

type DockerClient struct {
	client *client.Client
}

func NewDockerClient() (DockerClient, error) {
	c, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return DockerClient{}, fmt.Errorf("failed to create docker client: %w", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	_, err = c.Ping(ctx)
	if err != nil {
		return DockerClient{}, fmt.Errorf("docker connection failed: %w", err)
	}
	return DockerClient{client: c}, nil
}
