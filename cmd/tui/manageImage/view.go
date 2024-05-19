package manageimage

import (
	"fmt"
	"os"
	"strings"
	"time"

	"golang.org/x/term"

	ui_table "github.com/TheAimHero/dtui/internal/ui/table"
	ui_utils "github.com/TheAimHero/dtui/internal/ui/utils"
	"github.com/charmbracelet/lipgloss"
)

type ShowTextInput struct{}

var (
	physicalWidth, physicalHeight, _ = term.GetSize(int(os.Stdout.Fd())) // nolint:unused
	successDuration                  = 2 * time.Second
	errorDuration                    = 5 * time.Second

	lineStyle = lipgloss.NewStyle().Foreground(ui_table.HighlightColor)

	contentStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#F8F8F2"))

	titleStyle = func() lipgloss.Style {
		b := lipgloss.NormalBorder()
		b.Right = "├"
		return lipgloss.NewStyle().BorderStyle(b).BorderForeground(ui_table.HighlightColor)
	}()

	infoStyle = func() lipgloss.Style {
		b := lipgloss.NormalBorder()
		b.Left = "┤"
		return titleStyle.Copy().BorderStyle(b).BorderForeground(ui_table.HighlightColor)
	}()
)

func (m ImageModel) View() string {
	doc := strings.Builder{}
	doc.WriteString(ui_table.BaseTableStyle.Render(m.Table.View()))
	if m.Input.Focused() {
		doc.WriteString("\n" + lipgloss.NewStyle().Padding(1, 0, 0, 0).Render(m.Input.View()))
		doc.WriteString(strings.Repeat("\n", 14))
	} else {
		doc.WriteString(fmt.Sprintf("\n%s\n%s\n%s", m.headerView(), m.viewport.View(), m.footerView()))
	}
	doc.WriteString("\n" + m.message.ShowMessage())
	doc.WriteString("\n" + m.help.View(m.keys))
	doc.WriteString(strings.Repeat("\n", ui_utils.HeightPadding(doc, 8)))
	return doc.String()
}

func (m ImageModel) headerView() string {
	title := titleStyle.Render("Pull Logs")
	line := lineStyle.Render(strings.Repeat("─", max(0, m.viewport.Width-lipgloss.Width(title))))
	return lipgloss.JoinHorizontal(lipgloss.Center, title, line)
}

func (m ImageModel) footerView() string {
	info := infoStyle.Render(fmt.Sprintf("%3.f%%", m.viewport.ScrollPercent()*100))
	line := lineStyle.Render(strings.Repeat("─", max(0, m.viewport.Width-lipgloss.Width(info))))
	return lipgloss.JoinHorizontal(lipgloss.Center, line, info)
}
