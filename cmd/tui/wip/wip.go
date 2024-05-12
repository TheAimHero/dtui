package wip

import (
	tea "github.com/charmbracelet/bubbletea"
)

type wipModel struct {
	Title string
}

func (m wipModel) Init() tea.Cmd {
	return nil
}

func (m wipModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "enter", " ":
		}
	}

	return m, nil
}

func NewModel() tea.Model {
	return wipModel{
		Title: "Work In Progress",
	}
}
