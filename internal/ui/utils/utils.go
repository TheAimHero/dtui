package utils

import (
	"os"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

var (
	physicalWidth, physicalHeight, _ = term.GetSize(int(os.Stdout.Fd())) // nolint:unused
)

func HeightPadding(doc strings.Builder, fixHeight int) int {
	paddingHeight := physicalHeight - lipgloss.Height(doc.String()) - fixHeight
	if paddingHeight < 0 {
		paddingHeight = 0
	}
	return paddingHeight
}
