package logs

import (
	"github.com/TheAimHero/dtui/internal/ui/message"
	tea "github.com/charmbracelet/bubbletea"
)

func (m *logModel) GetLogs() (logModel, tea.Cmd) {
	var (
		cmds []tea.Cmd
		err  error
	)
	row := m.table.SelectedRow()
	if row == nil {
		return *m, nil
	}
	containerID := row[ContainerID]
	m.title = row[ContainerName]
	m.text = []string{}
	// close the previeous stream if existes
	if m.stream != nil {
		m.stream.Close()
	}
	m.stream, err = m.dockerClient.GetLogs(containerID)
	if err != nil {
		m.message.AddMessage("Error while fetching logs", message.ErrorMessage)
		cmds = append(cmds, m.message.ClearMessage(message.ErrorDuration))
	} else {
		m.message.AddMessage("Logs fetched for: "+row[ContainerName], message.SuccessMessage)
		cmds = append(cmds, m.message.ClearMessage(message.SuccessDuration))
	}
	return *m, tea.Batch(cmds...)
}
