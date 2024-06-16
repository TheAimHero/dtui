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

func deleteImage(m ImageModel) (ImageModel, tea.Cmd) {
	row := m.Table.SelectedRow()
	if row == nil {
		m.Message.AddMessage("No image to delete", message.InfoMessage)
		return m, m.Message.ClearMessage(message.InfoDuration)
	}

	err := m.DockerClient.DeleteImage(row[ImageID])
	if err != nil {
		m.Message.AddMessage(err.Error(), message.ErrorMessage)
		return m, m.Message.ClearMessage(errorDuration)
	}
	m.Message.AddMessage("Image deleted", message.SuccessMessage)
	return m, m.Message.ClearMessage(successDuration)
}

func (m ImageModel) DeleteImages() (ImageModel, tea.Cmd) {
	var errors []string
	if len(m.SelectedImages.ToSlice()) == 0 {
		return deleteImage(m)
	}
	for _, imageID := range m.SelectedImages.ToSlice() {
		err := m.DockerClient.DeleteImage(imageID)
		if err != nil {
			errors = append(errors, err.Error())
		}
	}
	if len(errors) > 0 {
		m.Message.AddMessage("Error while deleting some images", message.ErrorMessage)
		m.SelectedImages.Clear()
		return m, m.Message.ClearMessage(errorDuration)
	}
	m.Message.AddMessage("Images deleted", message.SuccessMessage)
	m.SelectedImages.Clear()
	return m, m.Message.ClearMessage(successDuration)
}

func (m *ImageModel) PullImages(imageName string) (ImageModel, tea.Cmd, io.ReadCloser) {
	var (
		stream io.ReadCloser
		err    error
	)
	// @fix: this causes tui to become unresponsive for a while
	stream, err = m.DockerClient.PullImage(imageName)
	m.Text = []string{}
	if err != nil {
		m.Message.AddMessage(err.Error(), message.ErrorMessage)
		return *m, m.Message.ClearMessage(errorDuration), stream
	}
	m.Message.AddMessage("Image pulled successfully", message.SuccessMessage)
	return *m, m.Message.ClearMessage(successDuration), stream
}

func (m ImageModel) SelectImage() (ImageModel, tea.Cmd) {
	if len(m.Table.Rows()) == 0 {
		return m, nil
	}
	imageID := m.Table.SelectedRow()[ImageID]
	if m.SelectedImages.Contains(imageID) {
		m.SelectedImages.Remove(imageID)
	} else {
		m.SelectedImages.Add(imageID)
	}
	m.Table.MoveDown(1)
	return m, nil
}

func (m ImageModel) SelectAllImages() (ImageModel, tea.Cmd) {
	var allIDs []string
	for _, row := range m.Table.Rows() {
		allIDs = append(allIDs, row[ImageID])
	}
	if m.SelectedImages.Cardinality() == len(m.Table.Rows()) {
		m.SelectedImages.Clear()
	} else {
		m.SelectedImages.Clear()
		m.SelectedImages.Append(allIDs...)
	}
	return m, nil
}
