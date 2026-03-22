package managecontainer

import (
	tea "github.com/charmbracelet/bubbletea"
	mapset "github.com/deckarep/golang-set/v2"

	"github.com/TheAimHero/dtui/internal/docker"
	"github.com/TheAimHero/dtui/internal/ui/components"
	"github.com/TheAimHero/dtui/internal/ui/message"
	"github.com/TheAimHero/dtui/internal/ui/prompt"
	"github.com/TheAimHero/dtui/internal/utils"
)

type ContainerModel struct {
	*components.BaseModel
	ContainerSvc       docker.ContainerService
	Containers         docker.Containers
	SelectedContainers mapset.Set[string]
	InProcess          mapset.Set[string]
	Confirmation       prompt.Model
	Message            message.Message
	containerKeys      keyMap
}

func (m ContainerModel) Init() tea.Cmd {
	return tea.Batch(utils.TickCommand(), m.Spinner.Tick)
}

func NewModel(containerSvc docker.ContainerService) (ContainerModel, error) {
	containers, err := containerSvc.FetchContainers()
	if err != nil {
		return ContainerModel{}, err
	}
	base := components.NewBaseModel(80, 40)
	m := ContainerModel{
		BaseModel:          &base,
		ContainerSvc:       containerSvc,
		Containers:         containers,
		SelectedContainers: mapset.NewSet[string](),
		InProcess:          mapset.NewSet[string](),
		Message:            message.Message{},
		containerKeys:      keys,
	}
	m.Table = m.getTable()
	return m, nil
}
