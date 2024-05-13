package logs

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/TheAimHero/dtui/internal/ui"
	"github.com/TheAimHero/dtui/internal/utils"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

const (
	ContainerID = iota
	ContainerName
	ContainerImage
	ContainerStatus
)

var (
	physicalWidth, physicalHeight, _ = term.GetSize(int(os.Stdout.Fd()))

	lineStyle = lipgloss.NewStyle().Foreground(ui.HighlightColor)

	contentStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#F8F8F2"))

	titleStyle = func() lipgloss.Style {
		b := lipgloss.NormalBorder()
		b.Right = "├"
		return lipgloss.NewStyle().BorderStyle(b).BorderForeground(ui.HighlightColor)
	}()
	infoStyle = func() lipgloss.Style {
		b := lipgloss.NormalBorder()
		b.Left = "┤"
		return titleStyle.Copy().BorderStyle(b).BorderForeground(ui.HighlightColor)
	}()
)

func (m logModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)
	switch msg := msg.(type) {
	case responseMsg:
		m.text = append(m.text, string(msg))
		m.viewport.SetContent(contentStyle.Render(strings.Join(m.text, "\n")))
		m.viewport.GotoBottom()
		return m, tea.Batch(waitForActivity(m.sub), tea.Println(msg))

	case time.Time:
		m.dockerClient.FetchContainers()
		tableRows := getTableRows(m.dockerClient.Containers)
		m.table.SetRows(tableRows)
		return m, utils.TickCommand()

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit

		case key.Matches(msg, m.keys.LogsUp):
			m.viewport.HalfViewUp()
		case key.Matches(msg, m.keys.Down):
			m.table.MoveDown(1)
			return m, nil

		case key.Matches(msg, m.keys.Up):
			m.table.MoveUp(1)
			return m, nil

		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll

		case key.Matches(msg, m.keys.LogsDown):
			m.viewport.HalfViewDown()

		case key.Matches(msg, m.keys.Select):
			m, cmd = m.GetLogs()
			return m, tea.Batch(cmd, listenForActivity(m.sub, m.stream))
		}

	case tea.WindowSizeMsg:
		physicalWidth, physicalHeight, _ = term.GetSize(int(os.Stdout.Fd()))
		m.table = getTable(m.dockerClient.Containers)
		m.viewport.Width = msg.Width - 20
		return m, cmd
	}
	m.viewport, cmd = m.viewport.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m logModel) View() string {
	doc := strings.Builder{}
	doc.WriteString(ui.BaseTableStyle.Copy().Margin(1).Render(m.table.View()))
	doc.WriteString(fmt.Sprintf("%s\n%s\n%s", m.headerView(), m.viewport.View(), m.footerView()))
	doc.WriteString("\n" + m.message.ShowMessage())
	doc.WriteString("\n" + m.help.View(m.keys))
	return doc.String()
}

func (m logModel) headerView() string {
	title := titleStyle.Render("Container Log: " + m.title)
	line := lineStyle.Render(strings.Repeat("─", max(0, m.viewport.Width-lipgloss.Width(title))))
	return lipgloss.JoinHorizontal(lipgloss.Center, title, line)
}

func (m logModel) footerView() string {
	info := infoStyle.Render(fmt.Sprintf("%3.f%%", m.viewport.ScrollPercent()*100))
	line := lineStyle.Render(strings.Repeat("─", max(0, m.viewport.Width-lipgloss.Width(info))))
	return lipgloss.JoinHorizontal(lipgloss.Center, line, info)
}
