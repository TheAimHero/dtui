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
	Up               key.Binding
	Down             key.Binding
	Left             key.Binding
	Right            key.Binding
	Help             key.Binding
	Quit             key.Binding
	StartContainer   key.Binding
	StopContainer    key.Binding
	StartContainers  key.Binding
	StopContainers   key.Binding
	DeleteContainer  key.Binding
	DeleteContainers key.Binding
	ToggleSelected   key.Binding
	ToggleSelectAll  key.Binding
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit, k.Up, k.Down, k.Left, k.Right}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Help, k.Quit},
		{k.Up, k.Down},
		{k.Left, k.Right},
		{k.StartContainer, k.StopContainer},
		{k.StartContainers, k.StopContainers},
		{k.DeleteContainer, k.DeleteContainers},
		{k.ToggleSelected, k.ToggleSelectAll},
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
	Left: key.NewBinding(
		key.WithKeys("left", "h"),
		key.WithHelp("←/h", "move left"),
	),
	Right: key.NewBinding(
		key.WithKeys("right", "l"),
		key.WithHelp("→/l", "move right"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "toggle help"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "ctrl+c"),
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
	DeleteContainer: key.NewBinding(
		key.WithKeys("d"),
		key.WithHelp("d", "delete container"),
	),
	DeleteContainers: key.NewBinding(
		key.WithKeys("D"),
		key.WithHelp("D", "delete selected containers"),
	),
	ToggleSelected: key.NewBinding(
		key.WithKeys(" "),
		key.WithHelp("space", "toggle selected"),
	),
	ToggleSelectAll: key.NewBinding(
		key.WithKeys("V"),
		key.WithHelp("V", "toggle select all"),
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
