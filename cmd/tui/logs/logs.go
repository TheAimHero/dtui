package logs

import (
	"io"

	"github.com/TheAimHero/dtui/internal/docker"
	ui_message "github.com/TheAimHero/dtui/internal/ui/message"
	"github.com/TheAimHero/dtui/internal/utils"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
)

type LogModel struct {
	Stream       io.ReadCloser
	Sub          chan utils.ResponseMsg
	Help         help.Model
	Title        string
	Keys         keyMap
	DockerClient docker.DockerClient
	Text         []string
	Message      ui_message.Message
	Table        table.Model
}

func (m LogModel) Init() tea.Cmd {
	var cmds []tea.Cmd
	cmds = append(cmds, utils.TickCommand())
	return tea.Batch(cmds...)
}

func NewModel(dockerClient docker.DockerClient) LogModel {
	err := dockerClient.FetchContainers()
	table := getTable(dockerClient.Containers)
	help := getHelpSection()
	m := LogModel{
		DockerClient: dockerClient,
		Table:        table,
		Sub:          make(chan utils.ResponseMsg),
		Text:         []string{},
		Help:         help,
		Keys:         keys,
	}
	if err != nil {
		m.Message.AddMessage("Error while fetching containers", ui_message.ErrorMessage)
	}
	return m
}
