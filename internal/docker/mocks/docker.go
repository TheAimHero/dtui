package mocks

import (
	"github.com/TheAimHero/dtui/internal/docker"
	"github.com/docker/docker/api/types"
)

type MockDockerClient struct {
	Containers docker.Containers
}

func NewMockDockerClient() *MockDockerClient {
	return &MockDockerClient{
		Containers: []types.Container{
			{ID: "1", Names: []string{"/container1"}, Image: "image1", Status: "running"},
			{ID: "2", Names: []string{"/container2"}, Image: "image2", Status: "running"},
			{ID: "3", Names: []string{"/container3"}, Image: "image3", Status: "running"},
		},
	}
}

func (d *MockDockerClient) FetchContainers() error {
	return nil
}

func (d *MockDockerClient) StopContainer(containerID string) error {
	return nil
}

func (d *MockDockerClient) StartContainer(containerID string) error {
	return nil
}

func (d *MockDockerClient) DeleteContainer(containerID string) error {
	for i, c := range d.Containers {
		if c.ID == containerID {
			d.Containers = append(d.Containers[:i], d.Containers[i+1:]...)
			break
		}
	}
	return nil
}

func (d *MockDockerClient) GetContainers() docker.Containers {
	return d.Containers
}
