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


func NewModel() tea.Model {
	return wipModel{
		Title: "Work In Progress",
	}
}
