package managecontianer

import (
	"fmt"
	"strings"

	"github.com/TheAimHero/dtui/internal/helpers"
	"github.com/TheAimHero/dtui/internal/ui"
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
	m.selectedContainers = []string{}
}

func (m containerModel) StartContainer() (tea.Model, tea.Cmd) {
	err := m.dockerClient.StartContainer(m.table.SelectedRow()[1])
	if err != nil {
		m.message.AddMessage(fmt.Sprintf("Error while starting container: %s", strings.Split(err.Error(), ":")[ContainerName]), ui.ErrorMessage)
		return m, m.message.ClearMessage(errorDuration)
	}
	m.message.AddMessage(fmt.Sprintf("Container %s started", m.table.SelectedRow()[ContainerName]), ui.SuccessMessage)
	return m, m.message.ClearMessage(successDuration)
}

func (m containerModel) StopContainer() (tea.Model, tea.Cmd) {
	err := m.dockerClient.StopContainer(m.table.SelectedRow()[ContainerID])
	if err != nil {
		m.message.AddMessage(fmt.Sprintf("Error while stopping container: %s", strings.Split(err.Error(), ":")[ContainerName]), ui.ErrorMessage)
		return m, m.message.ClearMessage(errorDuration)
	}
	m.message.AddMessage(fmt.Sprintf("Container %s stopped", m.table.SelectedRow()[ContainerName]), ui.SuccessMessage)
	return m, m.message.ClearMessage(successDuration)
}

func (m containerModel) StartContainers() (tea.Model, tea.Cmd) {
	defer m.ClearSelectedContainers()
	errors := make([]string, 0)
	for _, containerID := range m.selectedContainers {
		err := m.dockerClient.StartContainer(containerID)
		if err != nil {
			errors = append(errors, err.Error())
		}
	}
	if len(errors) > 0 {
		m.message.AddMessage("Error while starting some containers", ui.ErrorMessage)
		m.selectedContainers = []string{}
		return m, m.message.ClearMessage(errorDuration)
	}
	m.message.AddMessage("Containers started", ui.SuccessMessage)
	m.selectedContainers = []string{}
	return m, m.message.ClearMessage(successDuration)
}

func (m containerModel) StopContainers() (tea.Model, tea.Cmd) {
	errors := make([]string, 0)
	for _, containerID := range m.selectedContainers {
		err := m.dockerClient.StopContainer(containerID)
		if err != nil {
			errors = append(errors, err.Error())
		}
	}
	if len(errors) > 0 {
		m.message.AddMessage("Error while stopping some containers", ui.ErrorMessage)
	m.selectedContainers = []string{}
		return m, m.message.ClearMessage(errorDuration)
	}
	m.message.AddMessage("Containers stopped", ui.SuccessMessage)
	m.selectedContainers = []string{}
	return m, m.message.ClearMessage(successDuration)
}

func (m containerModel) SelectContainers() (tea.Model, tea.Cmd) {
	containerID := m.table.SelectedRow()[ContainerID]
	if helpers.InArray(containerID, m.selectedContainers) {
		m.selectedContainers = helpers.RemoveFromArray(containerID, m.selectedContainers)
	} else {
		m.selectedContainers = append(m.selectedContainers, containerID)
	}
  m.table.MoveDown(1)
	return m, nil
}
