package logs

import (
	"io"

	"github.com/TheAimHero/dtui/internal/docker"
	"github.com/TheAimHero/dtui/internal/ui"
	"github.com/TheAimHero/dtui/internal/utils"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

type logModel struct {
	stream        io.ReadCloser
	sub           chan responseMsg
	help          help.Model
	title         string
	viewport      viewport.Model
	keys          keyMap
	dockerClient  docker.DockerClient
	text          []string
	message       ui.Message
	table         table.Model
}

func (m logModel) Init() tea.Cmd {
	var cmd tea.Cmd
	m, cmd = m.GetLogs()
	return tea.Batch(
		listenForActivity(m.sub, m.stream),
		waitForActivity(m.sub),
		utils.TickCommand(),
		cmd,
	)
}

func getTable(containers docker.Containers) table.Model {
	tableColumns := getTableColumns()
	tableRows := getTableRows(containers)
	return ui.NewTable(tableColumns, tableRows)
}

func getViewPort() viewport.Model {
	vp := viewport.New(physicalWidth-10, 10)
	return vp
}

func NewModel(dockerClient docker.DockerClient) tea.Model {
	dockerClient.FetchContainers()
	viewport := getViewPort()
	table := getTable(dockerClient.Containers)
	row := table.SelectedRow()
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
	if row != nil {
		m.title = row[ContainerName]
	}
	return m
}
