package components

import (
	"os"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

var (
	_, physicalHeight, _ = term.GetSize(int(os.Stdout.Fd()))
)

func HeightPadding(doc strings.Builder, fixHeight int) int {
	paddingHeight := physicalHeight - lipgloss.Height(doc.String()) - fixHeight
	if paddingHeight < 0 {
		paddingHeight = 0
	}
	return paddingHeight
}

type ViewBuilder struct {
	width   int
	height  int
	doc     strings.Builder
	padding int
}

func NewViewBuilder(width, height int) *ViewBuilder {
	return &ViewBuilder{
		width:   width,
		height:  height,
		padding: 8,
	}
}

func (vb *ViewBuilder) AddCentered(view string) *ViewBuilder {
	vb.doc.WriteString("\n" + Centered(vb.width).Render(view))
	return vb
}

func (vb *ViewBuilder) AddPadded(view string) *ViewBuilder {
	vb.doc.WriteString("\n" + lipgloss.NewStyle().Padding(1, 0, 0, 0).Render(view))
	return vb
}

func (vb *ViewBuilder) AddSpacing(lines int) *ViewBuilder {
	vb.doc.WriteString(strings.Repeat("\n", lines))
	return vb
}

func (vb *ViewBuilder) Build() string {
	content := vb.doc.String()
	padding := vb.height - lipgloss.Height(content) - vb.padding
	if padding < 0 {
		padding = 0
	}
	vb.doc.WriteString(strings.Repeat("\n", padding))
	return vb.doc.String()
}

func (vb *ViewBuilder) SetPadding(p int) *ViewBuilder {
	vb.padding = p
	return vb
}
