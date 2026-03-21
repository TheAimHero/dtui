package manageimage

import (
	"fmt"
	"strings"
	"time"

	ui_table "github.com/TheAimHero/dtui/internal/ui/table"
	"github.com/charmbracelet/lipgloss"
)

type ShowTextInput struct{}

const (
	successDuration = 2 * time.Second
	errorDuration   = 5 * time.Second
)

func (m ImageModel) pullImage() string {
	var lines []string
	m.PullProgress.Range(func(key, value any) bool {
		imageName := key.(string)
		info := value.(PullProgressInfo)
		line := fmt.Sprintf("%s %s", m.PullSpinner.View(), imageName)
		if info.Progress != nil && info.Progress.Total > 0 {
			percent := float64(info.Progress.Current) / float64(info.Progress.Total) * 100
			line += fmt.Sprintf(": %s (%.0f%%)", info.Status, percent)
		} else {
			line += fmt.Sprintf(": %s", info.Status)
		}
		lines = append(lines, line)
		return true
	})
	return strings.Join(lines, "\n")
}

func (m ImageModel) View() string {
	doc := strings.Builder{}
	doc.WriteString(ui_table.Centered(m.Width).Render(m.Table.View()))
	if m.Input.Focused() || m.Input.Value() != "" {
		doc.WriteString("\n" + lipgloss.NewStyle().Padding(1, 0, 0, 0).Render(m.Input.View()))
	} else {
		doc.WriteString(strings.Repeat("\n", 2))
	}
	doc.WriteString("\n" + ui_table.Centered(m.Width).Render(m.Conformation.View()))
	doc.WriteString("\n" + ui_table.Centered(m.Width).Render(m.Message.ShowMessage()))
	doc.WriteString("\n" + ui_table.Centered(m.Width).Render(m.Help.View(m.Keys)))
	count := 0
	m.PullProgress.Range(func(_, _ any) bool {
		count++
		return false
	})
	if count > 0 {
		doc.WriteString("\n" + ui_table.Centered(m.Width).Render(m.pullImage()))
	} else {
		doc.WriteString("\n")
	}
	padding := m.Height - lipgloss.Height(doc.String()) - 8
	if padding < 0 {
		padding = 0
	}
	doc.WriteString(strings.Repeat("\n", padding))
	return doc.String()
}
