package tui

import (
	tea "github.com/charmbracelet/bubbletea"
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

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "1":
			m.ActiveTab = ContainerTab
			return m, m.Tabs[m.ActiveTab].Init()
		case "2":
			m.ActiveTab = ImageTab
			return m, m.Tabs[m.ActiveTab].Init()
		case "3":
			m.ActiveTab =LogsTab 
			return m, m.Tabs[m.ActiveTab].Init()
		case "4":
			m.ActiveTab = WipTab
			return m, m.Tabs[m.ActiveTab].Init()
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		m.Tabs[m.ActiveTab], cmd = m.Tabs[m.ActiveTab].Update(msg)
		return m, cmd
	}
	activeTab := m.Tabs[m.ActiveTab]
	updatedTab, cmd := activeTab.Update(msg)
	m.Tabs[m.ActiveTab] = updatedTab
	return m, cmd
}
