package tui

import (
	"os"
	"strings"

	"github.com/TheAimHero/dtui/internal/ui"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

var (
	activeTabBorder = lipgloss.Border{
		Top:         "─",
		Bottom:      " ",
		Left:        "│",
		Right:       "│",
		TopLeft:     "┌",
		TopRight:    "┐",
		BottomLeft:  "┘",
		BottomRight: "└",
	}
	tabBorder = lipgloss.Border{
		Top:         "─",
		Bottom:      "─",
		Left:        "│",
		Right:       "│",
		TopLeft:     "┌",
		TopRight:    "┐",
		BottomLeft:  "┴",
		BottomRight: "┴",
	}
	endTabBorder = lipgloss.Border{
		Top:         " ",
		Bottom:      "┐",
		Left:        " ",
		Right:       " ",
		TopLeft:     " ",
		TopRight:    " ",
		BottomLeft:  " ",
		BottomRight: " ",
	}
	inactiveTabBorder = lipgloss.NewStyle().Border(lipgloss.NormalBorder())
	docStyle          = lipgloss.NewStyle().Padding(0, 0, 0, 0)
	inactiveTabStyle  = lipgloss.NewStyle().Border(tabBorder, true).BorderForeground(ui.HighlightColor)
	padTabStyle       = lipgloss.NewStyle().Foreground(ui.HighlightColor)
	tabEndStyle       = lipgloss.NewStyle().Border(endTabBorder, false, false, true, false).BorderForeground(ui.HighlightColor)
	activeTabStyle    = inactiveTabStyle.Copy().Border(activeTabBorder, true)
	windowStyle       = lipgloss.NewStyle().BorderForeground(ui.HighlightColor).Align(lipgloss.Center).Border(lipgloss.NormalBorder()).UnsetBorderTop().Padding(2, 0)
)

func Tab(m MainModel) string {
	doc := strings.Builder{}
	physicalWidth, _, _ := term.GetSize(int(os.Stdout.Fd()))

	var renderedTabs []string

	for i := range m.Tabs {
		var style lipgloss.Style
		isFirst, isLast, isActive := i == 0, i == len(m.Tabs)-1, i == m.ActiveTab
		if isActive {
			style = activeTabStyle.Copy()
		} else {
			style = inactiveTabStyle.Copy()
		}
		border, _, _, _, _ := style.GetBorder()
		if isFirst && isActive {
			border.BottomLeft = "│"
		} else if isFirst && !isActive {
			border.BottomLeft = "├"
		} else if isLast && isActive {
			border.BottomRight = "└"
		} else if isLast && !isActive {
			border.BottomRight = "┴"
		}
		style = style.Border(border)
		renderedTabs = append(renderedTabs, style.Render(m.TabsTitle[i]))
	}
	row := lipgloss.JoinHorizontal(lipgloss.Top, renderedTabs...)
	paddingBorder := padTabStyle.Render(strings.Repeat("─", physicalWidth-lipgloss.Width(row)-1) + "┐")
	row = lipgloss.JoinHorizontal(lipgloss.Bottom, row, paddingBorder)
	doc.WriteString(row)
	doc.WriteString(windowStyle.Width((physicalWidth - windowStyle.GetHorizontalFrameSize())).Render(m.Tabs[m.ActiveTab].View()))
	return docStyle.Render(doc.String())
}
