package logs

import (
	"os"
	"strings"
	"time"

	"github.com/TheAimHero/dtui/internal/ui/message"
	"github.com/TheAimHero/dtui/internal/utils"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"golang.org/x/term"
)

func (m logModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)
	switch msg := msg.(type) {
	case utils.ResponseMsg:
		m.text = append(m.text, string(msg))
		m.viewport.SetContent(contentStyle.Render(strings.Join(m.text, "\n")))
		m.viewport.GotoBottom()
		cmds = append(cmds, utils.ResponseToStream(m.sub))

	case message.ClearMessage:
		m.message = message.Message{}

	case time.Time:
		err := m.dockerClient.FetchContainers()
		if err != nil {
			m.message.AddMessage("Error while fetching containers", message.ErrorMessage)
			cmds = append(cmds, m.message.ClearMessage(message.ErrorDuration))
		}
		tableRows := getTableRows(m.dockerClient.Containers)
		m.table.SetRows(tableRows)
		cmds = append(cmds, utils.TickCommand())

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit

		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll

		case key.Matches(msg, m.keys.Select):
			m, cmd = m.GetLogs()
			cmds = append(cmds, cmd, utils.ListenToStream(m.sub, m.stream))
		}

	case tea.WindowSizeMsg:
		physicalWidth, _, _ = term.GetSize(int(os.Stdout.Fd()))
		m.table = getTable(m.dockerClient.Containers)
		m.viewport.Width = msg.Width - 20
	}

	m.viewport, cmd = m.viewport.Update(msg)
	cmds = append(cmds, cmd)
	m.table, cmd = m.table.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}
