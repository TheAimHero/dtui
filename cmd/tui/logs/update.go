package logs

import (
	"os"
	"time"

	"github.com/TheAimHero/dtui/internal/ui/message"
	"github.com/TheAimHero/dtui/internal/utils"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"golang.org/x/term"
)

func (m LogModel) Update(msg tea.Msg) (LogModel, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)
	switch msg := msg.(type) {
	case utils.ResponseMsg:
		m.Text = append(m.Text, string(msg))
		cmds = append(cmds, utils.ResponseToStream(m.Sub))

	case message.ClearMessage:
		m.Message = message.Message{}

	case time.Time:
		err := m.DockerClient.FetchContainers()
		if err != nil {
			m.Message.AddMessage("Error while fetching containers", message.ErrorMessage)
			cmds = append(cmds, m.Message.ClearMessage(message.ErrorDuration))
		}
		tableRows := getTableRows(m.DockerClient.Containers)
		m.Table.SetRows(tableRows)
		cmds = append(cmds, utils.TickCommand())

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.Keys.Quit):
			return m, tea.Quit

		case key.Matches(msg, m.Keys.Help):
			m.Help.ShowAll = !m.Help.ShowAll

		case key.Matches(msg, m.Keys.Select):
			cmds = append(cmds, showLogs(m.Table.SelectedRow()[ContainerID]))
		}

	case tea.WindowSizeMsg:
		physicalWidth, _, _ = term.GetSize(int(os.Stdout.Fd()))
		m.Table = getTable(m.DockerClient.Containers)
	}
	m.Table, cmd = m.Table.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}
