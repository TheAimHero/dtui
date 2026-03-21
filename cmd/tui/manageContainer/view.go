package managecontainer

import (
	"strings"

	ui_table "github.com/TheAimHero/dtui/internal/ui/table"
	"github.com/charmbracelet/lipgloss"
)

func (m ContainerModel) View() string {
	doc := strings.Builder{}
	doc.WriteString(ui_table.Centered(m.Width).Render(m.Table.View()))
	doc.WriteString("\n" + ui_table.Centered(m.Width).Render(m.Confirmation.View()))
	if m.Input.Focused() || m.Input.Value() != "" {
		doc.WriteString("\n" + lipgloss.NewStyle().Padding(1, 0, 0, 0).Render(m.Input.View()))
	} else {
		doc.WriteString(strings.Repeat("\n", 2))
	}
	doc.WriteString("\n" + ui_table.Centered(m.Width).Render(m.Message.ShowMessage()))
	doc.WriteString("\n" + ui_table.Centered(m.Width).Render(m.Help.View(m.Keys)))
	padding := m.Height - lipgloss.Height(doc.String()) - 8
	if padding < 0 {
		padding = 0
	}
	doc.WriteString(strings.Repeat("\n", padding))
	return doc.String()
}
