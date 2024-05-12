package manageimage

import (
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"

	"github.com/TheAimHero/dtui/internal/ui"
)

var (
	highlightColor                   = lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"}
	helpStyle                        = lipgloss.NewStyle().Align(lipgloss.Bottom)
	physicalWidth, physicalHeight, _ = term.GetSize(int(os.Stdout.Fd()))
	successDuration                  = 2 * time.Second
	errorDuration                    = 5 * time.Second
)

func heightPadding(doc strings.Builder) int {
	paddingHeight := physicalHeight - lipgloss.Height(doc.String()) - 7
	if paddingHeight < 0 {
		paddingHeight = 0
	}
	return paddingHeight
}

func (m imageModel) View() string {
	doc := strings.Builder{}
	doc.WriteString(ui.BaseTableStyle.Render(m.table.View()))
	doc.WriteString("\n" + m.message.ShowMessage())
	doc.WriteString("\n" + m.help.View(m.keys))
	doc.WriteString(strings.Repeat("\n", heightPadding(doc)))
	return doc.String()
}

func (m imageModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		physicalWidth, physicalHeight, _ = term.GetSize(int(os.Stdout.Fd()))
		m.table = getTable(m.dockerClient.Images, m.selectedImages)
		return m, cmd

	case ui.ClearErrorMsg:
		m.message = ui.Message{}

	case time.Time:
		m.dockerClient.FetchImages()
		tableRows := getTableRows(m.dockerClient.Images, m.selectedImages)
		m.table.SetRows(tableRows)
		return m, tickCommand()

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit

		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll

		case key.Matches(msg, m.keys.SelectImage):
			return m.SelectImage()

		case key.Matches(msg, m.keys.SelectAllImages):
			return m.SelectAllImages()

		case key.Matches(msg, m.keys.DeleteImage):
			return m.DeleteImage()

		case key.Matches(msg, m.keys.DeleteSelectedImages):
			return m.DeleteImages()
		}
	}

	m.table, cmd = m.table.Update(msg)
	return m, cmd
}
