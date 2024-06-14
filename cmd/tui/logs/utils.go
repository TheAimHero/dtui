package logs

import (
	"os/exec"
	"strings"

	"github.com/TheAimHero/dtui/internal/docker"
	ui_table "github.com/TheAimHero/dtui/internal/ui/table"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	ContainerID = iota
	ContainerName
	ContainerImage
	ContainerStatus
)

func getTable(containers docker.Containers) table.Model {
	tableColumns := getTableColumns()
	tableRows := getTableRows(containers)
	table := ui_table.NewTable(tableColumns, tableRows)
	return table
}

func showLogs(containerID string) tea.Cmd {
	c := exec.Command("docker", "logs", containerID, "--follow")
	return tea.ExecProcess(c, func(err error) tea.Msg { return nil })
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
