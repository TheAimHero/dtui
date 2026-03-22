package managevolume

import (
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
	ActionDeleteVolume
	ActionPruneVolume
)

func (m VolumeModel) updateInput(msg tea.KeyMsg) (VolumeModel, tea.Cmd) {
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
		m.Keys.PruneVolume.SetEnabled(true)
		return m, cmd

	default:
		m.Input, cmd = m.Input.Update(msg)
		m.Table.SetRows(filterRows(m.Table.Rows(), m.Input.Value()))
		m.Table.SetCursor(0)
		return m, cmd
	}
}

func (m VolumeModel) handleAction(action ActionType) (tea.Model, tea.Cmd) {
	switch action {
	case ActionDeleteVolume:
		return m.DeleteVolume()

	case ActionPruneVolume:
		return m.PruneVolume()
	}
	return m, nil
}

func (m VolumeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

	case time.Time:
		var err error
		m.Volumes, err = m.VolumeSvc.FetchVolumes()
		if err != nil {
			m.Message.AddMessage("Error while fetching volumes", message.ErrorMessage)
			cmds = append(cmds, m.Message.ClearMessage(message.ErrorDuration), utils.TickCommand())
		}
		tableRows := getTableRows(m.Volumes, m.SelectedVolumes, m.InProgress, m.Spinner)
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

		case key.Matches(msg, m.Keys.DeleteVolume):
			m.Confirmation = prompt.NewModel("Are you sure you want to delete volume?", func() tea.Msg { return ActionDeleteVolume })

		case key.Matches(msg, m.Keys.Help):
			m.Help.ShowAll = !m.Help.ShowAll

		case key.Matches(msg, m.Keys.PruneVolume):
			m.Confirmation = prompt.NewModel("Are you sure you want to prune volumes?", func() tea.Msg { return ActionPruneVolume })

		case key.Matches(msg, m.Keys.ShowInput):
			m.Input = getInput()
			m.Keys.ShowInput.SetEnabled(false)
			m.Keys.PruneVolume.SetEnabled(false)
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
