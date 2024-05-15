package managecontianer

import (
	"strings"

	"github.com/TheAimHero/dtui/internal/docker"
	ui_table "github.com/TheAimHero/dtui/internal/ui/table"
	"github.com/charmbracelet/bubbles/table"
	mapset "github.com/deckarep/golang-set/v2"
)

func getTableRows(containers docker.Containers, selectedRows mapset.Set[string]) []table.Row {
	tableRows := []table.Row{}
	for _, container := range containers {
		var selected string
		if selectedRows.Contains(container.ID) {
			selected = " "
		} else {
			selected = "  "
		}
		tableRows = append(tableRows, table.Row{
			selected,
			container.ID,
			strings.Split(container.Names[0], "/")[1],
			container.Image,
			strings.ToUpper(string(container.Status[0])) + string(container.Status[1:]),
		})
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

func getTable(containers docker.Containers, selectedRows mapset.Set[string]) table.Model {
	tableColumns := getTableColumns()
	tableRows := getTableRows(containers, selectedRows)
	table := ui_table.NewTable(tableColumns, tableRows)
	table.KeyMap.HalfPageDown.Unbind()
	table.KeyMap.HalfPageUp.Unbind()
	table.KeyMap.GotoBottom.Unbind()
	table.KeyMap.GotoTop.Unbind()
	table.KeyMap.PageDown.Unbind()
	table.KeyMap.PageUp.Unbind()
	return table
}
