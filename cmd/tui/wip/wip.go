package wip

import (
	tea "github.com/charmbracelet/bubbletea"
)

type WipModel struct {
	Title string
}

func (m WipModel) Init() tea.Cmd {
	return nil
}

func NewModel() WipModel {
	return WipModel{
		Title: "Work In Progress",
	}
}
