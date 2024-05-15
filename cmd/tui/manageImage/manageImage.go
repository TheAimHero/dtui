package manageimage

import (
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	mapset "github.com/deckarep/golang-set/v2"

	"github.com/TheAimHero/dtui/internal/docker"
	"github.com/TheAimHero/dtui/internal/ui/message"
	ui_table "github.com/TheAimHero/dtui/internal/ui/table"
	"github.com/TheAimHero/dtui/internal/utils"
)

type imageModel struct {
	selectedImages mapset.Set[string]
	help           help.Model
	keys           keyMap
	dockerClient   docker.DockerClient
	message        message.Message
	table          table.Model
}

func (m imageModel) Init() tea.Cmd {
	return utils.TickCommand()
}

func getTable(images docker.Images, selectedImages mapset.Set[string]) table.Model {
	tableColumns := getTableColumns()
	tableRows := getTableRows(images, selectedImages)
	return ui_table.NewTable(tableColumns, tableRows)
}

func NewModel(dockerClient docker.DockerClient) tea.Model {
	err := dockerClient.FetchImages()
	m := imageModel{
		dockerClient:   dockerClient,
		table:          getTable(dockerClient.Images, mapset.NewSet[string]()),
		help:           getHelpSection(),
		selectedImages: mapset.NewSet[string](),
		keys:           keys,
	}
	if err != nil {
		m.message.AddMessage("Error while fetching images", message.ErrorMessage)
		m.message.ClearMessage(2 * time.Second)
	}
	return m
}
