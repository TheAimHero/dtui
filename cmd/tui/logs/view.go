package logs

import (
	"os"
	"strings"

	"github.com/TheAimHero/dtui/internal/ui/table"
	ui_utils "github.com/TheAimHero/dtui/internal/ui/utils"
	"golang.org/x/term"
)

var (
	physicalWidth, _, _ = term.GetSize(int(os.Stdout.Fd()))
)

func (m LogModel) View() string {
	doc := strings.Builder{}
	doc.WriteString(table.BaseTableStyle.Copy().Render(m.Table.View()))
	doc.WriteString("\n" + m.Message.ShowMessage())
	doc.WriteString("\n" + m.Help.View(m.Keys))
	doc.WriteString(strings.Repeat("\n", ui_utils.HeightPadding(doc, 8)))
	return doc.String()
}
