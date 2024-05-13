package utils

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

func TickCommand() tea.Cmd {
	return tea.Tick(300*time.Millisecond, func(t time.Time) tea.Msg {
		return t
	})
}
