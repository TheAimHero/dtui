package tui

import tea "github.com/charmbracelet/bubbletea"

func (m MainModel) View() string {
	return Tab(m)
}

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "1":
			m.ActiveTab = 0
			return m, nil
		case "2":
			m.ActiveTab = 1
			return m, nil
		case "3":
			m.ActiveTab = 2
			return m, nil
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}
	activeTab := m.Tabs[m.ActiveTab]
	updatedTab, cmd := activeTab.Update(msg)
	m.Tabs[m.ActiveTab] = updatedTab
	return m, cmd
}
