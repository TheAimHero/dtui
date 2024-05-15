package managecontianer

import (
	"fmt"
	"strings"

	"github.com/TheAimHero/dtui/internal/ui/message"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	ContainerSelected = iota
	ContainerID
	ContainerName
	ContainerImage
	ContainerStatus
)

func (m *containerModel) ClearSelectedContainers() {
	m.selectedContainers.Clear()
}

func (m containerModel) StartContainer() (tea.Model, tea.Cmd) {
	startMsg := message.Message{}
	row := m.table.SelectedRow()
	if row == nil {
		m.message.AddMessage("No container selected", message.ErrorMessage)
		return m, m.message.ClearMessage(message.ErrorDuration)
	}
	return m, func() tea.Msg {
		err := m.dockerClient.StartContainer(row[ContainerID])
		if err != nil {
			startMsg.AddMessage(fmt.Sprintf("Error while starting container: %s", strings.Split(err.Error(), ":")[1]), message.ErrorMessage)
			return startMsg
		}
		startMsg.AddMessage(fmt.Sprintf("Container %s started", m.table.SelectedRow()[ContainerName]), message.SuccessMessage)
    return startMsg
	}
}

func (m containerModel) StopContainer() (tea.Model, tea.Cmd) {
	stopMsg := message.Message{}
	return m, func() tea.Msg {
		err := m.dockerClient.StopContainer(m.table.SelectedRow()[ContainerID])
		if err != nil {
			stopMsg.AddMessage(fmt.Sprintf("Error while stopping container: %s", strings.Split(err.Error(), ":")[ContainerName]), message.ErrorMessage)
			return stopMsg
		}
		stopMsg.AddMessage(fmt.Sprintf("Container %s stopped", m.table.SelectedRow()[ContainerName]), message.SuccessMessage)
		return stopMsg
	}
}

func (m containerModel) StartContainers() (tea.Model, tea.Cmd) {
	startMsg := message.Message{}
	selectedContainers := m.selectedContainers.ToSlice()
	defer m.ClearSelectedContainers()
	errors := make([]string, 0)
	if len(selectedContainers) == 0 {
		m.message.AddMessage("No containers selected", message.ErrorMessage)
		return m, m.message.ClearMessage(message.ErrorDuration)
	}
	return m, func() tea.Msg {
		for _, containerID := range selectedContainers {
			err := m.dockerClient.StartContainer(containerID)
			if err != nil {
				errors = append(errors, err.Error())
			}
		}
		if len(errors) > 0 {
			startMsg.AddMessage("Error while starting some containers", message.ErrorMessage)
			m.selectedContainers.Clear()
			return startMsg
		}
		startMsg.AddMessage("Containers started", message.SuccessMessage)
		m.selectedContainers.Clear()
		return startMsg
	}
}

func (m containerModel) StopContainers() (tea.Model, tea.Cmd) {
	errors := make([]string, 0)
	selectedContainers := m.selectedContainers.ToSlice()
	stopMsg := message.Message{}
	return m, func() tea.Msg {
		for _, containerID := range selectedContainers {
			err := m.dockerClient.StopContainer(containerID)
			if err != nil {
				errors = append(errors, err.Error())
			}
		}
		if len(errors) > 0 {
			stopMsg.AddMessage("Error while stopping some containers", message.ErrorMessage)
			m.selectedContainers.Clear()
			return stopMsg
		}
		stopMsg.AddMessage("Containers stopped", message.SuccessMessage)
		m.selectedContainers.Clear()
		return stopMsg
	}
}

func (m containerModel) SelectContainers() (tea.Model, tea.Cmd) {
	containerID := m.table.SelectedRow()[ContainerID]
	if m.selectedContainers.Contains(containerID) {
		m.selectedContainers.Remove(containerID)
	} else {
		m.selectedContainers.Add(containerID)
	}
	m.table.MoveDown(1)
	return m, nil
}

func (m containerModel) SelectAllContainers() (tea.Model, tea.Cmd) {
	var allIDs []string
	for _, row := range m.table.Rows() {
		allIDs = append(allIDs, row[ContainerID])
	}
	if m.selectedContainers.Cardinality() == len(m.table.Rows()) {
		m.selectedContainers.Clear()
	} else {
		m.selectedContainers.Clear()
		m.selectedContainers.Append(allIDs...)
	}
	return m, nil
}

func (m containerModel) DeleteContainer() (tea.Model, tea.Cmd) {
	deleteMsg := message.Message{}
	row := m.table.SelectedRow()
	if row == nil {
		m.message.AddMessage("No container selected", message.ErrorMessage)
		return m, m.message.ClearMessage(message.ErrorDuration)
	}
	return m, func() tea.Msg {
		err := m.dockerClient.DeleteContainer(row[ContainerID])
		if err != nil {
			deleteMsg.AddMessage(fmt.Sprintf("Error while deleting container: %s", strings.Split(err.Error(), ":")[ContainerName]), message.ErrorMessage)
			return deleteMsg
		}
		deleteMsg.AddMessage(fmt.Sprintf("Container %s deleted", m.table.SelectedRow()[ContainerName]), message.SuccessMessage)
		return deleteMsg
	}
}

func (m containerModel) DeleteContainers() (tea.Model, tea.Cmd) {
	defer m.ClearSelectedContainers()
	selectedContainers := m.selectedContainers.ToSlice()
	if len(selectedContainers) == 0 {
		m.message.AddMessage("No containers selected", message.ErrorMessage)
		return m, m.message.ClearMessage(message.ErrorDuration)
	}
	deleteMsg := message.Message{}
	errors := make([]string, 0)
	return m, func() tea.Msg {
		for _, containerID := range selectedContainers {
			err := m.dockerClient.DeleteContainer(containerID)
			if err != nil {
				errors = append(errors, err.Error())
			}
		}
		if len(errors) > 0 {
			deleteMsg.AddMessage("Error while deleting some containers", message.ErrorMessage)
			m.selectedContainers.Clear()
			return deleteMsg

		}
		m.selectedContainers.Clear()
		deleteMsg.AddMessage("Containers deleted", message.SuccessMessage)
		return deleteMsg
	}
}
