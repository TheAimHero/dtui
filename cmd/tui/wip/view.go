package wip

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func (m wipModel) View() string {
	doc := strings.Builder{}
	baseStyle := lipgloss.NewStyle().Padding(1, 2)
	emphasisStyle := baseStyle.Copy().Bold(true).Foreground(lipgloss.Color("#F1FA8C"))
	titleStyle := baseStyle.Copy().Bold(true).Italic(true).Foreground(lipgloss.Color("#05C3DD"))

	doc.WriteString(titleStyle.Render("Work in progress... and Coming Soon..."))
	doc.WriteString("\n" + "What's next?\n")
	doc.WriteString(emphasisStyle.Render("-\tBuild Mode"))
	doc.WriteString(emphasisStyle.Render("-\tImage Mode"))
	doc.WriteString(emphasisStyle.Render("-\tLog Mode"))
	doc.WriteString(titleStyle.UnsetPadding().Render("\nGive a Star on GitHub\n"))
	doc.WriteString(titleStyle.Render("Thats all folks!"))
	return doc.String()
}
