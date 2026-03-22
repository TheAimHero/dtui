package managecontainer

import (
	"strings"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textinput"
	mapset "github.com/deckarep/golang-set/v2"

	"github.com/TheAimHero/dtui/internal/docker"
	"github.com/TheAimHero/dtui/internal/ui/components"
	"github.com/TheAimHero/dtui/internal/utils"
)

func filterRows(rows []table.Row, filter string) []table.Row {
	if filter == "" {
		return rows
	}
	var filteredRows []table.Row
	for _, row := range rows {
		if strings.Contains(row[ContainerName], filter) {
			filteredRows = append(filteredRows, row)
		}
	}
	return filteredRows
}

func getTableRows(containers docker.Containers, selectedContainers mapset.Set[string], inProcess mapset.Set[string], spinner spinner.Model) []table.Row {
	tableRows := []table.Row{}
	for _, container := range containers {
		var selected string
		var spinnerView string
		if selectedContainers.Contains(container.ID) {
			selected = "✓ "
		} else {
			selected = "  "
		}
		if inProcess.Contains(container.ID) {
			spinnerView = spinner.View()
		} else {
			spinnerView = ""
		}
		tableRows = append(tableRows, table.Row{
			selected,
			spinnerView,
			container.ID,
			strings.Split(container.Names[0], "/")[1],
			container.Image,
			strings.ToUpper(string(container.Status[0])) + string(container.Status[1:]),
		})
	}
	return tableRows
}

func getTableColumns(width int) []table.Column {
	return []table.Column{
		{Title: "Select", Width: utils.FloorMul(width, 0.05)},
		{Title: "Loading", Width: utils.FloorMul(width, 0.05)},
		{Title: "ID", Width: utils.FloorMul(width, 0.1)},
		{Title: "Name", Width: utils.FloorMul(width, 0.3)},
		{Title: "Image", Width: utils.FloorMul(width, 0.2)},
		{Title: "Status", Width: utils.FloorMul(width, 0.2)},
	}
}

func (m ContainerModel) getTable() table.Model {
	tableColumns := getTableColumns(m.Width)
	tableRows := getTableRows(
		m.Containers,
		m.SelectedContainers,
		m.InProcess,
		m.Spinner,
	)
	return components.NewStandardTable(tableColumns, tableRows)
}

func getInput() textinput.Model {
	ip := textinput.New()
	ip.Placeholder = "Container Name"
	ip.Prompt = "Container Filter: "
	return ip
}
