package tui

import (
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"golang.org/x/term"
)

const (
	ContainerTab = iota
	ImageTab
	LogsTab
	WipTab
)

func (m MainModel) View() string {
	return Tab(m)
}

func (m *MainModel) callInit() tea.Cmd {
	if m.InitTab.Contains(m.ActiveTab) {
		return nil
	}
  m.InitTab.Add(m.ActiveTab)
	return m.Tabs[m.ActiveTab].Init()
}

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "1":
			m.ActiveTab = ContainerTab
			return m, m.callInit()
		case "2":
			m.ActiveTab = ImageTab
			return m, m.callInit()
		case "3":
			m.ActiveTab = LogsTab
			return m, m.callInit()
		case "4":
			m.ActiveTab = WipTab
			return m, m.callInit()
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		physicalWidth, physicalHeight, _ := term.GetSize(int(os.Stdout.Fd()))
		m.Tabs[m.ActiveTab], cmd = m.Tabs[m.ActiveTab].Update(tea.WindowSizeMsg{Height: physicalHeight, Width: physicalWidth})
		return m, cmd
	}
	activeTab := m.Tabs[m.ActiveTab]
	updatedTab, cmd := activeTab.Update(msg)
	m.Tabs[m.ActiveTab] = updatedTab
	return m, tea.Batch(cmd)
}
