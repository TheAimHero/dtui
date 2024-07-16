package prompt

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	After  tea.Cmd
	Title  string
	Active bool
	Res    bool
}

func NewModel(title string, after tea.Cmd) Model {
	return Model{
		Title:  title,
		Active: true,
		Res:    false,
		After:  after,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "y", "Y":
			m.Res = true
			m.Active = false
			return m, m.After

		case "n", "N":
			m.Res = false
			m.Active = false

		case "ctrl+c", "q", "esc":
			m.Active = false
		}
	}
	return m, nil
}

func (m Model) View() string {
	if !m.Active {
		return ""
	}
	return fmt.Sprintf("%s: [Y/n]", m.Title)
}
