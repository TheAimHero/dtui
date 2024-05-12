package managecontianer

import (
	"strings"
	"time"

	"github.com/TheAimHero/dtui/internal/docker"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
)

func getTableRows(containers docker.Containers) []table.Row {
	tableRows := make([]table.Row, len(containers))
	for i, container := range containers {
		tableRows[i] = table.Row{
			container.ID,
			strings.Split(container.Names[0], "/")[1],
			strings.ToUpper(string(container.Status[0])) + string(container.Status[1:]),
			strings.ToUpper(string(container.State[0])) + string(container.State[1:]),
		}
	}
	return tableRows
}

func getTableColumns() []table.Column {
	width := (physicalWidth / 4) - 4
	return []table.Column{
		{Title: "ID", Width: width},
		{Title: "Name", Width: width},
		{Title: "Status", Width: width},
		{Title: "State", Width: width},
	}
}

func tickCommand() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return t
	})
}

