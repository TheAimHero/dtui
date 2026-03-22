package components

import (
	"github.com/charmbracelet/bubbles/textinput"
)

// NewFilterInput creates a text input for filtering
func NewFilterInput(prompt, placeholder string) textinput.Model {
	ti := textinput.New()
	ti.Placeholder = placeholder
	ti.Prompt = prompt
	ti.Focus()
	return ti
}

// NewPullInput creates a text input for pull operations
func NewPullInput() textinput.Model {
	ti := textinput.New()
	ti.Placeholder = "Image Name"
	ti.Prompt = "Image Pull Name: "
	ti.Focus()
	return ti
}
