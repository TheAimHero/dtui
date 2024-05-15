package managecontianer

import (
	"os"
	"strings"
	"time"

	"github.com/TheAimHero/dtui/internal/ui/message"
	ui_table "github.com/TheAimHero/dtui/internal/ui/table"
	ui_utils "github.com/TheAimHero/dtui/internal/ui/utils"
	"github.com/TheAimHero/dtui/internal/utils"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"golang.org/x/term"
)

var (
	physicalWidth, physicalHeight, _ = term.GetSize(int(os.Stdout.Fd())) // nolint:unused
)

func (m containerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		physicalWidth, physicalHeight, _ = term.GetSize(int(os.Stdout.Fd()))
		m.table = getTable(m.dockerClient.Containers, m.selectedContainers)
		return m, cmd

	case message.ClearErrorMsg:
		m.message = message.Message{}

	case message.Message:
		m.message = msg
		var duration time.Duration
		if msg.MsgType == message.SuccessMessage {
			duration = message.SuccessDuration
		} else {
			duration = message.ErrorDuration
		}
		return m, m.message.ClearMessage(duration)

	case time.Time:
		err := m.dockerClient.FetchContainers()
		if err != nil {
			m.message.AddMessage("Error while fetching containers", message.ErrorMessage)
			return m, tea.Batch(m.message.ClearMessage(message.ErrorDuration), utils.TickCommand())
		}
		tableRows := getTableRows(m.dockerClient.Containers, m.selectedContainers)
		m.table.SetRows(tableRows)

		return m, utils.TickCommand()

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit

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

		case key.Matches(msg, m.keys.DeleteContainer):
			return m.DeleteContainer()

		case key.Matches(msg, m.keys.DeleteContainers):
			return m.DeleteContainers()
		}

	}
	m.table, cmd = m.table.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m containerModel) View() string {
	doc := strings.Builder{}
	// doc.WriteString(m.execTime.View())
	doc.WriteString(ui_table.BaseTableStyle.Render(m.table.View()) + m.message.ShowMessage())
	doc.WriteString("\n" + m.help.View(m.keys))
	doc.WriteString(strings.Repeat("\n", ui_utils.HeightPadding(doc, 7)))
	return doc.String()
}
