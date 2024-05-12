package managecontianer

import (
	"fmt"
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
)

func (m containerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {

	case time.Time:
		m.dockerClient.FetchContainers()
		tableRows := getTableRows(m.dockerClient.Containers)
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
			err := m.dockerClient.StopContainer(m.table.SelectedRow()[1])
			if err != nil {
				m.message.AddMessage(fmt.Sprintf("Error while stopping container: %s", strings.Split(err.Error(), ":")[1]), "error")
				return m, nil
			}
			m.message.AddMessage(fmt.Sprintf("Container %s stopped", m.table.SelectedRow()[1]), "success")
			return m, nil

		case key.Matches(msg, m.keys.StartContainer):
			err := m.dockerClient.StartContainer(m.table.SelectedRow()[1])
			if err != nil {
				m.message.AddMessage(fmt.Sprintf("Error while starting container: %s", strings.Split(err.Error(), ":")[1]), "error")
				return m, nil
			}
			m.message.AddMessage(fmt.Sprintf("Container %s started", m.table.SelectedRow()[1]), "success")
			return m, nil
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
