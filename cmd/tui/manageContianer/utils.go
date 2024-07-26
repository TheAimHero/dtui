package managecontianer

import (
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textinput"
	mapset "github.com/deckarep/golang-set/v2"

	"github.com/TheAimHero/dtui/internal/docker"
	ui_table "github.com/TheAimHero/dtui/internal/ui/table"
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
	return []table.Column{
		{Title: "Select", Width: utils.FloorMul(physicalWidth, 0.05)},
		{Title: "Loading", Width: utils.FloorMul(physicalWidth, 0.05)},
		{Title: "ID", Width: utils.FloorMul(physicalWidth, 0.1)},
		{Title: "Name", Width: utils.FloorMul(physicalWidth, 0.3)},
		{Title: "Image", Width: utils.FloorMul(physicalWidth, 0.2)},
		{Title: "Status", Width: utils.FloorMul(physicalWidth, 0.2)},
	}
}

func (m ContainerModel) getTable() table.Model {
	tableColumns := getTableColumns()
	tableRows := getTableRows(
		m.DockerClient.Containers,
		m.SelectedContainers,
		m.InProcess,
		m.Spinner,
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

func getInput() textinput.Model {
	ip := textinput.New()
	ip.Placeholder = "Container Name"
	ip.Prompt = "Container Filter: "
	return ip
}
