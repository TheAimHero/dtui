package manageimage

import (
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"golang.org/x/term"

	"github.com/TheAimHero/dtui/internal/ui/message"
	ui_table "github.com/TheAimHero/dtui/internal/ui/table"
	ui_utils "github.com/TheAimHero/dtui/internal/ui/utils"
	"github.com/TheAimHero/dtui/internal/utils"
)

var (
	physicalWidth, physicalHeight, _ = term.GetSize(int(os.Stdout.Fd())) // nolint:unused
	successDuration                  = 2 * time.Second
	errorDuration                    = 5 * time.Second
)

func (m imageModel) View() string {
	doc := strings.Builder{}
	doc.WriteString(ui_table.BaseTableStyle.Render(m.table.View()))
	doc.WriteString("\n" + m.message.ShowMessage())
	doc.WriteString("\n" + m.help.View(m.keys))
	doc.WriteString(strings.Repeat("\n", ui_utils.HeightPadding(doc, 7)))
	return doc.String()
}

func (m imageModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		physicalWidth, physicalHeight, _ = term.GetSize(int(os.Stdout.Fd()))
		m.table = getTable(m.dockerClient.Images, m.selectedImages)
		return m, cmd

	case message.ClearErrorMsg:
		m.message = message.Message{}

	case time.Time:
		err := m.dockerClient.FetchImages()
		if err != nil {
			m.message.AddMessage("Error while fetching images", message.ErrorMessage)
			return m, tea.Batch(utils.TickCommand(), m.message.ClearMessage(message.ErrorDuration))
		}
		tableRows := getTableRows(m.dockerClient.Images, m.selectedImages)
		m.table.SetRows(tableRows)
		return m, utils.TickCommand()

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit

		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll

		case key.Matches(msg, m.keys.SelectImage):
			return m.SelectImage()

		case key.Matches(msg, m.keys.SelectAllImages):
			return m.SelectAllImages()

		case key.Matches(msg, m.keys.DeleteImage):
			return m.DeleteImage()

		case key.Matches(msg, m.keys.DeleteSelectedImages):
			return m.DeleteImages()
		}
	}

	m.table, cmd = m.table.Update(msg)
	return m, cmd
}
