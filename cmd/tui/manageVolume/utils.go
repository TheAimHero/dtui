package managevolume

import (
	"sort"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"
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
		if strings.Contains(row[VolumeName], filter) {
			filteredRows = append(filteredRows, row)
		}
	}
	return filteredRows
}

func getTableRows(volumes docker.Volumes, selectedVolumes mapset.Set[string], inProgress mapset.Set[string], sp spinner.Model) []table.Row {
	tableRows := []table.Row{}
	if len(volumes) == 0 {
		return tableRows
	}
	sort.SliceStable(volumes, func(i, j int) bool { return volumes[i].Name > volumes[j].Name })
	for _, v := range volumes {
		var (
			volSize     string
			selected    string
			spinnerView string
		)
		if selectedVolumes.Contains(v.Name) {
			selected = "✓ "
		} else {
			selected = "  "
		}
		if inProgress.Contains(v.Name) {
			spinnerView = sp.View()
		} else {
			spinnerView = ""
		}
		if v.UsageData != nil {
			volSize = utils.GetSize(v.UsageData.Size)
		} else {
			volSize = "Not Available"
		}
		tableRows = append(tableRows, table.Row{
			selected,
			spinnerView,
			v.Name,
			utils.GetDate(v.CreatedAt),
			v.Mountpoint,
			volSize,
		})
	}
	return tableRows
}

func getTableColumns(width int) []table.Column {
	return []table.Column{
		{Title: "Select", Width: utils.FloorMul(width, 0.05)},
		{Title: "Loading", Width: utils.FloorMul(width, 0.05)},
		{Title: "Name", Width: utils.FloorMul(width, 0.2)},
		{Title: "Created At", Width: utils.FloorMul(width, 0.15)},
		{Title: "Mountpoint", Width: utils.FloorMul(width, 0.4)},
		{Title: "Size", Width: utils.FloorMul(width, 0.1)},
	}
}

func (m VolumeModel) getTable() table.Model {
	tableColumns := getTableColumns(m.Width)
	tableRows := getTableRows(m.Volumes, m.SelectedVolumes, m.InProgress, m.Spinner)
	return components.NewStandardTable(tableColumns, tableRows)
}

func getSpinner() spinner.Model {
	s := spinner.New()
	s.Spinner = spinner.Points
	s.Spinner.FPS = 300 * time.Millisecond
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	return s
}

func getInput() textinput.Model {
	ip := textinput.New()
	ip.Placeholder = "Volume Name"
	ip.Prompt = "Volume Filter: "
	return ip
}
