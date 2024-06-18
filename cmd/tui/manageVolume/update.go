package managevolume

import (
	"os"
	"time"

	"github.com/TheAimHero/dtui/internal/ui/message"
	"github.com/TheAimHero/dtui/internal/utils"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"golang.org/x/term"
)

func (m VolumeModel) Update(msg tea.Msg) (VolumeModel, tea.Cmd) {
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
		err := m.DockerClient.FetchVolumes()
		if err != nil {
			m.Message.AddMessage("Error while fetching volumes", message.ErrorMessage)
			cmds = append(cmds, m.Message.ClearMessage(message.ErrorDuration), utils.TickCommand())
		}
		tableRows := getTableRows(
			m.DockerClient.Volumes,
			m.SelectedVolumes,
		)
		m.Table.SetRows(tableRows)
		cmds = append(cmds, utils.TickCommand(), cmd)

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.Keys.Quit):
			return m, tea.Quit

		case key.Matches(msg, m.Keys.Help):
			m.Help.ShowAll = !m.Help.ShowAll

		case key.Matches(msg, m.Keys.PruneVolume):
			m, cmd = m.PruneVolume()
			cmds = append(cmds, cmd)
		}
	}
	m.Table, cmd = m.Table.Update(msg)
	cmds = append(cmds, cmd)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}
