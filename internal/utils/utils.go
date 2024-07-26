package utils

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"time"

	"github.com/araddon/dateparse"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
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

func GetSize(bytes int64) string {
	size := float64(bytes)
	unit := ""
	switch {
	case size < KB:
		unit = "B"
	case size < MB:
		size /= KB
		unit = "KB"
	case size < GB:
		size /= MB
		unit = "MB"
	case size < TB:
		size /= GB
		unit = "GB"
	case size < PB:
		size /= TB
		unit = "TB"
	case size < EB:
		size /= PB
		unit = "PB"
	default:
		size /= EB
		unit = "EB"
	}
	return fmt.Sprintf("%.2f%s", size, unit)
}

func GetDate(dateStr string) string {
	date, _ := dateparse.ParseAny(dateStr)
	return date.Format("02/01/2006 15:04 MST")
}

func FloorMul(n int, m float64) int {
	return int(math.Floor(float64(n) * m))
}
