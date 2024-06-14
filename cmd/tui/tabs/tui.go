package tabs

import (
	logs "github.com/TheAimHero/dtui/cmd/tui/logs"
	managecontianer "github.com/TheAimHero/dtui/cmd/tui/manageContianer"
	manageimage "github.com/TheAimHero/dtui/cmd/tui/manageImage"
	wip "github.com/TheAimHero/dtui/cmd/tui/wip"
	"github.com/TheAimHero/dtui/internal/docker"
	tea "github.com/charmbracelet/bubbletea"
)

type MainModel struct {
	WipTab       wip.WipModel
	DockerClient docker.DockerClient
	TabsTitle    []string
	LogsTab      logs.LogModel
	ContainerTab managecontianer.ContainerModel
	ImageTab     manageimage.ImageModel
	ActiveTab    int
}

func (m MainModel) Init() tea.Cmd {
	return tea.Batch(m.ContainerTab.Init(), m.ImageTab.Init(), m.LogsTab.Init(), m.WipTab.Init())
}

func NewModel(dockerClient docker.DockerClient) tea.Model {
	containerModel := managecontianer.NewModel(dockerClient)
	imageModel := manageimage.NewModel(dockerClient)
	logsModel := logs.NewModel(dockerClient)
	wipModel := wip.NewModel()
	model := MainModel{
		TabsTitle:    []string{"Manage Container", "Manage Images", "View Logs", "Work In Progress"},
		ContainerTab: containerModel,
		ImageTab:     imageModel,
		LogsTab:      logsModel,
		WipTab:       wipModel,
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
	// p := tea.NewProgram(model, tea.WithAltScreen(), tea.WithMouseCellMotion())
	// for dev purpose
	p := tea.NewProgram(model)
	_, err = p.Run()
	if err != nil {
		return err
	}
	return nil
}
