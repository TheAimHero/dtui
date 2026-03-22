package components

import (
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
)

var (
	HighlightColor = lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"}
)

func Centered(width int) lipgloss.Style {
	return lipgloss.NewStyle().
		Width(width).
		Align(lipgloss.Center).
		Padding(1, 1)
}

const (
	DefaultMinTableHeight = 5
	DefaultMaxTableRatio  = 0.5
)

type TableConfig struct {
	MinHeight    int
	MaxHeightRat float64
	HeightOffset int
}

func DefaultTableConfig() TableConfig {
	return TableConfig{
		MinHeight:    DefaultMinTableHeight,
		MaxHeightRat: DefaultMaxTableRatio,
		HeightOffset: 10,
	}
}

func CalculateTableHeight(totalHeight int, config TableConfig) int {
	maxTableHeight := int(float64(totalHeight) * config.MaxHeightRat)
	tableHeight := totalHeight - config.HeightOffset

	if tableHeight < config.MinHeight {
		tableHeight = config.MinHeight
	}
	if tableHeight > maxTableHeight {
		tableHeight = maxTableHeight
	}

	return tableHeight
}

func ApplyTableStyles(t table.Model) table.Model {
	s := table.DefaultStyles()
	s.Header = s.Header.Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)
	return t
}

func NewStandardTable(columns []table.Column, rows []table.Row) table.Model {
	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
	)
	t.KeyMap.HalfPageDown.Unbind()
	t.KeyMap.HalfPageUp.Unbind()
	t.KeyMap.GotoBottom.Unbind()
	t.KeyMap.GotoTop.Unbind()
	t.KeyMap.PageDown.Unbind()
	t.KeyMap.PageUp.Unbind()
	return ApplyTableStyles(t)
}

func (c TableConfig) WithMinHeight(h int) TableConfig {
	c.MinHeight = h
	return c
}

func (c TableConfig) WithMaxRatio(r float64) TableConfig {
	c.MaxHeightRat = r
	return c
}

func (c TableConfig) WithOffset(o int) TableConfig {
	c.HeightOffset = o
	return c
}
