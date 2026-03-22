package wip

import (
	tea "github.com/charmbracelet/bubbletea"
	"golang.org/x/sys/unix"
	"golang.org/x/term"
)

type WipModel struct {
	Title  string
	Width  int
	Height int
}

func (m WipModel) Init() tea.Cmd {
	return nil
}

func NewModel() WipModel {
	width, height, _ := term.GetSize(unix.Stdout)
	return WipModel{
		Title:  "Work In Progress",
		Width:  width,
		Height: height,
	}
}
