package tui

import (
	"os"
	"strings"

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
	docStyle          = lipgloss.NewStyle().Padding(2, 0, 1, 0)
	highlightColor    = lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"}
	inactiveTabStyle  = lipgloss.NewStyle().Border(tabBorder, true).BorderForeground(highlightColor)
	padTabStyle       = inactiveTabBorder.Copy().BorderTop(false).BorderRight(false).BorderLeft(false).BorderForeground(highlightColor)
	tabEndStyle       = lipgloss.NewStyle().Border(endTabBorder, false, false, true, false).BorderForeground(highlightColor)
	activeTabStyle    = inactiveTabStyle.Copy().Border(activeTabBorder, true)
	windowStyle       = lipgloss.NewStyle().BorderForeground(highlightColor).Align(lipgloss.Center).Border(lipgloss.NormalBorder()).UnsetBorderTop().Padding(2, 0)
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
	paddingBorder := padTabStyle.Render(strings.Repeat(" ", physicalWidth-lipgloss.Width(row)-1))
	row = lipgloss.JoinHorizontal(lipgloss.Bottom, row, paddingBorder)
	row = lipgloss.JoinHorizontal(lipgloss.Bottom, row, tabEndStyle.Render(" "))
	doc.WriteString(row)
	doc.WriteString(windowStyle.Width((physicalWidth - windowStyle.GetHorizontalFrameSize())).Render(m.Tabs[m.ActiveTab].View()))
	return docStyle.Render(doc.String())
}
