package manageimage

import (
	"io"

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

func (m ImageModel) DeleteImage() (ImageModel, tea.Cmd) {
	row := m.Table.SelectedRow()
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

func (m ImageModel) DeleteImages() (ImageModel, tea.Cmd) {
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
	if len(errors) > 0 {
		m.message.AddMessage("Error while deleting some images", message.ErrorMessage)
		m.selectedImages.Clear()
		return m, m.message.ClearMessage(errorDuration)
	}
	m.message.AddMessage("Images deleted", message.SuccessMessage)
	m.selectedImages.Clear()
	return m, m.message.ClearMessage(successDuration)
}

func (m *ImageModel) PullImages(imageName string) (ImageModel, tea.Cmd, io.ReadCloser) {
	var (
		stream io.ReadCloser
		err    error
	)
	// @fix: this causes tui to become unresponsive for a while
	stream, err = m.dockerClient.PullImage(imageName)
	m.text = []string{}
	if err != nil {
		m.message.AddMessage(err.Error(), message.ErrorMessage)
		return *m, m.message.ClearMessage(errorDuration), stream
	}
	m.message.AddMessage("Image pulled successfully", message.SuccessMessage)
	return *m, m.message.ClearMessage(successDuration), stream
}

func (m ImageModel) SelectImage() (ImageModel, tea.Cmd) {
	if len(m.Table.Rows()) == 0 {
		return m, nil
	}
	imageID := m.Table.SelectedRow()[ImageID]
	if m.selectedImages.Contains(imageID) {
		m.selectedImages.Remove(imageID)
	} else {
		m.selectedImages.Add(imageID)
	}
	m.Table.MoveDown(1)
	return m, nil
}

func (m ImageModel) SelectAllImages() (ImageModel, tea.Cmd) {
	var allIDs []string
	for _, row := range m.Table.Rows() {
		allIDs = append(allIDs, row[ImageID])
	}
	if m.selectedImages.Cardinality() == len(m.Table.Rows()) {
		m.selectedImages.Clear()
	} else {
		m.selectedImages.Clear()
		m.selectedImages.Append(allIDs...)
	}
	return m, nil
}
