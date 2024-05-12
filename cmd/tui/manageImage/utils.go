package manageimage

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/TheAimHero/dtui/internal/docker"
	"github.com/TheAimHero/dtui/internal/size"
)

func getTableRows(images docker.Images) []table.Row {
	tableRows := make([]table.Row, len(images))
	for i, image := range images {
		tableRows[i] = table.Row{
			image.ID,
			image.RepoTags[0],
			time.Unix(image.Created, 0).Format("02/01/2006 15:04 MST"),
			size.GetSize(image.Size),
			fmt.Sprintf("%d", image.Containers),
		}
	}
	return tableRows
}

func getTableColumns() []table.Column {
	width := (physicalWidth / 5) - 4
	return []table.Column{
		{Title: "ID", Width: width},
		{Title: "Name", Width: width},
		{Title: "Created", Width: width},
		{Title: "Size", Width: width},
		{Title: "Containers", Width: width},
	}
}

func tickCommand() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return t
	})
}
