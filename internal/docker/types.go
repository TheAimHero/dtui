package docker

import (
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/api/types/volume"
)

type Container struct {
	ID     string
	Names  []string
	Image  string
	Status string
	State  string
}

func ContainerFromAPI(c types.Container) Container {
	return Container{
		ID:     c.ID,
		Names:  c.Names,
		Image:  c.Image,
		Status: c.Status,
		State:  c.State,
	}
}

type Containers []Container

type Image struct {
	ID       string
	RepoTags []string
	Created  int64
	Size     int64
}

func ImageFromAPI(i image.Summary) Image {
	tags := i.RepoTags
	if tags == nil {
		tags = []string{}
	}
	return Image{
		ID:       i.ID,
		RepoTags: tags,
		Created:  i.Created,
		Size:     i.Size,
	}
}

type Images []Image

type Volume struct {
	Name       string
	CreatedAt  string
	Mountpoint string
	UsageData  *VolumeUsageData
}

type VolumeUsageData struct {
	Size int64
}

func VolumeFromAPI(v *volume.Volume) Volume {
	vol := Volume{
		Name:       v.Name,
		CreatedAt:  v.CreatedAt,
		Mountpoint: v.Mountpoint,
	}
	if v.UsageData != nil {
		vol.UsageData = &VolumeUsageData{
			Size: v.UsageData.Size,
		}
	}
	return vol
}

type Volumes []*Volume
