package tabs

import (
	managecontianer "github.com/TheAimHero/dtui/cmd/tui/manageContianer"
	manageimage "github.com/TheAimHero/dtui/cmd/tui/manageImage"
	managevolume "github.com/TheAimHero/dtui/cmd/tui/manageVolume"
	wip "github.com/TheAimHero/dtui/cmd/tui/wip"
	"github.com/TheAimHero/dtui/internal/docker"
	tea "github.com/charmbracelet/bubbletea"
)

type MainModel struct {
	WipTab       wip.WipModel
	DockerClient docker.DockerClient
	TabsTitle    []string
	ContainerTab managecontianer.ContainerModel
	VolumeTab    managevolume.VolumeModel
	ImageTab     manageimage.ImageModel
	ActiveTab    int
}

func (m MainModel) Init() tea.Cmd {
	return tea.Batch(m.ContainerTab.Init(), m.ImageTab.Init(), m.WipTab.Init())
}

func NewModel(dockerClient docker.DockerClient) tea.Model {
	containerModel := managecontianer.NewModel(dockerClient)
	imageModel := manageimage.NewModel(dockerClient)
	volumeModel := managevolume.NewModel(dockerClient)
	wipModel := wip.NewModel()
	model := MainModel{
		TabsTitle:    []string{"Manage Container", "Manage Images", "Manage Volumes", "Work In Progress"},
		ContainerTab: containerModel,
		ImageTab:     imageModel,
		WipTab:       wipModel,
		VolumeTab:    volumeModel,
		DockerClient: dockerClient,
		ActiveTab:    0,
	}
	return model
}

func NewTui() error {
	dockerClient, err := docker.NewDockerClient()
	if err != nil {
		return err
	}
	model := NewModel(dockerClient)
	p := tea.NewProgram(model, tea.WithAltScreen(), tea.WithMouseCellMotion())
	// for dev purpose
	// p := tea.NewProgram(model)
	_, err = p.Run()
	if err != nil {
		return err
	}
	return nil
}
