package logs

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/TheAimHero/dtui/internal/docker"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
)

func getTableRows(containers docker.Containers) []table.Row {
	tableRows := []table.Row{}
	for _, container := range containers {
		tableRows = append(tableRows, table.Row{
			container.ID,
			strings.Split(container.Names[0], "/")[1],
			container.Image,
			strings.ToUpper(string(container.Status[0])) + string(container.Status[1:]),
		})
	}
	return tableRows
}

func getTableColumns() []table.Column {
	width := ((physicalWidth) / 4) - 4
	return []table.Column{
		{Title: "ID", Width: width},
		{Title: "Name", Width: width},
		{Title: "Image", Width: width},
		{Title: "Status", Width: width},
	}
}

type responseMsg string

func listenForActivity(sub chan<- responseMsg, stream io.ReadCloser) tea.Cmd {
	if stream == nil {
		return nil
	}
	return func() tea.Msg {
		defer stream.Close()
		scanner := bufio.NewScanner(stream)
		for scanner.Scan() {
			text := scanner.Text()
			sub <- responseMsg(text)
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "error reading input: %v\n", err)
		}
		return nil
	}
}

func waitForActivity(sub chan responseMsg) tea.Cmd {
	return func() tea.Msg {
		return responseMsg(<-sub)
	}
}
