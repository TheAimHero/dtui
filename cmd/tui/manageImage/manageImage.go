package manageimage

import (
	"sync"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	mapset "github.com/deckarep/golang-set/v2"

	"github.com/TheAimHero/dtui/internal/docker"
	"github.com/TheAimHero/dtui/internal/ui/components"
	"github.com/TheAimHero/dtui/internal/ui/message"
	"github.com/TheAimHero/dtui/internal/ui/prompt"
	"github.com/TheAimHero/dtui/internal/utils"
)

type ImageModel struct {
	*components.BaseModel
	ImageSvc       docker.ImageService
	Images         docker.Images
	SelectedImages mapset.Set[string]
	InProgress     mapset.Set[string]
	PullProgress   *sync.Map
	PullSpinner    spinner.Model
	Confirmation   prompt.Model
	Message        message.Message
	Keys           keyMap
	LoadingSpinner spinner.Model
}

type PullProgressInfo = docker.PullProgressInfo

func (m ImageModel) Init() tea.Cmd {
	var (
		cmds []tea.Cmd
	)
	cmds = append(cmds, utils.TickCommand(), m.PullSpinner.Tick, m.LoadingSpinner.Tick)
	return tea.Batch(cmds...)
}

func getTable(images docker.Images, selectedImages mapset.Set[string], inProcess mapset.Set[string], spinner spinner.Model, width int) table.Model {
	tableColumns := getTableColumns(width)
	tableRows := getTableRows(images, selectedImages, inProcess, spinner)
	return components.NewStandardTable(tableColumns, tableRows)
}

func NewModel(imageSvc docker.ImageService) (ImageModel, error) {
	images, err := imageSvc.FetchImages()
	base := components.NewBaseModel(80, 40)
	m := ImageModel{
		BaseModel:      &base,
		ImageSvc:       imageSvc,
		Images:         images,
		PullSpinner:    getSpinner(spinner.Dot),
		LoadingSpinner: getSpinner(spinner.Points),
		SelectedImages: mapset.NewSet[string](),
		InProgress:     mapset.NewSet[string](),
		PullProgress:   &sync.Map{},
		Keys:           keys,
	}
	m.Table = getTable(m.Images, m.SelectedImages, m.InProgress, m.PullSpinner, m.BaseModel.Width)
	if err != nil {
		return m, components.ErrorMessage(err.Error())
	}
	return m, nil
}
