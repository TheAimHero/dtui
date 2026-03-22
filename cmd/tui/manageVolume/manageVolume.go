package managevolume

import (
	"github.com/TheAimHero/dtui/internal/docker"
	"github.com/TheAimHero/dtui/internal/ui/components"
	"github.com/TheAimHero/dtui/internal/ui/message"
	"github.com/TheAimHero/dtui/internal/ui/prompt"
	"github.com/TheAimHero/dtui/internal/utils"
	tea "github.com/charmbracelet/bubbletea"
	mapset "github.com/deckarep/golang-set/v2"
)

type VolumeModel struct {
	*components.BaseModel
	VolumeSvc       docker.VolumeService
	Volumes         docker.Volumes
	Confirmation    prompt.Model
	Message         message.Message
	Keys            keyMap
	SelectedVolumes mapset.Set[string]
	InProgress      mapset.Set[string]
}

func (m VolumeModel) Init() tea.Cmd {
	return tea.Batch(utils.TickCommand(), m.Spinner.Tick)
}

func NewModel(volumeSvc docker.VolumeService) (VolumeModel, error) {
	volumes, err := volumeSvc.FetchVolumes()
	baseModel := components.NewBaseModel(80, 40)
	baseModel.Spinner = getSpinner()
	m := VolumeModel{
		BaseModel:       &baseModel,
		VolumeSvc:       volumeSvc,
		Volumes:         volumes,
		SelectedVolumes: mapset.NewSet[string](),
		InProgress:      mapset.NewSet[string](),
		Message:         message.Message{},
		Keys:            keys,
	}
	m.Table = m.getTable()
	if err != nil {
		return m, components.ErrorMessage(err.Error())
	}
	return m, nil
}
