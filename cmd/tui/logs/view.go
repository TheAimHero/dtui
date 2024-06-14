package logs

import (
	"fmt"
	"os"
	"strings"

	"github.com/TheAimHero/dtui/internal/ui/table"
	ui_utils "github.com/TheAimHero/dtui/internal/ui/utils"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

var (
	physicalWidth, _, _ = term.GetSize(int(os.Stdout.Fd()))

	lineStyle = lipgloss.NewStyle().Foreground(table.HighlightColor)

	contentStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#F8F8F2"))

	titleStyle = func() lipgloss.Style {
		b := lipgloss.NormalBorder()
		b.Right = "├"
		return lipgloss.NewStyle().BorderStyle(b).BorderForeground(table.HighlightColor)
	}()
	infoStyle = func() lipgloss.Style {
		b := lipgloss.NormalBorder()
		b.Left = "┤"
		return titleStyle.Copy().BorderStyle(b).BorderForeground(table.HighlightColor)
	}()
)

func (m LogModel) View() string {
	doc := strings.Builder{}
	doc.WriteString(table.BaseTableStyle.Copy().Render(m.Table.View()))
	doc.WriteString(fmt.Sprintf("\n%s\n%s\n%s", m.headerView(), m.Viewport.View(), m.footerView()))
	doc.WriteString("\n" + m.Message.ShowMessage())
	doc.WriteString("\n" + m.Help.View(m.Keys))
	doc.WriteString(strings.Repeat("\n", ui_utils.HeightPadding(doc, 8)))
	return doc.String()
}

func (m LogModel) headerView() string {
	var title string
	if m.Title == "" {
		title = titleStyle.Render("Select Container")
	} else {
		title = titleStyle.Render("Container Log: " + m.Title)
	}
	line := lineStyle.Render(strings.Repeat("─", max(0, m.Viewport.Width-lipgloss.Width(title))))
	return lipgloss.JoinHorizontal(lipgloss.Center, title, line)
}

func (m LogModel) footerView() string {
	info := infoStyle.Render(fmt.Sprintf("%3.f%%", m.Viewport.ScrollPercent()*100))
	line := lineStyle.Render(strings.Repeat("─", max(0, m.Viewport.Width-lipgloss.Width(info))))
	return lipgloss.JoinHorizontal(lipgloss.Center, line, info)
}
