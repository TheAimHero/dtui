package manageimage

import (
	"os"
	"strings"

	"github.com/TheAimHero/dtui/internal/ui"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"golang.org/x/term"
)

var (
	physicalWidth, physicalHeight, _ = term.GetSize(int(os.Stdout.Fd()))
)

func (m imageModel) View() string {
	doc := strings.Builder{}
	doc.WriteString(ui.BaseTableStyle.Render(m.table.View()) + m.message.ShowMessage())
	doc.WriteString("\n" + m.help.View(m.keys))
	return doc.String()
}

func (m imageModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch {

		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit

		case key.Matches(msg, m.keys.Down):
			m.table.MoveDown(1)
			return m, nil

		case key.Matches(msg, m.keys.Up):
			m.table.MoveUp(1)
			return m, nil

		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll

		}
	}
	return m, cmd
}
