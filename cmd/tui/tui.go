package tui

import (
	managecontianer "github.com/TheAimHero/dtui/cmd/tui/manageContianer"
	manageimage "github.com/TheAimHero/dtui/cmd/tui/manageImage"
	wip "github.com/TheAimHero/dtui/cmd/tui/wip"
	"github.com/TheAimHero/dtui/internal/docker"
	tea "github.com/charmbracelet/bubbletea"
)

func Init() error {
	dockerClient, err := docker.NewDockerClient()
	if err != nil {
		return err
	}
	containerModel := managecontianer.NewModel(dockerClient)
	imageModel := manageimage.NewModel(dockerClient)
	wipModel := wip.NewModel()
	model := MainModel{
		TabsTitle: []string{"Manage Container", "Manage Images", "Work In Progress"},
		Tabs:      []tea.Model{containerModel, imageModel, wipModel},
		ActiveTab: 0,
	}
	p := tea.NewProgram(model, tea.WithAltScreen())
	// p := tea.NewProgram(model)
	_, err = p.Run()
	if err != nil {
		return err
	}
	return nil
}
