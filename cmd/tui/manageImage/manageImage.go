package manageimage

import (
	"github.com/TheAimHero/dtui/internal/docker"
	"github.com/TheAimHero/dtui/internal/ui"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
)

type imageModel struct {
	help         help.Model
	message      ui.Message
	keys         keyMap
	dockerClient docker.DockerClient
	table        table.Model
}

func (m imageModel) Init() tea.Cmd {
	return nil
}

func getTable(images docker.Images) table.Model {
	tableColumns := getTableColumns()
	tableRows := getTableRows(images)
	return ui.NewTable(tableColumns, tableRows)
}

func NewModel(dockerClient docker.DockerClient) tea.Model {
	err := dockerClient.FetchImages()
	m := imageModel{
		dockerClient: dockerClient,
		table:        getTable(dockerClient.Images),
		help:         getHelpSection(),
		keys:         keys,
	}
	if err != nil {
		m.message.AddMessage("Error while fetching images", "error")
	}
	return m
}
