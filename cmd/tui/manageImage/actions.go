package manageimage

import (
	"fmt"
	"time"

	"github.com/TheAimHero/dtui/internal/ui/message"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	ImageStatus = iota
	ImageLoading
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

func (m ImageModel) PruneImages() (ImageModel, tea.Cmd) {
	err := m.DockerClient.PruneImage()
	if err != nil {
		m.Message.AddMessage("Error while pruning some images", message.ErrorMessage)
		m.SelectedImages.Clear()
		return m, m.Message.ClearMessage(errorDuration)
	}
	m.Message.AddMessage("Images pruned", message.SuccessMessage)
	m.Table.SetCursor(0)
	return m, m.Message.ClearMessage(successDuration)
}

func (m ImageModel) PullImage() (ImageModel, tea.Cmd) {
	imageName := m.Input.Value()
	m.Input.SetValue("")
	if imageName == "" {
		m.Message.AddMessage("Image name cannot be empty", message.InfoMessage)
		return m, m.Message.ClearMessage(message.InfoDuration)
	}
	pullMsg := message.Message{}
	return m, func() tea.Msg {
		m.PullProgress.Add(imageName)
		curTime := time.Now()
		_, err := m.DockerClient.PullImage(imageName)
		if err != nil {
			pullMsg.AddMessage("Error while pulling image: "+imageName, message.ErrorMessage)
			m.PullProgress.Remove(imageName)
			return pullMsg
		}
		m.PullProgress.Remove(imageName)
		pullMsg.AddMessage(fmt.Sprintf("Image %s pulled in %s", imageName, time.Since(curTime)), message.SuccessMessage)
		return pullMsg
	}
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
