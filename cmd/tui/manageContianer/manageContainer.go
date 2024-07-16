package managecontianer

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	mapset "github.com/deckarep/golang-set/v2"

	"github.com/TheAimHero/dtui/internal/docker"
	"github.com/TheAimHero/dtui/internal/ui/message"
	"github.com/TheAimHero/dtui/internal/ui/prompt"
	"github.com/TheAimHero/dtui/internal/utils"
)

type ContainerModel struct {
	Help               help.Model
	SelectedContainers mapset.Set[string]
	InProcess          mapset.Set[string]
	Keys               keyMap
	DockerClient       docker.DockerClient
	Confirmation       prompt.Model
	Message            message.Message
	Input              textinput.Model
	Spinner            spinner.Model
	Table              table.Model
}

func (m ContainerModel) Init() tea.Cmd {
	return tea.Batch(utils.TickCommand(), m.Spinner.Tick)
}

func NewModel(dockerClient docker.DockerClient) ContainerModel {
	err := dockerClient.FetchContainers()
	spinner := getSpinner()
	help := getHelpSection()
	input := getInput()
	m := ContainerModel{
		DockerClient:       dockerClient,
		Help:               help,
		Spinner:            spinner,
		SelectedContainers: mapset.NewSet[string](),
		InProcess:          mapset.NewSet[string](),
		Message:            message.Message{},
		Keys:               keys,
		Input:              input,
	}
	m.Table = m.getTable()
	if err != nil {
		m.Message.AddMessage("Error while fetching containers", message.ErrorMessage)
		m.Message.ClearMessage(message.SuccessDuration)
	}
	return m
}
