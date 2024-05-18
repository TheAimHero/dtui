package tabs

import (
	logs "github.com/TheAimHero/dtui/cmd/tui/logs"
	managecontianer "github.com/TheAimHero/dtui/cmd/tui/manageContianer"
	manageimage "github.com/TheAimHero/dtui/cmd/tui/manageImage"
	wip "github.com/TheAimHero/dtui/cmd/tui/wip"
	"github.com/TheAimHero/dtui/internal/docker"
	tea "github.com/charmbracelet/bubbletea"
	mapset "github.com/deckarep/golang-set/v2"
)

type MainModel struct {
	InitTab      mapset.Set[int]
	DockerClient docker.DockerClient
	TabsTitle    []string
	Tabs         []tea.Model
	ActiveTab    int
}

func (m MainModel) Init() tea.Cmd {
	cmds := []tea.Cmd{}
	cmds = append(cmds, m.Tabs[m.ActiveTab].Init())
	return tea.Batch(cmds...)
}

func NewModel(dockerClient docker.DockerClient) tea.Model {
	containerModel := managecontianer.NewModel(dockerClient)
	imageModel := manageimage.NewModel(dockerClient)
	logsModel := logs.NewModel(dockerClient)
	wipModel := wip.NewModel()
	model := MainModel{
		TabsTitle: []string{"Manage Container", "Manage Images", "View Logs", "Work In Progress"},
		Tabs:      []tea.Model{containerModel, imageModel, logsModel, wipModel},
		ActiveTab: 0,
		InitTab:   mapset.NewSet[int](),
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
