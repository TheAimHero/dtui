package manageimage

import (
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	mapset "github.com/deckarep/golang-set/v2"

	"github.com/TheAimHero/dtui/internal/docker"
	"github.com/TheAimHero/dtui/internal/ui/message"
	ui_table "github.com/TheAimHero/dtui/internal/ui/table"
	"github.com/TheAimHero/dtui/internal/utils"
)

type ImageModel struct {
	selectedImages mapset.Set[string]
	sub            chan utils.ResponseMsg
	text           []string
	viewport       viewport.Model
	help           help.Model
	keys           keyMap
	dockerClient   docker.DockerClient
	message        message.Message
	Input          textinput.Model
	Table          table.Model
}

func (m ImageModel) Init() tea.Cmd {
	var (
		cmds []tea.Cmd
	)
	cmds = append(cmds, utils.ResponseToStream(m.sub), utils.TickCommand())
	return tea.Batch(cmds...)
}

func getInput() textinput.Model {
	ip := textinput.New()
	ip.Placeholder = "Image Name"
	return ip
}

func getTable(images docker.Images, selectedImages mapset.Set[string]) table.Model {
	tableColumns := getTableColumns()
	tableRows := getTableRows(images, selectedImages)
	return ui_table.NewTable(tableColumns, tableRows)
}

func NewModel(dockerClient docker.DockerClient) ImageModel {
	err := dockerClient.FetchImages()
	m := ImageModel{
		dockerClient:   dockerClient,
		Table:          getTable(dockerClient.Images, mapset.NewSet[string]()),
		help:           getHelpSection(),
		viewport:       getViewPort(),
		sub:            make(chan utils.ResponseMsg),
		selectedImages: mapset.NewSet[string](),
		keys:           keys,
	}
	if err != nil {
		m.message.AddMessage("Error while fetching images", message.ErrorMessage)
		m.message.ClearMessage(2 * time.Second)
	}
	return m
}
