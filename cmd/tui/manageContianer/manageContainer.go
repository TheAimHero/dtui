package managecontianer

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/TheAimHero/dtui/internal/docker"
	"github.com/TheAimHero/dtui/internal/ui"
)

type containerModel struct {
	help         help.Model
	message      ui.Message
	keys         keyMap
	dockerClient docker.DockerClient
	table        table.Model
}

func (m containerModel) Init() tea.Cmd {
	return tickCommand()
}

func getTable(containers docker.Containers) table.Model {
	tableColumns := getTableColumns()
	tableRows := getTableRows(containers)
	return ui.NewTable(tableColumns, tableRows)
}

func NewModel(dockerClient docker.DockerClient) tea.Model {
	err := dockerClient.FetchContainers()
	m := containerModel{
		dockerClient: dockerClient,
		table:        getTable(dockerClient.Containers),
		help:         getHelpSection(),
		keys:         keys,
	}
	if err != nil {
		m.message.AddMessage("Error while fetching containers", "error")
		m.message.ClearMessage(successDuration)
	}
	return m
}
