package manageimage

import (
	"time"

	"github.com/charmbracelet/bubbles/table"
	mapset "github.com/deckarep/golang-set/v2"

	"github.com/TheAimHero/dtui/internal/docker"
	"github.com/TheAimHero/dtui/internal/size"
)

func getTableRows(images docker.Images, selectedRows mapset.Set[string]) []table.Row {
	tableRows := []table.Row{}
	for _, image := range images {
		var selected string
		var tag string
		if len(image.RepoTags) > 0 {
			tag = image.RepoTags[0]
		} else {
			tag = "<none>"
		}
		if selectedRows.Contains(image.ID) {
			selected = " "
		}
		tableRows = append(tableRows, table.Row{
			selected,
			image.ID,
			tag,
			time.Unix(image.Created, 0).Format("02/01/2006 15:04 MST"),
			size.GetSize(image.Size),
		})
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

