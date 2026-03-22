package managecontainer

import (
	"github.com/TheAimHero/dtui/internal/ui/components"
	"github.com/charmbracelet/bubbles/key"
)

type keyMap struct {
	Up               key.Binding
	Down             key.Binding
	Left             key.Binding
	Right            key.Binding
	Help             key.Binding
	Quit             key.Binding
	StartContainers  key.Binding
	StopContainers   key.Binding
	DeleteContainers key.Binding
	ToggleSelected   key.Binding
	ToggleSelectAll  key.Binding
	ShowInput        key.Binding
	EscapeInput      key.Binding
	SetFilter        key.Binding
	ShowLogs         key.Binding
	ExecContainer    key.Binding
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit, k.Up, k.Down, k.Left, k.Right}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Help, k.Quit},
		{k.Up, k.Down},
		{k.Left, k.Right},
		{k.StartContainers, k.StopContainers},
		{k.DeleteContainers},
		{k.ToggleSelected, k.ToggleSelectAll},
		{k.ShowInput, k.EscapeInput, k.ShowLogs, k.SetFilter},
		{k.ExecContainer},
	}
}

var keys = keyMap{
	Up:    components.NewNavigationKeys().Up,
	Down:  components.NewNavigationKeys().Down,
	Left:  components.NewNavigationKeys().Left,
	Right: components.NewNavigationKeys().Right,
	Help:  components.NewNavigationKeys().Help,
	Quit:  components.NewNavigationKeys().Quit,
	StartContainers: key.NewBinding(
		key.WithKeys("s"),
		key.WithHelp("s", "start container/s"),
	),
	StopContainers: key.NewBinding(
		key.WithKeys("p"),
		key.WithHelp("p", "stop container/s"),
	),
	DeleteContainers: key.NewBinding(
		key.WithKeys("d"),
		key.WithHelp("d", "delete container/s"),
	),
	ToggleSelected: key.NewBinding(
		key.WithKeys(" "),
		key.WithHelp("space", "toggle selected"),
	),
	ToggleSelectAll: key.NewBinding(
		key.WithKeys("V"),
		key.WithHelp("V", "toggle select all"),
	),
	ShowInput: key.NewBinding(
		key.WithKeys("/"),
		key.WithHelp("/", "show input"),
	),
	EscapeInput: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "escape input"),
		key.WithDisabled(),
	),
	SetFilter: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "set filter"),
		key.WithDisabled(),
	),
	ShowLogs: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "show logs"),
	),
	ExecContainer: key.NewBinding(
		key.WithKeys("e"),
		key.WithHelp("e", "exec into container"),
	),
}
