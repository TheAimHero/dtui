package tabs

import (
	"os"
	"strings"

	ui_table "github.com/TheAimHero/dtui/internal/ui/table"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

const (
	ContainerTab = iota
	ImageTab
	VolumeTab
	LogsTab
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
	docStyle                         = lipgloss.NewStyle().Padding(0, 0, 0, 0)
	inactiveTabStyle                 = lipgloss.NewStyle().Border(tabBorder, true).BorderForeground(ui_table.HighlightColor)
	padTabStyle                      = lipgloss.NewStyle().Foreground(ui_table.HighlightColor)
	activeTabStyle                   = inactiveTabStyle.Copy().Border(activeTabBorder, true)
	windowStyle                      = lipgloss.NewStyle().BorderForeground(ui_table.HighlightColor).Align(lipgloss.Center).Border(lipgloss.NormalBorder()).UnsetBorderTop().Padding(2, 0)
	physicalWidth, physicalHeight, _ = term.GetSize(int(os.Stdout.Fd())) // nolint:unused
)

func TabView(m MainModel) string {
	switch m.ActiveTab {
	case ContainerTab:
		return m.ContainerTab.View()
	case ImageTab:
		return m.ImageTab.View()
	case VolumeTab:
		return m.VolumeTab.View()
	case LogsTab:
		return m.LogsTab.View()
	case WipTab:
		return m.WipTab.View()
	default:
		return ""
	}
}

func (m MainModel) View() string {
	doc := strings.Builder{}

	var renderedTabs []string
	for i := range m.TabsTitle {
		var style lipgloss.Style
		isFirst, isLast, isActive := i == 0, i == len(m.TabsTitle)-1, i == m.ActiveTab
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
	repeatCount := physicalWidth - lipgloss.Width(row) - 1
	if repeatCount < 0 {
		repeatCount = 0
	}
	paddingBorder := padTabStyle.Render(strings.Repeat("─", repeatCount) + "┐")
	row = lipgloss.JoinHorizontal(lipgloss.Bottom, row, paddingBorder)
	doc.WriteString(row)

	doc.WriteString(windowStyle.Width((physicalWidth - windowStyle.GetHorizontalFrameSize())).Render(TabView(m)))
	return docStyle.Render(doc.String())
}
