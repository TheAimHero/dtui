package logs

import (
	"github.com/TheAimHero/dtui/internal/ui/message"
	tea "github.com/charmbracelet/bubbletea"
)

func (m *LogModel) GetLogs() (LogModel, tea.Cmd) {
	var (
		cmds []tea.Cmd
		err  error
	)
	row := m.Table.SelectedRow()
	if row == nil {
		return *m, nil
	}
	containerID := row[ContainerID]
	m.Title = row[ContainerName]
	m.Text = []string{}
	// close the previeous stream if existes
	if m.Stream != nil {
		m.Stream.Close()
	}
	m.Stream, err = m.DockerClient.GetLogs(containerID)
	if err != nil {
		m.Message.AddMessage("Error while fetching logs", message.ErrorMessage)
		cmds = append(cmds, m.Message.ClearMessage(message.ErrorDuration))
	} else {
		m.Message.AddMessage("Logs fetched for: "+row[ContainerName], message.SuccessMessage)
		cmds = append(cmds, m.Message.ClearMessage(message.SuccessDuration))
	}
	return *m, tea.Batch(cmds...)
}
