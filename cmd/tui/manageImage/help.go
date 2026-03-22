package manageimage

import (
	"github.com/TheAimHero/dtui/internal/ui/components"
	"github.com/charmbracelet/bubbles/key"
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
	Up:    components.NewNavigationKeys().Up,
	Down:  components.NewNavigationKeys().Down,
	Left:  components.NewNavigationKeys().Left,
	Right: components.NewNavigationKeys().Right,
	Help:  components.NewNavigationKeys().Help,
	Quit:  components.NewNavigationKeys().Quit,
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
