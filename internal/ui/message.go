package ui

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Message struct {
	value   string
	msgType int
}

const (
	ErrorMessage = iota
	SuccessMessage
)

const (
  ErrorDuration = 5 * time.Second
  SuccessDuration = 2 * time.Second
)

type ClearErrorMsg struct{}

var (
	msgStyle     = lipgloss.NewStyle().Padding(1, 0)
	errStyle     = msgStyle.Copy().Foreground(lipgloss.Color("#cb4154"))
	successStyle = msgStyle.Copy().Foreground(lipgloss.Color("#AADB1E"))
)

func (msg *Message) AddMessage(value string, messageType int) {
	msg.msgType = messageType
	msg.value = value
}

func (msg *Message) ShowMessage() string {
	switch msg.msgType {
	case ErrorMessage:
		return errStyle.Render(msg.value)
	case SuccessMessage:
		return successStyle.Render(msg.value)
	}

	return msgStyle.Render("")
}

func (msg Message) ClearMessage(t time.Duration) tea.Cmd {
	return tea.Tick(t, func(_ time.Time) tea.Msg {
		return ClearErrorMsg{}
	})
}
