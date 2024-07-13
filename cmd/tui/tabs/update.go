package tabs

import (
	"os"

	managecontianer "github.com/TheAimHero/dtui/cmd/tui/manageContianer"
	tea "github.com/charmbracelet/bubbletea"
	"golang.org/x/term"
)

func (m MainModel) getNextTab(_ tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.ActiveTab = (m.ActiveTab + 1) % len(m.Tabs)
	m.Tabs[m.ActiveTab%len(m.Tabs)], cmd = m.Tabs[m.ActiveTab%len(m.Tabs)].Update(tea.WindowSizeMsg{Width: physicalWidth, Height: physicalHeight})
	return m, tea.Sequence(m.Tabs[m.ActiveTab%len(m.Tabs)].Init(), cmd)
}

func (m MainModel) getPrevTab(_ tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.ActiveTab = (m.ActiveTab - 1 + len(m.Tabs)) % len(m.Tabs)
	m.Tabs[m.ActiveTab%len(m.Tabs)], cmd = m.Tabs[m.ActiveTab%len(m.Tabs)].Update(tea.WindowSizeMsg{Width: physicalWidth, Height: physicalHeight})
	return m, tea.Sequence(m.Tabs[m.ActiveTab%len(m.Tabs)].Init(), cmd)
}

func (m MainModel) updateCurrentTab(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.Tabs[m.ActiveTab%len(m.Tabs)], cmd = m.Tabs[m.ActiveTab%len(m.Tabs)].Update(msg)
	return m, cmd
}

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)
	if s, ok := m.Tabs[m.ActiveTab].(managecontianer.ContainerModel); ok && s.Input.Focused() {
		m.Tabs[m.ActiveTab], cmd = s.Update(msg)
		return m, cmd
	}
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "1", "2", "3", "4":
			m.ActiveTab = ContainerTab
			m.Tabs[m.ActiveTab], cmd = m.Tabs[m.ActiveTab].Update(msg)
			cmds = append(cmds, cmd)
			m.Tabs[m.ActiveTab], cmd = m.Tabs[m.ActiveTab].Update(tea.WindowSizeMsg{Width: physicalWidth, Height: physicalHeight})
			cmds = append(cmds, cmd)
			return m, tea.Batch(cmds...)

		case "right", "l":
			return m.getNextTab(msg)

		case "left", "h":
			return m.getPrevTab(msg)

		case "ctrl+c", "q":
			return m, tea.Sequence(tea.Quit)

		default:
			return m.updateCurrentTab(msg)
		}

	case tea.WindowSizeMsg:
		physicalWidth, physicalHeight, _ = term.GetSize(int(os.Stdout.Fd()))
		return m.updateCurrentTab(tea.WindowSizeMsg{Width: physicalWidth, Height: physicalHeight})

	default:
		return m.updateCurrentTab(msg)
	}
}
