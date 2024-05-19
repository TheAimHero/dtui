package managecontianer

import (
	"os"
	"time"

	"github.com/TheAimHero/dtui/internal/ui/message"
	"github.com/TheAimHero/dtui/internal/utils"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"golang.org/x/term"
)

func (m ContainerModel) Update(msg tea.Msg) (ContainerModel, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		physicalWidth, physicalHeight, _ = term.GetSize(int(os.Stdout.Fd()))
		m.table = getTable(m.dockerClient.Containers, m.selectedContainers)

	case message.ClearMessage:
		m.message = message.Message{}

	case message.Message:
		m.message = msg
		var duration time.Duration
		if msg.MsgType == message.SuccessMessage {
			duration = message.SuccessDuration
		} else {
			duration = message.ErrorDuration
		}
		cmds = append(cmds, m.message.ClearMessage(duration))

	case time.Time:
		err := m.dockerClient.FetchContainers()
		if err != nil {
			m.message.AddMessage("Error while fetching containers", message.ErrorMessage)
			cmds = append(cmds, m.message.ClearMessage(message.ErrorDuration), utils.TickCommand())
		}
		tableRows := getTableRows(m.dockerClient.Containers, m.selectedContainers)
		m.table.SetRows(tableRows)
		cmds = append(cmds, utils.TickCommand())

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit

		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll

		case key.Matches(msg, m.keys.StopContainer):
			m, cmd = m.StopContainer()
			cmds = append(cmds, cmd)

		case key.Matches(msg, m.keys.StartContainer):
			m, cmd = m.StartContainer()
			cmds = append(cmds, cmd)

		case key.Matches(msg, m.keys.StartContainers):
			m, cmd = m.StartContainers()
			cmds = append(cmds, cmd)

		case key.Matches(msg, m.keys.StopContainers):
			m, cmd = m.StopContainers()
			cmds = append(cmds, cmd)

		case key.Matches(msg, m.keys.ToggleSelected):
			m, cmd = m.SelectContainers()
			cmds = append(cmds, cmd)

		case key.Matches(msg, m.keys.ToggleSelectAll):
			m, cmd = m.SelectAllContainers()
			cmds = append(cmds, cmd)

		case key.Matches(msg, m.keys.DeleteContainer):
			m, cmd = m.DeleteContainer()
			cmds = append(cmds, cmd)

		case key.Matches(msg, m.keys.DeleteContainers):
			m, cmd = m.DeleteContainers()
			cmds = append(cmds, cmd)
		}
	}
	m.table, cmd = m.table.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}
