package message

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Message struct {
	Value   string
	MsgType int
}

const (
	ErrorMessage = iota
	SuccessMessage
	InfoMessage
)

const (
	ErrorDuration   = 5 * time.Second
	SuccessDuration = 2 * time.Second
	InfoDuration    = 2 * time.Second
)

type ClearMessage struct{}

var (
	msgStyle     = lipgloss.NewStyle().Padding(1, 0)
	errStyle     = msgStyle.Foreground(lipgloss.Color("#cb4154"))
	successStyle = msgStyle.Foreground(lipgloss.Color("#AADB1E"))
	infoStyle    = msgStyle.Foreground(lipgloss.Color("#A4DBE8"))
)

func (msg *Message) AddMessage(value string, messageType int) {
	msg.MsgType = messageType
	msg.Value = value
}

func (msg *Message) ShowMessage() string {
	switch msg.MsgType {
	case ErrorMessage:
		return errStyle.Render(msg.Value)
	case SuccessMessage:
		return successStyle.Render(msg.Value)
	case InfoMessage:
		return infoStyle.Render(msg.Value)
	}

	return msgStyle.Render("")
}

func (msg Message) ClearMessage(t time.Duration) tea.Cmd {
	return tea.Tick(t, func(_ time.Time) tea.Msg {
		return ClearMessage{}
	})
}
