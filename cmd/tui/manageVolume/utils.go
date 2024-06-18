package managevolume

import (
	"sort"
	"strings"

	"github.com/TheAimHero/dtui/internal/docker"
	ui_table "github.com/TheAimHero/dtui/internal/ui/table"
	"github.com/TheAimHero/dtui/internal/utils"
	"github.com/charmbracelet/bubbles/table"
	mapset "github.com/deckarep/golang-set/v2"
)

func getTableRows(volumes docker.Volumes, selectedVolumes mapset.Set[string]) []table.Row {
	tableRows := []table.Row{}
	if len(volumes) == 0 {
		return tableRows
	}
	sort.SliceStable(volumes, func(i, j int) bool { return volumes[i].Name > volumes[j].Name })
	for _, v := range volumes {
		var (
			selected string
			volSize  string
		)
		if selectedVolumes.Contains(v.Name) {
			selected = " "
		} else {
			selected = "  "
		}
		if v.UsageData != nil {
			volSize = utils.GetSize(v.UsageData.Size)
		} else {
			volSize = "Not Available"
		}
		tableRows = append(tableRows, table.Row{
			selected,
			v.Name,
			utils.GetDate(v.CreatedAt),
			v.Mountpoint,
			volSize,
		})
	}
	return tableRows
}

func getTableColumns() []table.Column {
	// width := ((physicalWidth) / 1)
	return []table.Column{
		{Title: "Select", Width: 8},
		{Title: "Name", Width: 20},
		{Title: "Created At", Width: 30},
		{Title: "Mountpoint", Width: physicalWidth - 100},
		{Title: "Size", Width: 13},
	}
}

func (m VolumeModel) getTable() table.Model {
	tableColumns := getTableColumns()
	tableRows := getTableRows(
		m.DockerClient.Volumes,
		m.SelectedVolumes,
	)
	table := ui_table.NewTable(tableColumns, tableRows)
	return table
}

func filterRows(rows []table.Row, filter string) []table.Row { // nolint:unused
	var filteredRows []table.Row
	for _, row := range rows {
		if strings.Contains(row[VolumeName], filter) {
			filteredRows = append(filteredRows, row)
		}
	}
	return filteredRows
}
