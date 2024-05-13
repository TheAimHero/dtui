package managecontianer

import (
	"fmt"
	"strings"

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
	m.selectedContainers.Clear()
}

func (m containerModel) StartContainer() (tea.Model, tea.Cmd) {
	err := m.dockerClient.StartContainer(m.table.SelectedRow()[ContainerID])
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
	for _, containerID := range m.selectedContainers.ToSlice() {
		err := m.dockerClient.StartContainer(containerID)
		if err != nil {
			errors = append(errors, err.Error())
		}
	}
	if len(errors) > 0 {
		m.message.AddMessage("Error while starting some containers", ui.ErrorMessage)
		m.selectedContainers.Clear()
		return m, m.message.ClearMessage(errorDuration)
	}
	m.message.AddMessage("Containers started", ui.SuccessMessage)
	m.selectedContainers.Clear()
	return m, m.message.ClearMessage(successDuration)
}

func (m containerModel) StopContainers() (tea.Model, tea.Cmd) {
	errors := make([]string, 0)
	for _, containerID := range m.selectedContainers.ToSlice() {
		err := m.dockerClient.StopContainer(containerID)
		if err != nil {
			errors = append(errors, err.Error())
		}
	}
	if len(errors) > 0 {
		m.message.AddMessage("Error while stopping some containers", ui.ErrorMessage)
		m.selectedContainers.Clear()
		return m, m.message.ClearMessage(errorDuration)
	}
	m.message.AddMessage("Containers stopped", ui.SuccessMessage)
	m.selectedContainers.Clear()
	return m, m.message.ClearMessage(successDuration)
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
	err := m.dockerClient.DeleteContainer(m.table.SelectedRow()[ContainerID])
	if err != nil {
		m.message.AddMessage(fmt.Sprintf("Error while deleting container: %s", strings.Split(err.Error(), ":")[ContainerName]), ui.ErrorMessage)
		return m, m.message.ClearMessage(errorDuration)
	}
	m.message.AddMessage(fmt.Sprintf("Container %s deleted", m.table.SelectedRow()[ContainerName]), ui.SuccessMessage)
	return m, m.message.ClearMessage(successDuration)
}

func (m containerModel) DeleteContainers() (tea.Model, tea.Cmd) {
	errors := make([]string, 0)
	for _, containerID := range m.selectedContainers.ToSlice() {
		err := m.dockerClient.DeleteContainer(containerID)
		if err != nil {
			errors = append(errors, err.Error())
		}
	}
	if len(errors) > 0 {
		m.message.AddMessage("Error while deleting some containers", ui.ErrorMessage)
		m.selectedContainers.Clear()
		return m, m.message.ClearMessage(errorDuration)
	}
	m.message.AddMessage("Containers deleted", ui.SuccessMessage)
	m.selectedContainers.Clear()
	return m, m.message.ClearMessage(successDuration)
}
