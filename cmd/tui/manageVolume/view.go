package managevolume

import (
	"github.com/TheAimHero/dtui/internal/ui/components"
)

func (m VolumeModel) View() string {
	vb := components.NewViewBuilder(m.Width, m.Height).
		AddCentered(m.Table.View()).
		AddCentered(m.Confirmation.View()).
		AddCentered(m.Message.ShowMessage()).
		AddCentered(m.Help.View(m.Keys))

	return vb.Build()
}
