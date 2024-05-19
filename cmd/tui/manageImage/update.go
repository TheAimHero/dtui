package manageimage

import (
	"os"
	"time"

	"github.com/TheAimHero/dtui/internal/ui/message"
	"github.com/TheAimHero/dtui/internal/utils"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"golang.org/x/term"
)

func (m imageModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		physicalWidth, physicalHeight, _ = term.GetSize(int(os.Stdout.Fd()))
		m.table = getTable(m.dockerClient.Images, m.selectedImages)

	case message.ClearMessage:
		m.message = message.Message{}

	case time.Time:
		err := m.dockerClient.FetchImages()
		if err != nil {
			m.message.AddMessage("Error while fetching images", message.ErrorMessage)
			cmds = append(cmds, utils.TickCommand(), m.message.ClearMessage(message.ErrorDuration))
		}
		tableRows := getTableRows(m.dockerClient.Images, m.selectedImages)
		m.table.SetRows(tableRows)
		cmds = append(cmds, utils.TickCommand())

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit

		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll

		case key.Matches(msg, m.keys.SelectImage):
			m, cmd = m.SelectImage()
			cmds = append(cmds, cmd)

		case key.Matches(msg, m.keys.SelectAllImages):
			m, cmd = m.SelectAllImages()
			cmds = append(cmds, cmd)

		case key.Matches(msg, m.keys.DeleteImage):
			m, cmd = m.DeleteImage()
			cmds = append(cmds, cmd)

		case key.Matches(msg, m.keys.DeleteSelectedImages):
			m, cmd = m.DeleteImages()
			cmds = append(cmds, cmd)
		}
	}
	m.table, cmd = m.table.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}
