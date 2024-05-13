package manageimage

import (
	"github.com/TheAimHero/dtui/internal/ui"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	ImageStatus = iota
	ImageID
	ImageTag
	ImageCreated
	ImageSize
)

func (m imageModel) DeleteImage() (tea.Model, tea.Cmd) {
	err := m.dockerClient.DeleteImage()
	if err != nil {
		m.message.AddMessage(err.Error(), ui.ErrorMessage)
		return m, m.message.ClearMessage(errorDuration)
	}
	m.message.AddMessage("Image deleted", ui.SuccessMessage)
	return m, m.message.ClearMessage(successDuration)
}

func (m imageModel) DeleteImages() (tea.Model, tea.Cmd) {
	errors := m.dockerClient.DeleteImages(m.selectedImages.ToSlice())
	if len(errors) > 0 {
		m.message.AddMessage("Error while deleting some images", ui.ErrorMessage)
		m.selectedImages.Clear()
		return m, m.message.ClearMessage(errorDuration)
	}
	m.message.AddMessage("Images deleted", ui.SuccessMessage)
	m.selectedImages.Clear()
	return m, m.message.ClearMessage(successDuration)
}

func (m imageModel) SelectImage() (tea.Model, tea.Cmd) {
	if len(m.table.Rows()) == 0 {
		return m, nil
	}
	imageID := m.table.SelectedRow()[ImageID]
	if m.selectedImages.Contains(imageID) {
		m.selectedImages.Remove(imageID)
	} else {
		m.selectedImages.Add(imageID)
	}
	m.table.MoveDown(1)
	return m, nil
}

func (m imageModel) SelectAllImages() (tea.Model, tea.Cmd) {
	var allIDs []string
	for _, row := range m.table.Rows() {
		allIDs = append(allIDs, row[ImageID])
	}
	if m.selectedImages.Cardinality() == len(m.table.Rows()) {
		m.selectedImages.Clear()
	} else {
		m.selectedImages.Clear()
		m.selectedImages.Append(allIDs...)
	}
	return m, nil
}
