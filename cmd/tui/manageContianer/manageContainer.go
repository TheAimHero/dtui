package managecontianer

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	mapset "github.com/deckarep/golang-set/v2"

	"github.com/TheAimHero/dtui/internal/docker"
	"github.com/TheAimHero/dtui/internal/ui/message"
	"github.com/TheAimHero/dtui/internal/utils"
)

type containerModel struct {
	selectedContainers mapset.Set[string]
	log                tea.Model
	help               help.Model
	keys               keyMap
	dockerClient       docker.DockerClient
	message            message.Message
	table              table.Model
}

func (m containerModel) Init() tea.Cmd {
	return tea.Batch(utils.TickCommand())
}

func NewModel(dockerClient docker.DockerClient) tea.Model {
	err := dockerClient.FetchContainers()
	table := getTable(dockerClient.Containers, mapset.NewSet[string]())
	help := getHelpSection()
	m := containerModel{
		dockerClient:       dockerClient,
		table:              table,
		help:               help,
		selectedContainers: mapset.NewSet[string](),
		message:            message.Message{},
		keys:               keys,
	}
	if err != nil {
		m.message.AddMessage("Error while fetching containers", message.ErrorMessage)
		m.message.ClearMessage(message.SuccessDuration)
	}
	return m
}
