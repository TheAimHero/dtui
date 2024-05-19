package managecontianer

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

func (m ContainerModel) View() string {
	doc := strings.Builder{}
	doc.WriteString(ui_table.BaseTableStyle.Render(m.table.View()))
	doc.WriteString("\n" + m.message.ShowMessage())
	doc.WriteString("\n" + m.help.View(m.keys))
  // @todo: adjust the fixHeight for consistency
	doc.WriteString(strings.Repeat("\n", ui_utils.HeightPadding(doc, 7)))
	return doc.String()
}
