package logs

import (
	"io"

	"github.com/TheAimHero/dtui/internal/docker"
	ui_message "github.com/TheAimHero/dtui/internal/ui/message"
	"github.com/TheAimHero/dtui/internal/utils"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

type logModel struct {
	stream       io.ReadCloser
	sub          chan responseMsg
	help         help.Model
	title        string
	viewport     viewport.Model
	keys         keyMap
	dockerClient docker.DockerClient
	text         []string
	message      ui_message.Message
	table        table.Model
}

func (m logModel) Init() tea.Cmd {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)
	m, cmd = m.GetLogs()
	cmds = append(cmds, cmd)
	cmds = append(cmds, waitForActivity(m.sub), utils.TickCommand())
	return tea.Batch(cmds...)
}

func NewModel(dockerClient docker.DockerClient) tea.Model {
	err := dockerClient.FetchContainers()
	viewport := getViewPort()
	table := getTable(dockerClient.Containers)
	help := getHelpSection()
	m := logModel{
		dockerClient: dockerClient,
		viewport:     viewport,
		table:        table,
		sub:          make(chan responseMsg),
		text:         []string{},
		help:         help,
		keys:         keys,
	}
	if err != nil {
		m.message.AddMessage("Error while fetching containers", ui_message.ErrorMessage)
	}
	return m
}
