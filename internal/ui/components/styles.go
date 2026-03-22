package components

import (
	"errors"

	"github.com/charmbracelet/lipgloss"
)

func ErrorMessage(message string) error {
	return errors.New(lipgloss.NewStyle().
		Foreground(lipgloss.Color("#cb4154")).
		Render("Error: " + message))
}

type HelpStyles struct {
	DescStyle     lipgloss.Style
	EllipsisStyle lipgloss.Style
	KeyStyle      lipgloss.Style
}

func DefaultHelpStyles() HelpStyles {
	return HelpStyles{
		DescStyle:     lipgloss.NewStyle().Foreground(lipgloss.Color("#F1FA8C")),
		EllipsisStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("#6272a4")),
		KeyStyle:      lipgloss.NewStyle().Foreground(lipgloss.Color("#BD93F9")),
	}
}

var (
	DescStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("#F1FA8C"))
	EllipsisStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#6272a4"))
	KeyStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("#BD93F9"))
)
