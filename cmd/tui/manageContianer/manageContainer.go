package managecontianer

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	mapset "github.com/deckarep/golang-set/v2"

	"github.com/TheAimHero/dtui/internal/docker"
	"github.com/TheAimHero/dtui/internal/ui"
)

type containerModel struct {
	help               help.Model
	message            ui.Message
	keys               keyMap
	dockerClient       docker.DockerClient
	selectedContainers mapset.Set[string]
	table              table.Model
}

func (m containerModel) Init() tea.Cmd {
	return tickCommand()
}

func getTable(containers docker.Containers, selectedRows mapset.Set[string]) table.Model {
	tableColumns := getTableColumns()
	tableRows := getTableRows(containers, selectedRows)
	return ui.NewTable(tableColumns, tableRows)
}

func NewModel(dockerClient docker.DockerClient) tea.Model {
	err := dockerClient.FetchContainers()
	m := containerModel{
		dockerClient:       dockerClient,
		table:              getTable(dockerClient.Containers, mapset.NewSet[string]()),
		help:               getHelpSection(),
		selectedContainers: mapset.NewSet[string](),
		message:            ui.Message{},
		keys:               keys,
	}
	if err != nil {
		m.message.AddMessage("Error while fetching containers", ui.ErrorMessage)
		m.message.ClearMessage(successDuration)
	}
	return m
}
