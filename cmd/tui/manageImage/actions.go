package manageimage

import (
	"github.com/TheAimHero/dtui/internal/ui/message"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	ImageStatus = iota
	ImageID
	ImageTag
	ImageCreated
	ImageSize
)

func (m imageModel) DeleteImage() (imageModel, tea.Cmd) {
	row := m.table.SelectedRow()
	if row == nil {
		m.message.AddMessage("No image selected", message.ErrorMessage)
		return m, m.message.ClearMessage(errorDuration)
	}

	err := m.dockerClient.DeleteImage(row[ImageID])
	if err != nil {
		m.message.AddMessage(err.Error(), message.ErrorMessage)
		return m, m.message.ClearMessage(errorDuration)
	}
	m.message.AddMessage("Image deleted", message.SuccessMessage)
	return m, m.message.ClearMessage(successDuration)
}

func (m imageModel) DeleteImages() (imageModel, tea.Cmd) {
	var errors []string
	if len(m.selectedImages.ToSlice()) == 0 {
		m.message.AddMessage("No images selected", message.ErrorMessage)
		return m, m.message.ClearMessage(errorDuration)
	}
	for _, imageID := range m.selectedImages.ToSlice() {
		err := m.dockerClient.DeleteImage(imageID)
		if err != nil {
			errors = append(errors, err.Error())
		}
	}
	// errors := m.dockerClient.DeleteImages(m.selectedImages.ToSlice())
	if len(errors) > 0 {
		m.message.AddMessage("Error while deleting some images", message.ErrorMessage)
		m.selectedImages.Clear()
		return m, m.message.ClearMessage(errorDuration)
	}
	m.message.AddMessage("Images deleted", message.SuccessMessage)
	m.selectedImages.Clear()
	return m, m.message.ClearMessage(successDuration)
}

func (m imageModel) SelectImage() (imageModel, tea.Cmd) {
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

func (m imageModel) SelectAllImages() (imageModel, tea.Cmd) {
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
