package managevolume

import (
	"github.com/TheAimHero/dtui/internal/ui/components"
	"github.com/charmbracelet/bubbles/key"
)

type keyMap struct {
	Up           key.Binding
	Down         key.Binding
	Left         key.Binding
	Right        key.Binding
	Help         key.Binding
	Quit         key.Binding
	ShowInput    key.Binding
	EscapeInput  key.Binding
	SetFilter    key.Binding
	PruneVolume  key.Binding
	DeleteVolume key.Binding
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit, k.Up, k.Down, k.Left, k.Right}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Help, k.Quit},
		{k.Up, k.Down},
		{k.Left, k.Right},
		{k.ShowInput, k.EscapeInput, k.SetFilter},
		{k.PruneVolume, k.DeleteVolume},
	}
}

var keys = keyMap{
	Up:    components.NewNavigationKeys().Up,
	Down:  components.NewNavigationKeys().Down,
	Left:  components.NewNavigationKeys().Left,
	Right: components.NewNavigationKeys().Right,
	Help:  components.NewNavigationKeys().Help,
	Quit:  components.NewNavigationKeys().Quit,
	ShowInput: key.NewBinding(
		key.WithKeys("/"),
		key.WithHelp("/", "filter"),
	),
	EscapeInput: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "clear filter"),
		key.WithDisabled(),
	),
	SetFilter: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "apply filter"),
		key.WithDisabled(),
	),
	PruneVolume: key.NewBinding(
		key.WithKeys("p"),
		key.WithHelp("p", "prune volume"),
	),
	DeleteVolume: key.NewBinding(
		key.WithKeys("d"),
		key.WithHelp("d", "delete volume"),
	),
}
