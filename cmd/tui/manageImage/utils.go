package manageimage

import (
	"time"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	mapset "github.com/deckarep/golang-set/v2"

	"github.com/TheAimHero/dtui/internal/docker"
	"github.com/TheAimHero/dtui/internal/size"
)

func getTableRows(images docker.Images, selectedRows mapset.Set[string]) []table.Row {
	tableRows := make([]table.Row, len(images))
	for i, image := range images {
		var selected string
		if selectedRows.Contains(image.ID) {
			selected = " "
		}
		tableRows[i] = table.Row{
			selected,
			image.ID,
			image.RepoTags[0],
			time.Unix(image.Created, 0).Format("02/01/2006 15:04 MST"),
			size.GetSize(image.Size),
		}
	}
	return tableRows
}

func getTableColumns() []table.Column {
	width := (physicalWidth / 5) - 4
	return []table.Column{
		{Title: "Select", Width: width},
		{Title: "ID", Width: width},
		{Title: "Name", Width: width},
		{Title: "Created", Width: width},
		{Title: "Size", Width: width},
	}
}

func tickCommand() tea.Cmd {
	return tea.Tick(300*time.Millisecond, func(t time.Time) tea.Msg {
		return t
	})
}
