package docker

import (
	"io"
)

type ContainerService interface {
	FetchContainers() (Containers, error)
	StopContainer(id string) error
	StartContainer(id string) error
	DeleteContainer(id string) error
	GetLogs(id string) (io.ReadCloser, error)
}

type ImageService interface {
	FetchImages() (Images, error)
	DeleteImage(id string) error
	PruneImage() error
	PullImage(name string, ch chan<- PullProgressEvent) error
}

type VolumeService interface {
	FetchVolumes() (Volumes, error)
	DeleteVolume(id string, force bool) error
	PruneVolume() error
}
