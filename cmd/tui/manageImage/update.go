package manageimage

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

func (m ImageModel) updateInput(msg tea.KeyMsg) (ImageModel, tea.Cmd) {
	var cmd tea.Cmd

	switch {
	case key.Matches(msg, m.Keys.EscapeInput):
		m.Table.Focus()
		m.Input = textinput.Model{}
		m.Table, cmd = m.Table.Update(msg)
		m.Keys.ShowInput.SetEnabled(true)
		m.Keys.EscapeInput.SetEnabled(false)
		return m, cmd

	case key.Matches(msg, m.Keys.Submit):
		m.Table.Focus()
		m.Input.Blur()
		m.Keys.ShowInput.SetEnabled(true)
		m.Keys.EscapeInput.SetEnabled(false)
		m.Keys.Submit.SetEnabled(false)
		m, cmd = m.PullImage()
		m.Input.SetValue("")
		return m, tea.Batch(cmd)

	default:
		m.Input, cmd = m.Input.Update(msg)
		return m, cmd
	}
}

func (m ImageModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		physicalWidth, physicalHeight, _ = term.GetSize(int(os.Stdout.Fd()))
		m.Table = getTable(m.DockerClient.Images, m.SelectedImages, m.InProgress, m.LoadingSpinner)

	case message.ClearMessage:
		m.Message = message.Message{}

	case time.Time:
		err := m.DockerClient.FetchImages()
		if err != nil {
			m.Message.AddMessage("Error while fetching images", message.ErrorMessage)
			cmds = append(cmds, utils.TickCommand(), m.Message.ClearMessage(message.ErrorDuration))
		}
		tableRows := getTableRows(m.DockerClient.Images, m.SelectedImages, m.InProgress, m.LoadingSpinner)
		m.Table.SetRows(tableRows)
		cmds = append(cmds, utils.TickCommand())

	case message.Message:
		m.Message = msg
		var duration time.Duration
		if msg.MsgType == message.SuccessMessage {
			duration = message.SuccessDuration
		} else {
			duration = message.ErrorDuration
		}
		cmds = append(cmds, m.Message.ClearMessage(duration))

	case tea.KeyMsg:
		if m.Input.Focused() {
			return m.updateInput(msg)
		}
		switch {
		case key.Matches(msg, m.Keys.Quit):
			return m, tea.Quit

		case key.Matches(msg, m.Keys.Help):
			m.Help.ShowAll = !m.Help.ShowAll

		case key.Matches(msg, m.Keys.SelectImage):
			m, cmd = m.SelectImage()
			cmds = append(cmds, cmd)

		case key.Matches(msg, m.Keys.SelectAllImages):
			m, cmd = m.SelectAllImages()
			cmds = append(cmds, cmd)

		case key.Matches(msg, m.Keys.DeleteImages):
			m, cmd = m.DeleteImages()
			cmds = append(cmds, cmd)

		case key.Matches(msg, m.Keys.PruneImages):
			m, cmd = m.PruneImages()
			cmds = append(cmds, cmd)

		case key.Matches(msg, m.Keys.ShowInput):
			m.Input = getInput()
			m.Keys.ShowInput.SetEnabled(false)
			m.Keys.EscapeInput.SetEnabled(true)
			m.Keys.Submit.SetEnabled(true)
			cmds = append(cmds, m.Input.Focus())
		}
	}
	m.Table, cmd = m.Table.Update(msg)
	cmds = append(cmds, cmd)
	m.PullSpinner, cmd = m.PullSpinner.Update(msg)
	cmds = append(cmds, cmd)
	m.LoadingSpinner, cmd = m.LoadingSpinner.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}
