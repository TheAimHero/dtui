package managevolume

import (
	"sort"
	"strings"

	"github.com/TheAimHero/dtui/internal/docker"
	ui_table "github.com/TheAimHero/dtui/internal/ui/table"
	"github.com/TheAimHero/dtui/internal/utils"
	"github.com/charmbracelet/bubbles/table"
)

func getTableRows(volumes docker.Volumes) []table.Row {
	tableRows := []table.Row{}
	if len(volumes) == 0 {
		return tableRows
	}
	sort.SliceStable(volumes, func(i, j int) bool { return volumes[i].Name > volumes[j].Name })
	for _, v := range volumes {
		var (
			volSize string
		)
		if v.UsageData != nil {
			volSize = utils.GetSize(v.UsageData.Size)
		} else {
			volSize = "Not Available"
		}
		tableRows = append(tableRows, table.Row{
			v.Name,
			utils.GetDate(v.CreatedAt),
			v.Mountpoint,
			volSize,
		})
	}
	return tableRows
}

func getTableColumns() []table.Column {
	return []table.Column{
		{Title: "Name", Width: utils.FloorMul(physicalWidth, 0.2)},
		{Title: "Created At", Width: utils.FloorMul(physicalWidth, 0.15)},
		{Title: "Mountpoint", Width: utils.FloorMul(physicalWidth, 0.45)},
		{Title: "Size", Width: utils.FloorMul(physicalWidth, 0.1)},
	}
}

func (m VolumeModel) getTable() table.Model {
	tableColumns := getTableColumns()
	tableRows := getTableRows(m.DockerClient.Volumes)
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
