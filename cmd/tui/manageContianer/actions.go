package managecontianer

import (
	"fmt"
	"strings"

	"github.com/TheAimHero/dtui/internal/ui/message"
	tea "github.com/charmbracelet/bubbletea"
	mapset "github.com/deckarep/golang-set/v2"
)

const (
	ContainerSelected = iota
	ContainerInProcess
	ContainerID
	ContainerName
	ContainerImage
	ContainerStatus
)

func (m *ContainerModel) ClearSelectedContainers() {
	m.SelectedContainers.Clear()
}

func (m ContainerModel) StartContainer() (ContainerModel, tea.Cmd) {
	startMsg := message.Message{}
	row := m.Table.SelectedRow()
	if row == nil {
		m.Message.AddMessage("No container selected", message.InfoMessage)
		return m, m.Message.ClearMessage(message.InfoDuration)
	}
	m.InProcess.Add(row[ContainerID])
	return m, func() tea.Msg {
		err := m.DockerClient.StartContainer(row[ContainerID])
		if err != nil {
			startMsg.AddMessage(fmt.Sprintf("Error while starting container: %s", strings.Split(err.Error(), ":")[1]), message.ErrorMessage)
			return startMsg
		}
		startMsg.AddMessage(fmt.Sprintf("Container %s started", m.Table.SelectedRow()[ContainerName]), message.SuccessMessage)
		m.InProcess.Remove(row[ContainerID])
		return startMsg
	}
}

func (m ContainerModel) StopContainer() (ContainerModel, tea.Cmd) {
	stopMsg := message.Message{}
	row := m.Table.SelectedRow()
	if row == nil {
		m.Message.AddMessage("No container selected", message.InfoMessage)
		return m, m.Message.ClearMessage(message.InfoDuration)
	}
	m.InProcess.Add(row[ContainerID])
	return m, func() tea.Msg {
		err := m.DockerClient.StopContainer(m.Table.SelectedRow()[ContainerID])
		if err != nil {
			stopMsg.AddMessage(fmt.Sprintf("Error while stopping container: %s", strings.Split(err.Error(), ":")[ContainerName]), message.ErrorMessage)
			return stopMsg
		}
		stopMsg.AddMessage(fmt.Sprintf("Container %s stopped", m.Table.SelectedRow()[ContainerName]), message.SuccessMessage)
		m.InProcess.Remove(row[ContainerID])
		return stopMsg
	}
}

func (m ContainerModel) StartContainers() (ContainerModel, tea.Cmd) {
	startMsg := message.Message{}
	defer m.ClearSelectedContainers()
	errors := make([]string, 0)
	if m.SelectedContainers.Cardinality() == 0 {
		m.Message.AddMessage("No containers selected", message.InfoMessage)
		return m, m.Message.ClearMessage(message.InfoDuration)
	}
	tableContainers := mapset.NewSet[string]()
	for _, row := range m.Table.Rows() {
		tableContainers.Add(row[ContainerID])
	}
	toStart := tableContainers.Intersect(m.SelectedContainers).ToSlice()
  if len(toStart) == 0 {
    m.Message.AddMessage("No containers selected", message.InfoMessage)
    return m, m.Message.ClearMessage(message.InfoDuration)
  }
	return m, func() tea.Msg {
		for _, containerID := range toStart {
			go func(containerID string) {
				m.InProcess.Add(containerID)
				err := m.DockerClient.StartContainer(containerID)
				if err != nil {
					errors = append(errors, err.Error())
				}
				m.InProcess.Remove(containerID)
			}(containerID)
		}
		if len(errors) > 0 {
			startMsg.AddMessage("Error while starting some containers", message.ErrorMessage)
			m.SelectedContainers.Clear()
			return startMsg
		}
		startMsg.AddMessage("Containers started", message.SuccessMessage)
		m.SelectedContainers.Clear()
		return startMsg
	}
}

func (m ContainerModel) StopContainers() (ContainerModel, tea.Cmd) {
	errors := make([]string, 0)
	selectedContainers := m.SelectedContainers.ToSlice()
	stopMsg := message.Message{}
	if m.SelectedContainers.Cardinality() == 0 {
		m.Message.AddMessage("No containers selected", message.InfoMessage)
		return m, m.Message.ClearMessage(message.InfoDuration)
	}
  tableContainers := mapset.NewSet[string]()
  for _, row := range m.Table.Rows() {
    tableContainers.Add(row[ContainerID])
  }
  toStop := tableContainers.Intersect(m.SelectedContainers).ToSlice()
  if len(toStop) == 0 {
    m.Message.AddMessage("No containers selected", message.InfoMessage)
    return m, m.Message.ClearMessage(message.InfoDuration)
  }
	defer m.ClearSelectedContainers()
	return m, func() tea.Msg {
		for _, containerID := range selectedContainers {
			go func(containerID string) {
				m.InProcess.Add(containerID)
				err := m.DockerClient.StopContainer(containerID)
				if err != nil {
					errors = append(errors, err.Error())
				}
				m.InProcess.Remove(containerID)
			}(containerID)
		}
		if len(errors) > 0 {
			stopMsg.AddMessage("Error while stopping some containers", message.ErrorMessage)
			m.SelectedContainers.Clear()
			return stopMsg
		}
		stopMsg.AddMessage("Containers stopped", message.SuccessMessage)
		m.SelectedContainers.Clear()
		return stopMsg
	}
}

func (m ContainerModel) SelectContainers() (ContainerModel, tea.Cmd) {
	row := m.Table.SelectedRow()
	if row == nil {
		m.Message.AddMessage("No containers to select", message.InfoMessage)
		return m, m.Message.ClearMessage(message.InfoDuration)
	}
	containerID := row[ContainerID]
	if m.SelectedContainers.Contains(containerID) {
		m.SelectedContainers.Remove(containerID)
	} else {
		m.SelectedContainers.Add(containerID)
	}
	m.Table.MoveDown(1)
	return m, nil
}

func (m ContainerModel) SelectAllContainers() (ContainerModel, tea.Cmd) {
	var allIDs []string
	for _, row := range m.Table.Rows() {
		allIDs = append(allIDs, row[ContainerID])
	}
	if m.SelectedContainers.Cardinality() == len(m.Table.Rows()) {
		m.SelectedContainers.Clear()
	} else {
		m.SelectedContainers.Clear()
		m.SelectedContainers.Append(allIDs...)
	}
	return m, nil
}

func (m ContainerModel) DeleteContainer() (ContainerModel, tea.Cmd) {
	deleteMsg := message.Message{}
	row := m.Table.SelectedRow()
	if row == nil {
		m.Message.AddMessage("No container selected", message.InfoMessage)
		return m, m.Message.ClearMessage(message.InfoDuration)
	}
	return m, func() tea.Msg {
		m.InProcess.Add(row[ContainerID])
		err := m.DockerClient.DeleteContainer(row[ContainerID])
		if err != nil {
			deleteMsg.AddMessage(fmt.Sprintf("Error while deleting container: %s", strings.Split(err.Error(), ":")[ContainerName]), message.ErrorMessage)
			return deleteMsg
		}
		deleteMsg.AddMessage(fmt.Sprintf("Container %s deleted", m.Table.SelectedRow()[ContainerName]), message.SuccessMessage)
		m.InProcess.Remove(row[ContainerID])
		return deleteMsg
	}
}

func (m ContainerModel) DeleteContainers() (ContainerModel, tea.Cmd) {
	defer m.ClearSelectedContainers()
	tableContainers := mapset.NewSet[string]()
	rows := m.Table.Rows()
	for _, row := range rows {
		tableContainers.Add(row[ContainerID])
	}
	if m.SelectedContainers.Cardinality() == 0 {
		m.Message.AddMessage("No containers selected", message.InfoMessage)
		return m, m.Message.ClearMessage(message.InfoDuration)
	}
	toDelete := tableContainers.Intersect(m.SelectedContainers).ToSlice()
  if len(toDelete) == 0 {
    m.Message.AddMessage("No containers selected", message.InfoMessage)
    return m, m.Message.ClearMessage(message.InfoDuration)
  }
	deleteMsg := message.Message{}
	errors := make([]string, 0)
	return m, func() tea.Msg {
		for _, containerID := range toDelete {
			go func(containerID string) {
				m.InProcess.Add(containerID)
				err := m.DockerClient.DeleteContainer(containerID)
				if err != nil {
					errors = append(errors, err.Error())
				}
				m.InProcess.Remove(containerID)
			}(containerID)
		}
		if len(errors) > 0 {
			deleteMsg.AddMessage("Error while deleting some containers", message.ErrorMessage)
			m.SelectedContainers.Clear()
			return deleteMsg

		}
		m.SelectedContainers.Clear()
		deleteMsg.AddMessage("Containers deleted", message.SuccessMessage)
		return deleteMsg
	}
}
