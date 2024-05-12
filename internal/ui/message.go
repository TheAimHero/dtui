package ui

import "github.com/charmbracelet/lipgloss"

type Message struct {
	msgType string
	value   string
}

var (
	msgStyle     = lipgloss.NewStyle().Padding(1, 0)
	errStyle     = msgStyle.Copy().Foreground(lipgloss.Color("#cb4154"))
	successStyle = msgStyle.Copy().Foreground(lipgloss.Color("#AADB1E"))
)

func (msg *Message) AddMessage(value string, messageType string) {
	msg.msgType = messageType
	msg.value = value
}

func (msg *Message) ShowMessage() string {
	switch msg.msgType {
	case "error":
		return errStyle.Render(msg.value)
	case "success":
		return successStyle.Render(msg.value)
	}
	return msgStyle.Render("")
}
