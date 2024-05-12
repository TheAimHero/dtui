package managecontianer

import (
	"os"
	"strings"
	"time"

	"github.com/TheAimHero/dtui/internal/ui"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

var (
	physicalWidth, physicalHeight, _ = term.GetSize(int(os.Stdout.Fd()))
	errorDuration                    = 5 * time.Second
	successDuration                  = 2 * time.Second
)

func (m containerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		physicalWidth, physicalHeight, _ = term.GetSize(int(os.Stdout.Fd()))
		m.table = getTable(m.dockerClient.Containers, m.selectedContainers)
		return m, cmd

	case ui.ClearErrorMsg:
		m.message = ui.Message{}

	case time.Time:
		m.dockerClient.FetchContainers()
		tableRows := getTableRows(m.dockerClient.Containers, m.selectedContainers)
		m.table.SetRows(tableRows)
		return m, tickCommand()

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit

		case key.Matches(msg, m.keys.Down):
			m.table.MoveDown(1)
			return m, nil

		case key.Matches(msg, m.keys.Up):
			m.table.MoveUp(1)
			return m, nil

		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll

		case key.Matches(msg, m.keys.StopContainer):
			return m.StopContainer()

		case key.Matches(msg, m.keys.StartContainer):
			return m.StartContainer()

		case key.Matches(msg, m.keys.StartContainers):
			return m.StartContainers()

		case key.Matches(msg, m.keys.StopContainers):
			return m.StopContainers()

		case key.Matches(msg, m.keys.ToggleSelected):
			return m.SelectContainers()

		case key.Matches(msg, m.keys.ToggleSelectAll):
			return m.SelectAllContainers()
		}

	}
	return m, cmd
}

func (m containerModel) View() string {
	doc := strings.Builder{}
	align := lipgloss.NewStyle().Align(lipgloss.NoTabConversion)
	doc.WriteString(align.Render(ui.BaseTableStyle.Render(m.table.View()) + m.message.ShowMessage()))
	doc.WriteString("\n" + m.help.View(m.keys))
	return doc.String()
}
