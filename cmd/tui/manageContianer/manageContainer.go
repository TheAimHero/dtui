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
	SelectedContainers mapset.Set[string]
	InProcess         mapset.Set[string]
	Help               help.Model
	Keys               keyMap
	DockerClient       docker.DockerClient
	Message            message.Message
	Table              table.Model
	Spinner            spinner.Model
}

func (m ContainerModel) Init() tea.Cmd {
	return tea.Batch(utils.TickCommand(), m.Spinner.Tick)
}

func NewModel(dockerClient docker.DockerClient) ContainerModel {
	err := dockerClient.FetchContainers()
	spinner := getSpinner()
	help := getHelpSection()
	m := ContainerModel{
		DockerClient:       dockerClient,
		Help:               help,
		Spinner:            spinner,
		SelectedContainers: mapset.NewSet[string](),
		InProcess:         mapset.NewSet[string](),
		Message:            message.Message{},
		Keys:               keys,
	}
	m.Table = m.getTable()
	if err != nil {
		m.Message.AddMessage("Error while fetching containers", message.ErrorMessage)
		m.Message.ClearMessage(message.SuccessDuration)
	}
	return m
}
