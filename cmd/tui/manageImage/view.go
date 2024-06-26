package manageimage

import (
	"os"
	"strings"
	"time"

	"golang.org/x/term"

	ui_table "github.com/TheAimHero/dtui/internal/ui/table"
	ui_utils "github.com/TheAimHero/dtui/internal/ui/utils"
)

type ShowTextInput struct{}

var (
	physicalWidth, physicalHeight, _ = term.GetSize(int(os.Stdout.Fd())) // nolint:unused
	successDuration                  = 2 * time.Second
	errorDuration                    = 5 * time.Second
)

func (m ImageModel) View() string {
	doc := strings.Builder{}
	doc.WriteString(ui_table.BaseTableStyle.Render(m.Table.View()))
	doc.WriteString("\n" + m.Message.ShowMessage())
	doc.WriteString("\n" + m.Help.View(m.Keys))
	doc.WriteString(strings.Repeat("\n", ui_utils.HeightPadding(doc, 8)))
	return doc.String()
}
