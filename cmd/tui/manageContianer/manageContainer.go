package managecontianer

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	mapset "github.com/deckarep/golang-set/v2"

	"github.com/TheAimHero/dtui/internal/docker"
	"github.com/TheAimHero/dtui/internal/ui/message"
	"github.com/TheAimHero/dtui/internal/utils"
)

type ContainerModel struct {
	selectedContainers mapset.Set[string]
	inProcesss         mapset.Set[string]
	help               help.Model
	keys               keyMap
	dockerClient       docker.DockerClient
	message            message.Message
	table              table.Model
	spinner            spinner.Model
}

func (m ContainerModel) Init() tea.Cmd {
	return tea.Batch(utils.TickCommand(), m.spinner.Tick)
}

func NewModel(dockerClient docker.DockerClient) ContainerModel {
	err := dockerClient.FetchContainers()
	spinner := getSpinner()
	help := getHelpSection()
	m := ContainerModel{
		dockerClient:       dockerClient,
		help:               help,
		spinner:            spinner,
		selectedContainers: mapset.NewSet[string](),
		inProcesss:         mapset.NewSet[string](),
		message:            message.Message{},
		keys:               keys,
	}
	m.table = m.getTable()
	if err != nil {
		m.message.AddMessage("Error while fetching containers", message.ErrorMessage)
		m.message.ClearMessage(message.SuccessDuration)
	}
	return m
}
