package tabs

import (
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"golang.org/x/term"
)

func (m MainModel) getNextTab(_ tea.Msg) (tea.Model, tea.Cmd) {
	var cmd  tea.Cmd
	switch m.ActiveTab {
	case ContainerTab:
		m.ActiveTab = ImageTab
		m.ImageTab, cmd = m.ImageTab.Update(tea.WindowSizeMsg{})
		return m, tea.Sequence(m.ImageTab.Init(), cmd)

	case ImageTab:
		m.ActiveTab = LogsTab
		m.LogsTab, cmd = m.LogsTab.Update(tea.WindowSizeMsg{Width: physicalWidth, Height: physicalHeight})
		return m, tea.Sequence(m.LogsTab.Init(), cmd)

	case LogsTab:
		m.ActiveTab = WipTab
		m.WipTab, cmd = m.WipTab.Update(tea.WindowSizeMsg{Width: physicalWidth, Height: physicalHeight})
		return m, tea.Sequence(m.WipTab.Init(), cmd)

	default:
		m.ActiveTab = ContainerTab
		m.ContainerTab, cmd = m.ContainerTab.Update(tea.WindowSizeMsg{Width: physicalWidth, Height: physicalHeight})
		return m, tea.Sequence(m.ContainerTab.Init(), cmd)
	}
}

func (m MainModel) getPrevTab(_ tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch m.ActiveTab {
	case ContainerTab:
		m.ActiveTab = WipTab
		m.WipTab, cmd = m.WipTab.Update(tea.WindowSizeMsg{Width: physicalWidth, Height: physicalHeight})
		return m, tea.Sequence(m.WipTab.Init(), cmd)

	case ImageTab:
		m.ActiveTab = ContainerTab
		m.ContainerTab, cmd = m.ContainerTab.Update(tea.WindowSizeMsg{Width: physicalWidth, Height: physicalHeight})
		return m, tea.Sequence(m.ContainerTab.Init(), cmd)

	case LogsTab:
		m.ActiveTab = ImageTab
		m.ImageTab, cmd = m.ImageTab.Update(tea.WindowSizeMsg{Width: physicalWidth, Height: physicalHeight})
		return m, tea.Sequence(m.ImageTab.Init(), cmd)

	default:
		m.ActiveTab = LogsTab
		m.LogsTab, cmd = m.LogsTab.Update(tea.WindowSizeMsg{Width: physicalWidth, Height: physicalHeight})
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
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)
	if m.ImageTab.Input.Focused() {
		m.ImageTab, cmd = m.ImageTab.Update(msg)
		return m, cmd
	}
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "1":
			m.ActiveTab = ContainerTab
			m.ContainerTab, cmd = m.ContainerTab.Update(msg)
			cmds = append(cmds, cmd)
			m.ContainerTab, cmd = m.ContainerTab.Update(tea.WindowSizeMsg{Width: physicalWidth, Height: physicalHeight})
			cmds = append(cmds, cmd)
			return m, tea.Batch(cmds...)

		case "2":
			m.ActiveTab = ImageTab
			m.ImageTab, cmd = m.ImageTab.Update(msg)
			cmds = append(cmds, cmd)
			m.ImageTab, cmd = m.ImageTab.Update(tea.WindowSizeMsg{Width: physicalWidth, Height: physicalHeight})
			cmds = append(cmds, cmd)
			return m, tea.Batch(cmds...)

		case "3":
			m.ActiveTab = LogsTab
			m.LogsTab, cmd = m.LogsTab.Update(msg)
			cmds = append(cmds, cmd)
			m.LogsTab, cmd = m.LogsTab.Update(tea.WindowSizeMsg{Width: physicalWidth, Height: physicalHeight})
			cmds = append(cmds, cmd)
			return m, tea.Batch(cmds...)

		case "4":
			m.ActiveTab = WipTab
			m.WipTab, cmd = m.WipTab.Update(msg)
			cmds = append(cmds, cmd)
			m.WipTab, cmd = m.WipTab.Update(tea.WindowSizeMsg{Width: physicalWidth, Height: physicalHeight})
			cmds = append(cmds, cmd)
			return m, tea.Batch(cmds...)

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
		physicalWidth, physicalHeight, _ = term.GetSize(int(os.Stdout.Fd()))
		return m.updateCurrentTab(tea.WindowSizeMsg{Width: physicalWidth, Height: physicalHeight})

	default:
		return m.updateCurrentTab(msg)
	}
}
