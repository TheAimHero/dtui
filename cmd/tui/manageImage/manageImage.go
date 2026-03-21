package manageimage

import (
	"sync"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	mapset "github.com/deckarep/golang-set/v2"

	"github.com/TheAimHero/dtui/internal/docker"
	"github.com/TheAimHero/dtui/internal/ui/message"
	"github.com/TheAimHero/dtui/internal/ui/prompt"
	"github.com/TheAimHero/dtui/internal/ui/styles"
	ui_table "github.com/TheAimHero/dtui/internal/ui/table"
	"github.com/TheAimHero/dtui/internal/utils"
)

type ImageModel struct {
	SelectedImages mapset.Set[string]
	PullProgress   *sync.Map
	InProgress     mapset.Set[string]
	PullSpinner    spinner.Model
	LoadingSpinner spinner.Model
	Help           help.Model
	Keys           keyMap
	DockerClient   docker.DockerClient
	Message        message.Message
	Input          textinput.Model
	Conformation   prompt.Model
	Table          table.Model
	Width          int
	Height         int
}

type PullProgressInfo = docker.PullProgressInfo

func (m ImageModel) Init() tea.Cmd {
	var (
		cmds []tea.Cmd
	)
	cmds = append(cmds, utils.TickCommand(), m.PullSpinner.Tick, m.LoadingSpinner.Tick)
	return tea.Batch(cmds...)
}

func getTable(images docker.Images, selectedImages mapset.Set[string], inProcesss mapset.Set[string], spinner spinner.Model, width int) table.Model {
	tableColumns := getTableColumns(width)
	tableRows := getTableRows(images, selectedImages, inProcesss, spinner)
	return ui_table.NewTable(tableColumns, tableRows)
}

func NewModel(dockerClient docker.DockerClient) (ImageModel, error) {
	err := dockerClient.FetchImages()
	m := ImageModel{
		DockerClient:   dockerClient,
		Help:           getHelpSection(),
		PullSpinner:    getSpinner(spinner.Dot),
		LoadingSpinner: getSpinner(spinner.Points),
		SelectedImages: mapset.NewSet[string](),
		InProgress:     mapset.NewSet[string](),
		PullProgress:   &sync.Map{},
		Keys:           keys,
		Width:          80,
		Height:         40,
	}
	m.Table = getTable(m.DockerClient.Images, m.SelectedImages, m.InProgress, m.PullSpinner, m.Width)
	if err != nil {
		return m, styles.ErrorMessage(err.Error())
	}
	return m, nil
}
