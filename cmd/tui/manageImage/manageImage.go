package manageimage

import (
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	mapset "github.com/deckarep/golang-set/v2"

	"github.com/TheAimHero/dtui/internal/docker"
	"github.com/TheAimHero/dtui/internal/ui/message"
	ui_table "github.com/TheAimHero/dtui/internal/ui/table"
	"github.com/TheAimHero/dtui/internal/utils"
)

type ImageModel struct {
	SelectedImages mapset.Set[string]
	PullProgress   mapset.Set[string]
	InProgress     mapset.Set[string]
	PullSpinner    spinner.Model
	LoadingSpinner spinner.Model
	Help           help.Model
	Keys           keyMap
	DockerClient   docker.DockerClient
	Message        message.Message
	Input          textinput.Model
	Table          table.Model
}

func (m ImageModel) Init() tea.Cmd {
	var (
		cmds []tea.Cmd
	)
	cmds = append(cmds, utils.TickCommand(), m.PullSpinner.Tick, m.LoadingSpinner.Tick)
	return tea.Batch(cmds...)
}

func getTable(images docker.Images, selectedImages mapset.Set[string], inProcesss mapset.Set[string], spinner spinner.Model) table.Model {
	tableColumns := getTableColumns()
	tableRows := getTableRows(images, selectedImages, inProcesss, spinner)
	return ui_table.NewTable(tableColumns, tableRows)
}

func NewModel(dockerClient docker.DockerClient) ImageModel {
	err := dockerClient.FetchImages()
	m := ImageModel{
		DockerClient:   dockerClient,
		Help:           getHelpSection(),
		PullSpinner:    getSpinner(spinner.Dot),
		LoadingSpinner: getSpinner(spinner.Points),
		SelectedImages: mapset.NewSet[string](),
		InProgress:     mapset.NewSet[string](),
		PullProgress:   mapset.NewSet[string](),
		Keys:           keys,
	}
	m.Table = getTable(m.DockerClient.Images, m.SelectedImages, m.InProgress, m.PullSpinner)
	if err != nil {
		m.Message.AddMessage("Error while fetching images", message.ErrorMessage)
		m.Message.ClearMessage(2 * time.Second)
	}
	return m
}
