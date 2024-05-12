package managecontianer

import (
	"strings"
	"time"

	"github.com/TheAimHero/dtui/internal/docker"
	"github.com/TheAimHero/dtui/internal/helpers"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
)

func getTableRows(containers docker.Containers, selectedRows []string) []table.Row {
	tableRows := make([]table.Row, len(containers))
	for i, container := range containers {
		var selected string
		if helpers.InArray(container.ID, selectedRows) {
			selected = " "
		} else {
			selected = "  "
		}
		tableRows[i] = table.Row{
			selected,
			container.ID,
			strings.Split(container.Names[0], "/")[1],
			container.Image,
			strings.ToUpper(string(container.Status[0])) + string(container.Status[1:]),
		}
	}
	return tableRows
}

func getTableColumns() []table.Column {
	width := ((physicalWidth) / 5) - 4
	return []table.Column{
		{Title: "Select", Width: width},
		{Title: "ID", Width: width},
		{Title: "Name", Width: width},
		{Title: "Image", Width: width},
		{Title: "Status", Width: width},
	}
}

func tickCommand() tea.Cmd {
	return tea.Tick(300*time.Millisecond, func(t time.Time) tea.Msg {
		return t
	})
}
