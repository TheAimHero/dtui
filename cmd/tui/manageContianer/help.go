package managecontianer

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/lipgloss"
)

var (
	descStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("#F1FA8C"))
	ellipsisStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#6272a4"))
	keyStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("#BD93F9"))
)

type keyMap struct {
	Up              key.Binding
	Down            key.Binding
	Help            key.Binding
	Quit            key.Binding
	StartContainer  key.Binding
	StopContainer   key.Binding
	StartContainers key.Binding
	StopContainers  key.Binding
	ToggleSelected  key.Binding
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit, k.Up, k.Down}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down},
		{k.Help, k.Quit},
		{k.StartContainer, k.StopContainer},
		{k.StartContainers, k.StopContainers},
		{k.ToggleSelected},
	}
}

var keys = keyMap{
	Up: key.NewBinding(
		key.WithKeys("up", "k"),
		key.WithHelp("↑/k", "move up"),
	),
	Down: key.NewBinding(
		key.WithKeys("down", "j"),
		key.WithHelp("↓/j", "move down"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "toggle help"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "esc", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
	StartContainer: key.NewBinding(
		key.WithKeys("s"),
		key.WithHelp("s", "start container"),
	),
	StopContainer: key.NewBinding(
		key.WithKeys("p"),
		key.WithHelp("p", "stop container"),
	),
	StartContainers: key.NewBinding(
		key.WithKeys("S"),
		key.WithHelp("S", "start selected containers"),
	),
	StopContainers: key.NewBinding(
		key.WithKeys("P"),
		key.WithHelp("P", "stop selected containers"),
	),
	ToggleSelected: key.NewBinding(
		key.WithKeys(" "),
		key.WithHelp("space", "toggle selected"),
	),
}

func getHelpSection() help.Model {
	m := help.New()
	s := m.Styles
	s.ShortDesc = descStyle
	s.FullDesc = descStyle
	s.FullKey = keyStyle
	s.ShortKey = keyStyle
	s.Ellipsis = ellipsisStyle
	s.FullSeparator = ellipsisStyle
	s.ShortSeparator = ellipsisStyle
	m.Styles = s
	return m
}
