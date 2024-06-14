package managevolume

import (
	"os"
	"strings"

	ui_table "github.com/TheAimHero/dtui/internal/ui/table"
	ui_utils "github.com/TheAimHero/dtui/internal/ui/utils"
	"golang.org/x/term"
)

var (
	physicalWidth, physicalHeight, _ = term.GetSize(int(os.Stdout.Fd())) // nolint:unused
)

func (m VolumeModel) View() string {
	doc := strings.Builder{}
	doc.WriteString(ui_table.BaseTableStyle.Render(m.Table.View()))
	doc.WriteString("\n" + m.Message.ShowMessage())
	doc.WriteString("\n" + m.Help.View(m.Keys))
	doc.WriteString(strings.Repeat("\n", ui_utils.HeightPadding(doc, 8)))
	return doc.String()
}
