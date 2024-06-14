package managecontianer

import (
	"os"
	"time"

	"github.com/TheAimHero/dtui/internal/ui/message"
	"github.com/TheAimHero/dtui/internal/utils"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"golang.org/x/term"
)

func (m ContainerModel) updateInput(msg tea.KeyMsg) (ContainerModel, tea.Cmd) {
	var cmd tea.Cmd

	switch msg.String() {
	case "esc":
		m.Table.Focus()
		m.Input = textinput.Model{}
		m.Table, cmd = m.Table.Update(msg)
		m.Keys.ShowInput.SetEnabled(true)
		m.Keys.EscapeInput.SetEnabled(false)
		m.Keys.SetFilter.SetEnabled(false)
		return m, cmd

	case "enter":
		m.Table.Focus()
		m.Input.Blur()
		m.Keys.ShowInput.SetEnabled(true)
		m.Keys.EscapeInput.SetEnabled(false)
		m.Keys.SetFilter.SetEnabled(false)
		return m, cmd

	default:
		m.Input, cmd = m.Input.Update(msg)
		m.Table.SetRows(filterRows(m.Table.Rows(), m.Input.Value()))
		m.Table.SetCursor(0)
		return m, cmd
	}
}

func (m ContainerModel) Update(msg tea.Msg) (ContainerModel, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		physicalWidth, physicalHeight, _ = term.GetSize(int(os.Stdout.Fd()))
		m.Table = m.getTable()

	case message.ClearMessage:
		m.Message = message.Message{}

	case message.Message:
		m.Message = msg
		var duration time.Duration
		if msg.MsgType == message.SuccessMessage {
			duration = message.SuccessDuration
		} else {
			duration = message.ErrorDuration
		}
		cmds = append(cmds, m.Message.ClearMessage(duration))

	case time.Time:
		err := m.DockerClient.FetchContainers()
		if err != nil {
			m.Message.AddMessage("Error while fetching containers", message.ErrorMessage)
			cmds = append(cmds, m.Message.ClearMessage(message.ErrorDuration), utils.TickCommand())
		}
		tableRows := getTableRows(
			m.DockerClient.Containers,
			m.SelectedContainers,
			m.InProcess,
			m.Spinner,
		)
		m.Table.SetRows(filterRows(tableRows, m.Input.Value()))
		m.Input, cmd = m.Input.Update(msg)
		cmds = append(cmds, utils.TickCommand(), tea.Println(m.Input.Value()), cmd)

	case tea.KeyMsg:
		if m.Input.Focused() {
			return m.updateInput(msg)
		}
		switch {
		case key.Matches(msg, m.Keys.Quit):
			return m, tea.Quit

		case key.Matches(msg, m.Keys.Help):
			m.Help.ShowAll = !m.Help.ShowAll

		case key.Matches(msg, m.Keys.StopContainer):
			m, cmd = m.StopContainer()
			cmds = append(cmds, cmd)

		case key.Matches(msg, m.Keys.StartContainer):
			m, cmd = m.StartContainer()
			cmds = append(cmds, cmd)

		case key.Matches(msg, m.Keys.StartContainers):
			m, cmd = m.StartContainers()
			cmds = append(cmds, cmd)

		case key.Matches(msg, m.Keys.StopContainers):
			m, cmd = m.StopContainers()
			cmds = append(cmds, cmd)

		case key.Matches(msg, m.Keys.ToggleSelected):
			m, cmd = m.SelectContainers()
			cmds = append(cmds, cmd)

		case key.Matches(msg, m.Keys.ToggleSelectAll):
			m, cmd = m.SelectAllContainers()
			cmds = append(cmds, cmd)

		case key.Matches(msg, m.Keys.DeleteContainer):
			m, cmd = m.DeleteContainer()
			cmds = append(cmds, cmd)

		case key.Matches(msg, m.Keys.DeleteContainers):
			m, cmd = m.DeleteContainers()
			cmds = append(cmds, cmd)

		case key.Matches(msg, m.Keys.ShowInput):
			m.Input = getInput()
			m.Keys.ShowInput.SetEnabled(false)
			m.Keys.EscapeInput.SetEnabled(true)
			m.Keys.SetFilter.SetEnabled(true)
			cmds = append(cmds, m.Input.Focus())
		}
	}
	m.Table, cmd = m.Table.Update(msg)
	cmds = append(cmds, cmd)
	m.Spinner, cmd = m.Spinner.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}
