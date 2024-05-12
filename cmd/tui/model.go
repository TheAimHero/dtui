package tui

import (
	"github.com/TheAimHero/dtui/internal/docker"
	tea "github.com/charmbracelet/bubbletea"
)

type MainModel struct {
	DockerClient docker.DockerClient
	TabsTitle    []string
	Tabs         []tea.Model
	ActiveTab    int
}

func (m MainModel) Init() tea.Cmd {
	return nil
}

