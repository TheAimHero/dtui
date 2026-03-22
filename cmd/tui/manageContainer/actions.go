package managecontainer

import (
	"os/exec"
	"sync"

	"github.com/TheAimHero/dtui/internal/docker"
	"github.com/TheAimHero/dtui/internal/ui/message"
	tea "github.com/charmbracelet/bubbletea"
	mapset "github.com/deckarep/golang-set/v2"
)

// ContainerActionResult is sent when a container action completes.
// The Update method handles this message to update InProcess state.
type ContainerActionResult struct {
	ContainerID string
	Action      ActionType
	Err         error
}

// ContainerBatchResult aggregates all results from a batch operation.
type ContainerBatchResult struct {
	Results []ContainerActionResult
	Action  ActionType
}

const (
	ContainerSelected = iota
	ContainerInProcess
	ContainerID
	ContainerName
	ContainerImage
	ContainerStatus
)

func isValidContainerID(id string) bool {
	return len(id) > 0 && len(id) <= 128
}

func (m *ContainerModel) ClearSelectedContainers() {
	m.SelectedContainers.Clear()
}

func startContainer(m ContainerModel) (ContainerModel, tea.Cmd) {
	row := m.Table.SelectedRow()
	if row == nil {
		m.Message.AddMessage("No container selected", message.InfoMessage)
		return m, m.Message.ClearMessage(message.InfoDuration)
	}
	containerID := row[ContainerID]
	m.InProcess.Add(containerID)
	return m, func() tea.Msg {
		err := m.ContainerSvc.StartContainer(containerID)
		return ContainerActionResult{
			ContainerID: containerID,
			Action:      ActionStartContainer,
			Err:         err,
		}
	}
}

func (m ContainerModel) StartContainers() (ContainerModel, tea.Cmd) {
	defer m.ClearSelectedContainers()
	if m.SelectedContainers.Cardinality() == 0 {
		return startContainer(m)
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
	// Add all to InProcess synchronously before returning
	for _, containerID := range toStart {
		m.InProcess.Add(containerID)
	}
	return m, func() tea.Msg {
		results := make([]ContainerActionResult, len(toStart))
		var wg sync.WaitGroup
		for i, containerID := range toStart {
			wg.Add(1)
			go func(i int, containerID string) {
				defer wg.Done()
				err := m.ContainerSvc.StartContainer(containerID)
				results[i] = ContainerActionResult{
					ContainerID: containerID,
					Action:      ActionStartContainer,
					Err:         err,
				}
			}(i, containerID)
		}
		wg.Wait()
		return ContainerBatchResult{
			Results: results,
			Action:  ActionStartContainer,
		}
	}
}

func stopContainer(m ContainerModel) (ContainerModel, tea.Cmd) {
	row := m.Table.SelectedRow()
	if row == nil {
		m.Message.AddMessage("No container selected", message.InfoMessage)
		return m, m.Message.ClearMessage(message.InfoDuration)
	}
	containerID := row[ContainerID]
	m.InProcess.Add(containerID)
	return m, func() tea.Msg {
		err := m.ContainerSvc.StopContainer(containerID)
		return ContainerActionResult{
			ContainerID: containerID,
			Action:      ActionStopContainer,
			Err:         err,
		}
	}
}

func (m ContainerModel) StopContainers() (ContainerModel, tea.Cmd) {
	if m.SelectedContainers.Cardinality() == 0 {
		return stopContainer(m)
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
	// Add all to InProcess synchronously before returning
	for _, containerID := range toStop {
		m.InProcess.Add(containerID)
	}
	return m, func() tea.Msg {
		results := make([]ContainerActionResult, len(toStop))
		var wg sync.WaitGroup
		for i, containerID := range toStop {
			wg.Add(1)
			go func(i int, containerID string) {
				defer wg.Done()
				err := m.ContainerSvc.StopContainer(containerID)
				results[i] = ContainerActionResult{
					ContainerID: containerID,
					Action:      ActionStopContainer,
					Err:         err,
				}
			}(i, containerID)
		}
		wg.Wait()
		return ContainerBatchResult{
			Results: results,
			Action:  ActionStopContainer,
		}
	}
}

func deleteContainer(m ContainerModel) (ContainerModel, tea.Cmd) {
	row := m.Table.SelectedRow()
	if row == nil {
		m.Message.AddMessage("No container selected", message.InfoMessage)
		return m, m.Message.ClearMessage(message.InfoDuration)
	}
	containerID := row[ContainerID]
	m.InProcess.Add(containerID)
	return m, func() tea.Msg {
		err := m.ContainerSvc.DeleteContainer(containerID)
		return ContainerActionResult{
			ContainerID: containerID,
			Action:      ActionDeleteContainer,
			Err:         err,
		}
	}
}

func (m ContainerModel) DeleteContainers() (ContainerModel, tea.Cmd) {
	defer m.ClearSelectedContainers()
	if m.SelectedContainers.Cardinality() == 0 {
		return deleteContainer(m)
	}
	rows := m.Table.Rows()
	tableContainers := mapset.NewSet[string]()
	for _, row := range rows {
		tableContainers.Add(row[ContainerID])
	}
	toDelete := tableContainers.Intersect(m.SelectedContainers).ToSlice()
	if len(toDelete) == 0 {
		m.Message.AddMessage("No containers selected", message.InfoMessage)
		return m, m.Message.ClearMessage(message.InfoDuration)
	}
	// Add all to InProcess synchronously before returning
	for _, containerID := range toDelete {
		m.InProcess.Add(containerID)
	}
	return m, func() tea.Msg {
		results := make([]ContainerActionResult, len(toDelete))
		var wg sync.WaitGroup
		for i, containerID := range toDelete {
			wg.Add(1)
			go func(i int, containerID string) {
				defer wg.Done()
				err := m.ContainerSvc.DeleteContainer(containerID)
				results[i] = ContainerActionResult{
					ContainerID: containerID,
					Action:      ActionDeleteContainer,
					Err:         err,
				}
			}(i, containerID)
		}
		wg.Wait()
		return ContainerBatchResult{
			Results: results,
			Action:  ActionDeleteContainer,
		}
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

func (m ContainerModel) ExecContainer() (ContainerModel, tea.Cmd) {
	row := m.Table.SelectedRow()
	if row == nil {
		m.Message.AddMessage("No container selected", message.InfoMessage)
		return m, m.Message.ClearMessage(message.InfoDuration)
	}
	containerID := row[ContainerID]
	if containerID == "" {
		m.Message.AddMessage("No container selected", message.InfoMessage)
		return m, m.Message.ClearMessage(message.InfoDuration)
	}
	var container docker.Container
	for _, c := range m.Containers {
		if c.ID == containerID {
			container = c
			break
		}
	}
	if container.State != "running" {
		m.Message.AddMessage("Container is not running", message.InfoMessage)
		return m, m.Message.ClearMessage(message.InfoDuration)
	}
	if !isValidContainerID(containerID) {
		m.Message.AddMessage("Invalid container ID", message.ErrorMessage)
		return m, m.Message.ClearMessage(message.ErrorDuration)
	}
	c := exec.Command("docker", "container", "exec", "-it", containerID, "sh")
	return m, tea.ExecProcess(c, func(err error) tea.Msg { return tea.ClearScreen })
}

func (m ContainerModel) ShowLogs() (ContainerModel, tea.Cmd) {
	row := m.Table.SelectedRow()
	if row == nil {
		m.Message.AddMessage("No container selected", message.InfoMessage)
		return m, m.Message.ClearMessage(message.InfoDuration)
	}
	containerID := row[ContainerID]
	if containerID == "" {
		m.Message.AddMessage("No container selected", message.InfoMessage)
		return m, m.Message.ClearMessage(message.InfoDuration)
	}
	if !isValidContainerID(containerID) {
		m.Message.AddMessage("Invalid container ID", message.ErrorMessage)
		return m, m.Message.ClearMessage(message.ErrorDuration)
	}
	c := exec.Command("docker", "logs", containerID, "--follow")
	return m, tea.ExecProcess(c, func(err error) tea.Msg { return tea.ClearScreen })
}
