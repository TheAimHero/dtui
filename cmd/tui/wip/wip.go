package wip

import (
	"os"

	tea "github.com/charmbracelet/bubbletea"
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
	width, height, _ := term.GetSize(int(os.Stdout.Fd()))
	return WipModel{
		Title:  "Work In Progress",
		Width:  width,
		Height: height,
	}
}
