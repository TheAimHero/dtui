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
	batch := make([]tea.Cmd, len(m.Tabs))
	for i := range m.Tabs {
		batch[i] = m.Tabs[i].Init()
	}
	return tea.Batch(batch...)
}
