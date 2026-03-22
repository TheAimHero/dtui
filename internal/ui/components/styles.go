package components

import (
	"github.com/charmbracelet/lipgloss"
)

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
