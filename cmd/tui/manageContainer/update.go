package managecontainer

import (
	"fmt"
	"time"

	"github.com/TheAimHero/dtui/internal/ui/components"
	"github.com/TheAimHero/dtui/internal/ui/message"
	"github.com/TheAimHero/dtui/internal/ui/prompt"
	"github.com/TheAimHero/dtui/internal/utils"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type ActionType int

const (
	ActionNoOp ActionType = iota
	ActionStartContainer
	ActionStopContainer
	ActionDeleteContainer
)

func actionVerb(action ActionType) string {
	switch action {
	case ActionStartContainer:
		return "starting"
	case ActionStopContainer:
		return "stopping"
	case ActionDeleteContainer:
		return "deleting"
	}
	return ""
}

func actionVerbPast(action ActionType) string {
	switch action {
	case ActionStartContainer:
		return "started"
	case ActionStopContainer:
		return "stopped"
	case ActionDeleteContainer:
		return "deleted"
	}
	return ""
}

func (m ContainerModel) updateInput(msg tea.KeyMsg) (ContainerModel, tea.Cmd) {
	var cmd tea.Cmd

	switch msg.String() {
	case "esc":
		m.Table.Focus()
		m.Input = textinput.Model{}
		m.Table, cmd = m.Table.Update(msg)
		m.containerKeys.ShowInput.SetEnabled(true)
		m.containerKeys.EscapeInput.SetEnabled(false)
		m.containerKeys.SetFilter.SetEnabled(false)
		return m, cmd

	case "enter":
		m.Table.Focus()
		m.Input.Blur()
		m.containerKeys.ShowInput.SetEnabled(true)
		m.containerKeys.EscapeInput.SetEnabled(false)
		m.containerKeys.SetFilter.SetEnabled(false)
		m.containerKeys.ShowLogs.SetEnabled(true)
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
		m.Width = msg.Width
		m.Height = msg.Height
		tableConfig := components.DefaultTableConfig().WithOffset(12)
		tableHeight := components.CalculateTableHeight(m.Height, tableConfig)
		m.Table = m.getTable()
		m.Table.SetHeight(tableHeight)

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

	case ContainerActionResult:
		m.InProcess.Remove(msg.ContainerID)
		if msg.Err != nil {
			m.Message.AddMessage(fmt.Sprintf("Error while %s container: %s", actionVerb(msg.Action), msg.Err.Error()), message.ErrorMessage)
			cmds = append(cmds, m.Message.ClearMessage(message.ErrorDuration))
		} else {
			m.Message.AddMessage(fmt.Sprintf("Container %s %s", msg.ContainerID, actionVerbPast(msg.Action)), message.SuccessMessage)
			cmds = append(cmds, m.Message.ClearMessage(message.SuccessDuration))
		}

	case ContainerBatchResult:
		for _, result := range msg.Results {
			m.InProcess.Remove(result.ContainerID)
		}
		hasErrors := false
		for _, result := range msg.Results {
			if result.Err != nil {
				hasErrors = true
				break
			}
		}
		if hasErrors {
			m.Message.AddMessage(fmt.Sprintf("Error while %s some containers", actionVerb(msg.Action)), message.ErrorMessage)
			cmds = append(cmds, m.Message.ClearMessage(message.ErrorDuration))
		} else {
			m.Message.AddMessage(fmt.Sprintf("Containers %s", actionVerbPast(msg.Action)), message.SuccessMessage)
			cmds = append(cmds, m.Message.ClearMessage(message.SuccessDuration))
		}

	case time.Time:
		containers, err := m.ContainerSvc.FetchContainers()
		if err != nil {
			m.Message.AddMessage("Error while fetching containers", message.ErrorMessage)
			cmds = append(cmds, m.Message.ClearMessage(message.ErrorDuration), utils.TickCommand())
		} else {
			m.Containers = containers
		}
		tableRows := getTableRows(
			m.Containers,
			m.SelectedContainers,
			m.InProcess,
			m.Spinner,
		)
		m.Table.SetRows(filterRows(tableRows, m.Input.Value()))
		m.Input, cmd = m.Input.Update(msg)
		cmds = append(cmds, utils.TickCommand(), cmd)

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

		case key.Matches(msg, m.containerKeys.StartContainers):
			m.Confirmation = prompt.NewModel("Are you sure you want to start selected containers?", func() tea.Msg { return ActionStartContainer })

		case key.Matches(msg, m.containerKeys.StopContainers):
			m.Confirmation = prompt.NewModel("Are you sure you want to stop selected containers?", func() tea.Msg { return ActionStopContainer })

		case key.Matches(msg, m.containerKeys.ToggleSelected):
			m, cmd = m.SelectContainers()
			cmds = append(cmds, cmd)

		case key.Matches(msg, m.containerKeys.ToggleSelectAll):
			m, cmd = m.SelectAllContainers()
			cmds = append(cmds, cmd)

		case key.Matches(msg, m.containerKeys.DeleteContainers):
			m.Confirmation = prompt.NewModel("Are you sure you want to delete selected containers?", func() tea.Msg { return ActionDeleteContainer })

		case key.Matches(msg, m.containerKeys.ShowLogs):
			m, cmd = m.ShowLogs()
			cmds = append(cmds, cmd)

		case key.Matches(msg, m.containerKeys.ExecContainer):
			m, cmd = m.ExecContainer()
			cmds = append(cmds, cmd)

		case key.Matches(msg, m.containerKeys.ShowInput):
			m.Input = getInput()
			m.containerKeys.ShowInput.SetEnabled(false)
			m.containerKeys.ShowLogs.SetEnabled(false)
			m.containerKeys.EscapeInput.SetEnabled(true)
			m.containerKeys.SetFilter.SetEnabled(true)
			cmds = append(cmds, m.Input.Focus())
		}
	}
	m.Table, cmd = m.Table.Update(msg)
	cmds = append(cmds, cmd)
	m.Spinner, cmd = m.Spinner.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}
