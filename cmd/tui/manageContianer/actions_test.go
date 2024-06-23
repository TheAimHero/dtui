package managecontianer

import (
	"testing"
	"time"

	"github.com/TheAimHero/dtui/internal/docker"
	"github.com/TheAimHero/dtui/internal/docker/mocks"
	tea "github.com/charmbracelet/bubbletea"
)

func Includes(id string, containers docker.Containers) bool {
	for _, c := range containers {
		if c.ID == id {
			return true
		}
	}
	return false
}

func RunAction(action func() (ContainerModel, tea.Cmd)) ContainerModel {
	m, cmd := action()
	if cmd != nil {
		cmd()
	}
	return m
}

func TestContainerModel_ClearSelectedContainers(t *testing.T) {
	m := NewModel(mocks.NewMockDockerClient())
	containers := m.DockerClient.GetContainers()
	for _, c := range containers {
		m.SelectedContainers.Add(c.ID)
	}
	m.ClearSelectedContainers()
	if m.SelectedContainers.Cardinality() != 0 {
		t.Errorf("Expected SelectedContainers to be empty, got %v", m.SelectedContainers)
	}
}

func TestContainerModel_SelectAllContainers(t *testing.T) {
	m := NewModel(mocks.NewMockDockerClient())
	containers := m.DockerClient.GetContainers()
	numContainers := len(containers)
	m.SelectAllContainers()
	if m.SelectedContainers.Cardinality() != 3 {
		t.Errorf("Expected SelectedContainers to be %v, got %v", numContainers, m.SelectedContainers)
	}
}

func TestContainerModel_SelectContainers(t *testing.T) {
	m := NewModel(mocks.NewMockDockerClient())
	id := m.Table.SelectedRow()[ContainerID]
	m = RunAction(m.SelectContainers)
	if m.SelectedContainers.Cardinality() != 1 {
		t.Errorf("Expected SelectedContainers to be 1, got %v", m.SelectedContainers)
	}
	if !Includes(id, m.DockerClient.GetContainers()) {
		t.Errorf("Expected container %v to be selected, got %v", id, m.DockerClient.GetContainers())
	}
}

func TestContainerModel_DeleteContainers(t *testing.T) {
	m := NewModel(mocks.NewMockDockerClient())
	//test delete only the container having the table row selected
	id := m.Table.SelectedRow()[ContainerID]
	prevLen := len(m.DockerClient.GetContainers())
	RunAction(m.DeleteContainers)
	if Includes(id, m.DockerClient.GetContainers()) {
		t.Errorf("Expected container %v to be deleted, got %v", id, m.DockerClient.GetContainers())
	}
	if len(m.DockerClient.GetContainers()) != prevLen-1 {
		t.Errorf("Expected containers to be deleted and len to be %v, got %v", prevLen-1, len(m.DockerClient.GetContainers()))
	}
	// test delete all containers
	m.SelectAllContainers()
	RunAction(m.DeleteContainers)
	// wait for all the go routines to finish
	time.Sleep(time.Second)
	if len(m.DockerClient.GetContainers()) > 0 {
		t.Errorf("Expected containers to be deleted and len to be 0, got %v", len(m.DockerClient.GetContainers()))
	}
}

func TestContainerModel_StartContainers(t *testing.T) {
	m := NewModel(mocks.NewMockDockerClient())
	id := m.Table.SelectedRow()[ContainerID]
	RunAction(m.StartContainers)
	if !Includes(id, m.DockerClient.GetContainers()) {
		t.Errorf("Expected container %v to be started, got %v", id, m.DockerClient.GetContainers())
	}
}
