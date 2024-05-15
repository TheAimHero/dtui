package logs

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/lipgloss"
)

var (
	descStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("#F1FA8C")).Align(lipgloss.Left, lipgloss.Center).Padding(0, 2)
	ellipsisStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#6272a4"))
	keyStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("#BD93F9")).Align(lipgloss.Left, lipgloss.Center)
)

type keyMap struct {
	Up       key.Binding
	Down     key.Binding
	Help     key.Binding
	Quit     key.Binding
	LogsDown key.Binding
	LogsUp   key.Binding
	Select   key.Binding
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit, k.Up, k.Down}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down},
		{k.Help, k.Quit},
		{k.LogsUp, k.LogsDown},
		{k.Select},
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
	LogsUp: key.NewBinding(
		key.WithKeys("c-u"),
		key.WithHelp("c-u", "scroll up"),
	),
	LogsDown: key.NewBinding(
		key.WithKeys("c-d"),
		key.WithHelp("c-d", "scroll down"),
	),
	Select: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "select log"),
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