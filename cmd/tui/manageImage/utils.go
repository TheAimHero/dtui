package manageimage

import (
	"time"

	"github.com/TheAimHero/dtui/internal/docker"
	"github.com/TheAimHero/dtui/internal/size"
	"github.com/charmbracelet/bubbles/table"
)

func getTableRows(images docker.Images) []table.Row {
	tableRows := make([]table.Row, len(images))
	for i, image := range images {
		tableRows[i] = table.Row{
			image.ID,
			image.RepoTags[0],
			time.Unix(image.Created, 0).Format("02/01/2006 15:04 MST"),
			size.GetSize(image.Size),
		}
	}
	return tableRows
}

func getTableColumns() []table.Column {
	width := (physicalWidth / 4) - 4
	return []table.Column{
		{Title: "ID", Width: width},
		{Title: "Name", Width: width},
		{Title: "Created", Width: width},
		{Title: "Size", Width: width},
	}
}
