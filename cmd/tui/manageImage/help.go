package manageimage

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
	Left            key.Binding
	Right           key.Binding
	Help            key.Binding
	Quit            key.Binding
	DeleteImages    key.Binding
	PruneImages     key.Binding
	SelectImage     key.Binding
	SelectAllImages key.Binding
	EscapeInput     key.Binding
	ShowInput       key.Binding
	Submit          key.Binding
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit, k.Up, k.Down, k.Left, k.Right}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down},
		{k.Left, k.Right},
		{k.Help, k.Quit},
		{k.SelectImage, k.SelectAllImages},
		{k.DeleteImages, k.PruneImages},
		{k.EscapeInput, k.ShowInput, k.Submit},
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
	SelectImage: key.NewBinding(
		key.WithKeys(" "),
		key.WithHelp("space", "select image"),
	),
	SelectAllImages: key.NewBinding(
		key.WithKeys("V"),
		key.WithHelp("V", "select all images"),
	),
	DeleteImages: key.NewBinding(
		key.WithKeys("d"),
		key.WithHelp("d", "delete image"),
	),
	EscapeInput: key.NewBinding(
		key.WithKeys("esc", "ctrl+c"),
		key.WithHelp("esc/ctrl+c", "escape input"),
		key.WithDisabled(),
	),
	PruneImages: key.NewBinding(
		key.WithKeys("p"),
		key.WithHelp("p", "prune images"),
	),
	ShowInput: key.NewBinding(
		key.WithKeys("/"),
		key.WithHelp("/", "show input"),
	),
	Submit: key.NewBinding(
		key.WithKeys("enter"),
		key.WithDisabled(),
		key.WithHelp("enter", "submit"),
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
