package tabs

import (
	managecontainer "github.com/TheAimHero/dtui/cmd/tui/manageContainer"
	manageimage "github.com/TheAimHero/dtui/cmd/tui/manageImage"
	managevolume "github.com/TheAimHero/dtui/cmd/tui/manageVolume"
	wip "github.com/TheAimHero/dtui/cmd/tui/wip"
	"github.com/TheAimHero/dtui/internal/docker"
	"github.com/TheAimHero/dtui/internal/ui/components"
	tea "github.com/charmbracelet/bubbletea"
)

type MainModel struct {
	WipTab       wip.WipModel
	ContainerSvc docker.ContainerService
	ImageSvc     docker.ImageService
	VolumeSvc    docker.VolumeService
	TabsTitle    []string
	Tabs         []tea.Model
	ActiveTab    int
	Width        int
	Height       int
}

func (m MainModel) Init() tea.Cmd {
	cmds := []tea.Cmd{}
	for _, m := range m.Tabs {
		cmds = append(cmds, m.Init())
	}
	return tea.Batch(cmds...)
}

func NewModel(containerSvc docker.ContainerService, imageSvc docker.ImageService, volumeSvc docker.VolumeService) (tea.Model, error) {
	containerModel, err := managecontainer.NewModel(containerSvc)
	if err != nil {
		return nil, components.ErrorMessage(err.Error())
	}
	imageModel, err := manageimage.NewModel(imageSvc)
	if err != nil {
		return nil, components.ErrorMessage(err.Error())
	}
	volumeModel, err := managevolume.NewModel(volumeSvc)
	if err != nil {
		return nil, components.ErrorMessage(err.Error())
	}
	wipModel := wip.NewModel()
	model := MainModel{
		TabsTitle:    []string{"Manage Container", "Manage Images", "Manage Volumes", "Work In Progress"},
		Tabs:         []tea.Model{&containerModel, &imageModel, &volumeModel, &wipModel},
		ContainerSvc: containerSvc,
		ImageSvc:     imageSvc,
		VolumeSvc:    volumeSvc,
		ActiveTab:    0,
		Width:        80,
		Height:       40,
	}
	return model, nil
}

func NewTui() error {
	dockerClient, err := docker.NewDockerClient()
	if err != nil {
		return err
	}
	model, err := NewModel(&dockerClient, &dockerClient, &dockerClient)
	if err != nil {
		return err
	}
	_, err = tea.NewProgram(model, tea.WithAltScreen(), tea.WithMouseCellMotion()).Run()
	if err != nil {
		return err
	}
	return nil
}
