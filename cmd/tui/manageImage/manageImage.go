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

type ImageModel struct {
	SelectedImages mapset.Set[string]
	Help           help.Model
	Keys           keyMap
	DockerClient   docker.DockerClient
	Message        message.Message
	Table          table.Model
}

func (m ImageModel) Init() tea.Cmd {
	var (
		cmds []tea.Cmd
	)
	cmds = append(cmds, utils.TickCommand())
	return tea.Batch(cmds...)
}

func getTable(images docker.Images, selectedImages mapset.Set[string]) table.Model {
	tableColumns := getTableColumns()
	tableRows := getTableRows(images, selectedImages)
	return ui_table.NewTable(tableColumns, tableRows)
}

func NewModel(dockerClient docker.DockerClient) ImageModel {
	err := dockerClient.FetchImages()
	m := ImageModel{
		DockerClient:   dockerClient,
		Table:          getTable(dockerClient.Images, mapset.NewSet[string]()),
		Help:           getHelpSection(),
		SelectedImages: mapset.NewSet[string](),
		Keys:           keys,
	}
	if err != nil {
		m.Message.AddMessage("Error while fetching images", message.ErrorMessage)
		m.Message.ClearMessage(2 * time.Second)
	}
	return m
}
