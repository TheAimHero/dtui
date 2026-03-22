package managecontainer

import (
	"github.com/TheAimHero/dtui/internal/ui/components"
)

func (m ContainerModel) View() string {
	vb := components.NewViewBuilder(m.Width, m.Height).
		AddCentered(m.Table.View()).
		AddCentered(m.Confirmation.View())

	if m.Input.Focused() || m.Input.Value() != "" {
		vb.AddPadded(m.Input.View())
	} else {
		vb.AddSpacing(2)
	}

	vb.AddCentered(m.Message.ShowMessage()).
		AddCentered(m.Help.View(m.Keys))

	return vb.Build()
}
