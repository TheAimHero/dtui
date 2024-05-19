package utils

import (
	"bufio"
	"io"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type ResponseMsg string

func TickCommand() tea.Cmd {
	return tea.Tick(300*time.Millisecond, func(t time.Time) tea.Msg {
		return t
	})
}

func ListenToStream(sub chan<- ResponseMsg, stream io.ReadCloser) tea.Cmd {
	if stream == nil {
		return nil
	}
	return func() tea.Msg {
		defer stream.Close()
		scanner := bufio.NewScanner(stream)
		for scanner.Scan() {
			text := scanner.Text()
			sub <- ResponseMsg(text)
		}
		if err := scanner.Err(); err != nil {
			return nil
		}
		return nil
	}
}

func ResponseToStream(sub chan ResponseMsg) tea.Cmd {
	return func() tea.Msg {
		return ResponseMsg(<-sub)
	}
}
