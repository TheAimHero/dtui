package tabs

import (
	"strings"

	"github.com/TheAimHero/dtui/internal/ui/components"
	"github.com/charmbracelet/lipgloss"
)

const (
	ContainerTab = iota
	ImageTab
	VolumeTab
	WipTab
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
	docStyle         = lipgloss.NewStyle().Padding(0, 0, 0, 0)
	inactiveTabStyle = lipgloss.NewStyle().Border(tabBorder, true).BorderForeground(components.HighlightColor)
	padTabStyle      = lipgloss.NewStyle().Foreground(components.HighlightColor)
	activeTabStyle   = inactiveTabStyle.Border(activeTabBorder, true)
)

func TabView(m MainModel) string {
	return m.Tabs[m.ActiveTab].View()
}

func (m MainModel) View() string {
	doc := strings.Builder{}

	var renderedTabs []string
	for i := range m.TabsTitle {
		var style lipgloss.Style
		isFirst, isLast, isActive := i == 0, i == len(m.TabsTitle)-1, i == m.ActiveTab
		if isActive {
			style = activeTabStyle
		} else {
			style = inactiveTabStyle
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
	repeatCount := m.Width - lipgloss.Width(row)
	if repeatCount < 0 {
		repeatCount = 0
	}
	paddingBorder := padTabStyle.Render(strings.Repeat("─", repeatCount) + "┐")
	row = lipgloss.JoinHorizontal(lipgloss.Bottom, row, paddingBorder)
	doc.WriteString(row)

	doc.WriteString(TabView(m))
	return doc.String()
}
