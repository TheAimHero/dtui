package managevolume

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
	Up          key.Binding
	Down        key.Binding
	Left        key.Binding
	Right       key.Binding
	Help        key.Binding
	Quit        key.Binding
	PruneVolume key.Binding
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit, k.Up, k.Down, k.Left, k.Right}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Help, k.Quit},
		{k.Up, k.Down},
		{k.Left, k.Right},
		{k.PruneVolume},
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
	PruneVolume: key.NewBinding(
		key.WithKeys("p"),
		key.WithHelp("p", "prune volume"),
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
