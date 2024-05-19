package managecontianer

import (
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/table"
	mapset "github.com/deckarep/golang-set/v2"

	"github.com/TheAimHero/dtui/internal/docker"
	ui_table "github.com/TheAimHero/dtui/internal/ui/table"
)

func getTableRows(containers docker.Containers, selectedContainers mapset.Set[string], inProcesss mapset.Set[string], spinner spinner.Model) []table.Row {
	tableRows := []table.Row{}
	for _, container := range containers {
		var selected string
		var spinnerView string
		if selectedContainers.Contains(container.ID) {
			selected = " "
		} else {
			selected = "  "
		}
		if inProcesss.Contains(container.ID) {
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

func getTableColumns() []table.Column {
	width := ((physicalWidth) / 4) - 10
	return []table.Column{
		{Title: "Select", Width: 8},
		{Title: "Loading", Width: 9},
		{Title: "ID", Width: width},
		{Title: "Name", Width: width},
		{Title: "Image", Width: width},
		{Title: "Status", Width: width},
	}
}

func (m ContainerModel) getTable() table.Model {
	tableColumns := getTableColumns()
	tableRows := getTableRows(
		m.dockerClient.Containers,
		m.selectedContainers,
		m.inProcesss,
		m.spinner,
	)
	table := ui_table.NewTable(tableColumns, tableRows)
	return table
}

func getSpinner() spinner.Model {
	s := spinner.New()
	s.Spinner = spinner.Points
	s.Spinner.FPS = 300 * time.Millisecond
	return s
}
