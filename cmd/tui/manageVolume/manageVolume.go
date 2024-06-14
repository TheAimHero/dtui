package managevolume

import (
	"github.com/TheAimHero/dtui/internal/docker"
	"github.com/TheAimHero/dtui/internal/ui/message"
	"github.com/TheAimHero/dtui/internal/utils"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	mapset "github.com/deckarep/golang-set/v2"
)

type VolumeModel struct {
	SelectedVolumes mapset.Set[string]
	Help            help.Model
	Keys            keyMap
	DockerClient    docker.DockerClient
	Message         message.Message
	Table           table.Model
}

func (m VolumeModel) Init() tea.Cmd {
	return tea.Batch(utils.TickCommand())
}

func NewModel(dockerClient docker.DockerClient) VolumeModel {
	err := dockerClient.FetchVolumes()
	help := getHelpSection()
	m := VolumeModel{
		DockerClient:    dockerClient,
		Help:            help,
		Message:         message.Message{},
		Keys:            keys,
		SelectedVolumes: mapset.NewSet[string](),
	}
	m.Table = m.getTable()
	if err != nil {
		m.Message.AddMessage("Error while fetching containers", message.ErrorMessage)
		m.Message.ClearMessage(message.SuccessDuration)
	}
	return m
}
