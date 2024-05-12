package managecontianer

import (
	"strings"

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
	width := (physicalWidth / 3) - 4
	tableColumns := []table.Column{
		{Title: "ID", Width: width},
		{Title: "Name", Width: width},
		{Title: "State", Width: width},
	}
	tableRows := make([]table.Row, len(containers))
	for i, container := range containers {
		tableRows[i] = table.Row{
			container.ID,
			strings.Split(container.Names[0], "/")[1],
			strings.ToUpper(string(container.State[0])) + string(container.State[1:]),
		}
	}
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
	}
	return m
}
