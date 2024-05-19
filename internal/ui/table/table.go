package table

import (
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
)

var (
	HighlightColor = lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"}
	BaseTableStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(HighlightColor).
			Padding(1, 1)
)

func NewTable(tableColumns []table.Column, tableRows []table.Row) table.Model {
	t := table.New(
		table.WithColumns(tableColumns),
		table.WithRows(tableRows),
		table.WithFocused(true),
		table.WithHeight(10),
	)
	t.KeyMap.HalfPageDown.Unbind()
	t.KeyMap.HalfPageUp.Unbind()
	t.KeyMap.GotoBottom.Unbind()
	t.KeyMap.GotoTop.Unbind()
	t.KeyMap.PageDown.Unbind()
	t.KeyMap.PageUp.Unbind()

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	return t
}
