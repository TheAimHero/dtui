package managecontianer

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/TheAimHero/dtui/internal/docker"
	"github.com/TheAimHero/dtui/internal/ui"
)

type containerModel struct {
	help               help.Model
	message            ui.Message
	keys               keyMap
	dockerClient       docker.DockerClient
	selectedContainers []string
	table              table.Model
}

func (m containerModel) Init() tea.Cmd {
	return tickCommand()
}

func getTable(containers docker.Containers, selectedRows []string) table.Model {
	tableColumns := getTableColumns()
	tableRows := getTableRows(containers, selectedRows)
	return ui.NewTable(tableColumns, tableRows)
}

func NewModel(dockerClient docker.DockerClient) tea.Model {
	err := dockerClient.FetchContainers()
	m := containerModel{
		dockerClient:       dockerClient,
		table:              getTable(dockerClient.Containers, []string{}),
		help:               getHelpSection(),
		selectedContainers: make([]string, 0),
		message:            ui.Message{},
		keys:               keys,
	}
	if err != nil {
		m.message.AddMessage("Error while fetching containers", ui.ErrorMessage)
		m.message.ClearMessage(successDuration)
	}
	return m
}
