package styles

import (
	"errors"

	"github.com/charmbracelet/lipgloss"
)

func ErrorMessage(message string) error {
	return errors.New(lipgloss.NewStyle().
		Foreground(lipgloss.Color("#cb4154")).
    Render("Error: " + "Docker is not running...\nStart Docker and try again."))
}
