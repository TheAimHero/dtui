package manageimage

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"time"

	"golang.org/x/term"

	ui_table "github.com/TheAimHero/dtui/internal/ui/table"
	ui_utils "github.com/TheAimHero/dtui/internal/ui/utils"
	"github.com/charmbracelet/lipgloss"
)

type ShowTextInput struct{}

var (
	physicalWidth, physicalHeight, _ = term.GetSize(int(os.Stdout.Fd())) // nolint:unused
	successDuration                  = 2 * time.Second
	errorDuration                    = 5 * time.Second
)

func (m ImageModel) pullImage() string {
	images := m.PullProgress.ToSlice()
	sort.SliceStable(images, func(i, j int) bool { return images[i] > images[j] })
	return fmt.Sprintf("%s\tPulling images: %s", m.PullSpinner.View(), strings.Join(images, ", "))
}

func (m ImageModel) View() string {
	doc := strings.Builder{}
	doc.WriteString(ui_table.BaseTableStyle.Render(m.Table.View()))
	if m.Input.Focused() || m.Input.Value() != "" {
		doc.WriteString("\n" + lipgloss.NewStyle().Padding(1, 0, 0, 0).Render(m.Input.View()))
	} else {
		doc.WriteString(strings.Repeat("\n", 2))
	}
	doc.WriteString("\n" + m.Message.ShowMessage())
	doc.WriteString("\n" + m.Help.View(m.Keys))
	if m.PullProgress.Cardinality() > 0 {
		doc.WriteString("\n" + lipgloss.PlaceVertical(physicalHeight-lipgloss.Height(doc.String())-10, lipgloss.Bottom, m.pullImage()))
	} else {
		doc.WriteString("\n")
	}
	doc.WriteString(strings.Repeat("\n", ui_utils.HeightPadding(doc, 8)))
	return doc.String()
}
