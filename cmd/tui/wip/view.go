package wip

import (
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

var (
	physicalWidth, physicalHeight, _ = term.GetSize(int(os.Stdout.Fd()))
)

func heightPadding(doc strings.Builder) int {
	paddingHeight := physicalHeight - lipgloss.Height(doc.String()) - 7
	if paddingHeight < 0 {
		paddingHeight = 0
	}
	return paddingHeight
}

func (m wipModel) View() string {
	doc := strings.Builder{}
	baseStyle := lipgloss.NewStyle().Padding(1, 2)
	emphasisStyle := baseStyle.Copy().Bold(true).Foreground(lipgloss.Color("#F1FA8C"))
	titleStyle := baseStyle.Copy().Bold(true).Italic(true).Foreground(lipgloss.Color("#05C3DD"))

	doc.WriteString(titleStyle.Render("Work in progress... and Coming Soon..."))
	doc.WriteString("\n" + "What's next?\n")
	doc.WriteString(emphasisStyle.Render("-\tBuild Mode"))
	doc.WriteString(emphasisStyle.Render("-\tImage Mode"))
	doc.WriteString(emphasisStyle.Render("-\tLog Mode"))
	doc.WriteString(titleStyle.UnsetPadding().Render("\nGive a Star on GitHub\n"))
	doc.WriteString(titleStyle.Render("Thats all folks!"))
	doc.WriteString(strings.Repeat("\n", heightPadding(doc)))
	return doc.String()
}

func (m wipModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		physicalWidth, physicalHeight, _ = term.GetSize(int(os.Stdout.Fd()))
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "enter", " ":
		}
	}

	return m, nil
}
