package logs

import (
	"bufio"
	"io"
	"strings"

	"github.com/TheAimHero/dtui/internal/docker"
	"github.com/TheAimHero/dtui/internal/ui"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	ContainerID = iota
	ContainerName
	ContainerImage
	ContainerStatus
)

type responseMsg string

func getTable(containers docker.Containers) table.Model {
	tableColumns := getTableColumns()
	tableRows := getTableRows(containers)
	table := ui.NewTable(tableColumns, tableRows)
	table.KeyMap.HalfPageDown.Unbind()
	table.KeyMap.HalfPageUp.Unbind()
	table.KeyMap.GotoBottom.Unbind()
	table.KeyMap.GotoTop.Unbind()
	table.KeyMap.PageDown.Unbind()
	table.KeyMap.PageUp.Unbind()
	return table
}

func getViewPort() viewport.Model {
	vp := viewport.New(physicalWidth-10, 10)
	vp.KeyMap.Down.Unbind()
	vp.KeyMap.Up.Unbind()
	vp.KeyMap.PageDown.Unbind()
	vp.KeyMap.PageUp.Unbind()
	return vp
}

func getTableRows(containers docker.Containers) []table.Row {
	tableRows := []table.Row{}
	for _, container := range containers {
		tableRows = append(tableRows, table.Row{
			container.ID,
			strings.Split(container.Names[0], "/")[1],
			container.Image,
			strings.ToUpper(string(container.Status[0])) + string(container.Status[1:]),
		})
	}
	return tableRows
}

func getTableColumns() []table.Column {
	width := ((physicalWidth) / 4) - 4
	return []table.Column{
		{Title: "ID", Width: width},
		{Title: "Name", Width: width},
		{Title: "Image", Width: width},
		{Title: "Status", Width: width},
	}
}

func listenForActivity(sub chan<- responseMsg, stream io.ReadCloser) tea.Cmd {
	if stream == nil {
		return nil
	}
	return func() tea.Msg {
		defer stream.Close()
		scanner := bufio.NewScanner(stream)
		for scanner.Scan() {
			text := scanner.Text()
			sub <- responseMsg(text)
		}
		if err := scanner.Err(); err != nil {
			return nil
		}
		return nil
	}
}

func waitForActivity(sub chan responseMsg) tea.Cmd {
	return func() tea.Msg {
		return responseMsg(<-sub)
	}
}
