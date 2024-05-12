package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

const (
	ContainerTab = iota
	ImageTab
	WipTab
)

func (m MainModel) View() string {
	return Tab(m)
}

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "1":
			m.ActiveTab = ContainerTab
			return m, nil
		case "2":
			m.ActiveTab = ImageTab
			return m, nil
		case "3":
			m.ActiveTab = WipTab
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
