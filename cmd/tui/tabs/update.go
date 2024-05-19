package tabs

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m MainModel) getNextTab(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch m.ActiveTab {
	case ContainerTab:
		m.ActiveTab = ImageTab
		m.ImageTab, cmd = m.ImageTab.Update(msg)
		return m, tea.Sequence(m.ImageTab.Init(), cmd)

	case ImageTab:
		m.ActiveTab = LogsTab
		m.LogsTab, cmd = m.LogsTab.Update(msg)
		return m, tea.Sequence(m.LogsTab.Init(), cmd)

	case LogsTab:
		m.ActiveTab = WipTab
		m.WipTab, cmd = m.WipTab.Update(msg)
		return m, tea.Sequence(m.WipTab.Init(), cmd)

	default:
		m.ActiveTab = ContainerTab
		m.ContainerTab, cmd = m.ContainerTab.Update(msg)
		return m, tea.Sequence(m.ContainerTab.Init(), cmd)
	}
}

func (m MainModel) getPrevTab(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch m.ActiveTab {
	case ContainerTab:
		m.ActiveTab = WipTab
		m.WipTab, cmd = m.WipTab.Update(msg)
		return m, tea.Sequence(m.WipTab.Init(), cmd)

	case ImageTab:
		m.ActiveTab = ContainerTab
		m.ContainerTab, cmd = m.ContainerTab.Update(msg)
		return m, tea.Sequence(m.ContainerTab.Init(), cmd)

	case LogsTab:
		m.ActiveTab = ImageTab
		m.ImageTab, cmd = m.ImageTab.Update(msg)
		return m, tea.Sequence(m.ImageTab.Init(), cmd)

	default:
		m.ActiveTab = LogsTab
		m.LogsTab, cmd = m.LogsTab.Update(msg)
		return m, tea.Sequence(m.LogsTab.Init(), cmd)
	}
}

func (m MainModel) updateCurrentTab(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch m.ActiveTab {
	case ContainerTab:
		m.ContainerTab, cmd = m.ContainerTab.Update(msg)

	case ImageTab:
		m.ImageTab, cmd = m.ImageTab.Update(msg)

	case LogsTab:
		m.LogsTab, cmd = m.LogsTab.Update(msg)

	default:
		m.WipTab, cmd = m.WipTab.Update(msg)
	}
	return m, cmd
}

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "1":
			m.ActiveTab = ContainerTab
			m.ContainerTab, cmd = m.ContainerTab.Update(msg)
			return m, cmd

		case "2":
			m.ActiveTab = ImageTab
			m.ImageTab, cmd = m.ImageTab.Update(msg)
			return m, cmd

		case "3":
			m.ActiveTab = LogsTab
			m.LogsTab, cmd = m.LogsTab.Update(msg)
			return m, cmd

		case "4":
			m.ActiveTab = WipTab
			m.WipTab, cmd = m.WipTab.Update(msg)
			return m, cmd

		case "right", "l":
			return m.getNextTab(msg)

		case "left", "h":
			return m.getPrevTab(msg)

		case "ctrl+c", "q":
			return m, tea.Quit

		default:
			return m.updateCurrentTab(msg)
		}

	case tea.WindowSizeMsg:
		return m.updateCurrentTab(msg)

	default:
		return m.updateCurrentTab(msg)
	}
}
