package manageimage

import (
	"fmt"
	"strings"
	"time"

	"github.com/TheAimHero/dtui/internal/ui/components"
)

type ShowTextInput struct{}

const (
	successDuration = 2 * time.Second
	errorDuration   = 5 * time.Second
)

func (m ImageModel) pullImage() string {
	var lines []string
	m.PullProgress.Range(func(key, value any) bool {
		imageName := key.(string)
		info := value.(PullProgressInfo)
		line := fmt.Sprintf("%s %s", m.PullSpinner.View(), imageName)
		if info.Progress != nil && info.Progress.Total > 0 {
			percent := float64(info.Progress.Current) / float64(info.Progress.Total) * 100
			line += fmt.Sprintf(": %s (%.0f%%)", info.Status, percent)
		} else {
			line += fmt.Sprintf(": %s", info.Status)
		}
		lines = append(lines, line)
		return true
	})
	return strings.Join(lines, "\n")
}

func (m ImageModel) View() string {
	vb := components.NewViewBuilder(m.BaseModel.Width, m.BaseModel.Height).
		AddCentered(m.Table.View())

	if m.Input.Focused() || m.Input.Value() != "" {
		vb.AddPadded(m.Input.View())
	} else {
		vb.AddSpacing(2)
	}

	vb.AddCentered(m.Confirmation.View()).
		AddCentered(m.Message.ShowMessage()).
		AddCentered(m.Help.View(m.Keys))

	count := 0
	m.PullProgress.Range(func(_, _ any) bool {
		count++
		return false
	})
	if count > 0 {
		vb.AddCentered(m.pullImage())
	} else {
		vb.AddSpacing(1)
	}

	return vb.Build()
}
