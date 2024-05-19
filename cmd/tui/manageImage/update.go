package manageimage

import (
	"io"
	"os"
	"strings"
	"time"

	"github.com/TheAimHero/dtui/internal/ui/message"
	"github.com/TheAimHero/dtui/internal/utils"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"golang.org/x/term"
)

func (m ImageModel) updateInput(msg tea.KeyMsg) (ImageModel, tea.Cmd) {
	var (
		cmd    tea.Cmd
		cmds   []tea.Cmd
		stream io.ReadCloser
	)
	switch msg.String() {
	case "esc":
		m.Table.Focus()
		m.Input = textinput.Model{}
		m.Table, cmd = m.Table.Update(msg)
		m.keys.ShowInput.SetEnabled(true)
		m.keys.EscapeInput.SetEnabled(false)
		return m, cmd

	case "enter":
		m.Table.Focus()
		imageName := m.Input.Value()
		m.keys.ShowInput.SetEnabled(true)
		m.Input = textinput.Model{}
		if len(imageName) == 0 {
			m.message.AddMessage("Please enter image name", message.ErrorMessage)
			return m, m.message.ClearMessage(message.ErrorDuration)
		}
		m, cmd, stream = m.PullImages(imageName)
		cmds = append(cmds, cmd)
		m.Table, cmd = m.Table.Update(msg)
		cmds = append(cmds, cmd)
		m.viewport, cmd = m.viewport.Update(msg)
		cmds = append(cmds, cmd)
		cmds = append(cmds, utils.ListenToStream(m.sub, stream))
		return m, tea.Batch(cmds...)

	default:
		m.Input, cmd = m.Input.Update(msg)
		return m, cmd
	}
}

func (m ImageModel) Update(msg tea.Msg) (ImageModel, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		physicalWidth, physicalHeight, _ = term.GetSize(int(os.Stdout.Fd()))
		m.Table = getTable(m.dockerClient.Images, m.selectedImages)
		m.viewport.Width = msg.Width - 20

	case message.ClearMessage:
		m.message = message.Message{}

	case utils.ResponseMsg:
		m.text = append(m.text, string(msg))
		m.viewport.SetContent(contentStyle.Render(strings.Join(m.text, "\n")))
		m.viewport.GotoBottom()
		cmds = append(cmds, utils.ResponseToStream(m.sub))

	case time.Time:
		err := m.dockerClient.FetchImages()
		if err != nil {
			m.message.AddMessage("Error while fetching images", message.ErrorMessage)
			cmds = append(cmds, utils.TickCommand(), m.message.ClearMessage(message.ErrorDuration))
		}
		tableRows := getTableRows(m.dockerClient.Images, m.selectedImages)
		m.Table.SetRows(tableRows)
		cmds = append(cmds, utils.TickCommand())

	case message.Message:
		m.message = msg
		var duration time.Duration
		if msg.MsgType == message.SuccessMessage {
			duration = message.SuccessDuration
		} else {
			duration = message.ErrorDuration
		}
		cmds = append(cmds, m.message.ClearMessage(duration))

	case tea.KeyMsg:
		if m.Input.Focused() {
			return m.updateInput(msg)
		}
		switch {
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit

		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll

		case key.Matches(msg, m.keys.SelectImage):
			m, cmd = m.SelectImage()
			cmds = append(cmds, cmd)

		case key.Matches(msg, m.keys.SelectAllImages):
			m, cmd = m.SelectAllImages()
			cmds = append(cmds, cmd)

		case key.Matches(msg, m.keys.DeleteImage):
			m, cmd = m.DeleteImage()
			cmds = append(cmds, cmd)

		case key.Matches(msg, m.keys.DeleteSelectedImages):
			m, cmd = m.DeleteImages()
			cmds = append(cmds, cmd)

		case key.Matches(msg, m.keys.ShowInput):
			m.Input = getInput()
			m.keys.ShowInput.SetEnabled(false)
			m.keys.EscapeInput.SetEnabled(true)
			cmds = append(cmds, m.Input.Focus())

		}
	}
	m.Table, cmd = m.Table.Update(msg)
	cmds = append(cmds, cmd)
	m.viewport, cmd = m.viewport.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}
