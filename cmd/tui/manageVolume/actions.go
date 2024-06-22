package managevolume

import (
	"github.com/TheAimHero/dtui/internal/ui/message"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	VolumeSelected = iota
	VolumeName
	VolumeCreated
	VolumeMountpoint
	VolumeSize
)

func (m VolumeModel) PruneVolume() (VolumeModel, tea.Cmd) {
	deleteMsg := message.Message{}
	row := m.Table.SelectedRow()
	if row == nil {
		m.Message.AddMessage("No volume selected", message.InfoMessage)
		return m, m.Message.ClearMessage(message.InfoDuration)
	}
	return m, func() tea.Msg {
		err := m.DockerClient.PruneVolume()
		if err != nil {
			deleteMsg.AddMessage("Error while pruning volumes", message.ErrorMessage)
			return deleteMsg
		}
		deleteMsg.AddMessage("Volumes Pruned", message.SuccessMessage)
		m.Table.SetCursor(0)
		return deleteMsg
	}
}
