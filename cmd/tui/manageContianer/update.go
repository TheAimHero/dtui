package managecontianer

import (
	"os"
	"time"

	"github.com/TheAimHero/dtui/internal/ui/message"
	"github.com/TheAimHero/dtui/internal/ui/prompt"
	"github.com/TheAimHero/dtui/internal/utils"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"golang.org/x/term"
)

type ActionType int

const (
	ActionNoOp ActionType = iota
	ActionStartContainer
	ActionStopContainer
	ActionDeleteContainer
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
		m.Keys.ShowLogs.SetEnabled(true)
		return m, cmd

	default:
		m.Input, cmd = m.Input.Update(msg)
		m.Table.SetRows(filterRows(m.Table.Rows(), m.Input.Value()))
		m.Table.SetCursor(0)
		return m, cmd
	}
}

func (m ContainerModel) handleAction(action ActionType) (tea.Model, tea.Cmd) {
	switch action {
	case ActionStartContainer:
		return m.StartContainers()

	case ActionStopContainer:
		return m.StopContainers()

	case ActionDeleteContainer:
		return m.DeleteContainers()
	}
	return m, nil
}

func (m ContainerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

	case ActionType:
		return m.handleAction(msg)

	case tea.KeyMsg:
		if m.Input.Focused() {
			return m.updateInput(msg)
		}
		if m.Confirmation.Active {
			m.Confirmation, cmd = m.Confirmation.Update(msg)
			cmds = append(cmds, cmd)
			return m, tea.Batch(cmds...)
		}
		switch {
		case key.Matches(msg, m.Keys.Quit):
			return m, tea.Quit

		case key.Matches(msg, m.Keys.Help):
			m.Help.ShowAll = !m.Help.ShowAll

		case key.Matches(msg, m.Keys.StartContainers):
			m.Confirmation = prompt.NewModel("Are you sure you want to start selected containers?", func() tea.Msg { return ActionStartContainer })

		case key.Matches(msg, m.Keys.StopContainers):
			m.Confirmation = prompt.NewModel("Are you sure you want to stop selected containers?", func() tea.Msg { return ActionStopContainer })

		case key.Matches(msg, m.Keys.ToggleSelected):
			m, cmd = m.SelectContainers()
			cmds = append(cmds, cmd)

		case key.Matches(msg, m.Keys.ToggleSelectAll):
			m, cmd = m.SelectAllContainers()
			cmds = append(cmds, cmd)

		case key.Matches(msg, m.Keys.DeleteContainers):
			m.Confirmation = prompt.NewModel("Are you sure you want to delete selected containers?", func() tea.Msg { return ActionDeleteContainer })

		case key.Matches(msg, m.Keys.ShowLogs):
			m, cmd = m.ShowLogs()
			cmds = append(cmds, cmd)

		case key.Matches(msg, m.Keys.ExecContainer):
			m, cmd = m.ExecContainer()
			cmds = append(cmds, cmd)

		case key.Matches(msg, m.Keys.ShowInput):
			m.Input = getInput()
			m.Keys.ShowInput.SetEnabled(false)
			m.Keys.ShowLogs.SetEnabled(false)
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
