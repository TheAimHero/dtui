package manageimage

import (
	"time"

	"github.com/TheAimHero/dtui/internal/docker"
	"github.com/TheAimHero/dtui/internal/size"
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
	width := (physicalWidth / 4) - 4
	tableColumns := []table.Column{
		{Title: "ID", Width: width},
		{Title: "Name", Width: width},
		{Title: "Created", Width: width},
		{Title: "Size", Width: width},
	}
	tableRows := make([]table.Row, len(images))
	for i, image := range images {
		tableRows[i] = table.Row{
			image.ID,
			image.RepoTags[0],
			time.Unix(image.Created, 0).Format("02/01/2006 15:04 MST"),
			size.GetSize(image.Size),
		}
	}
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
