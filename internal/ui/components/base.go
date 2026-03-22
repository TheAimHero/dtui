package components

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textinput"
)

// BaseModel holds fields common to all management models
type BaseModel struct {
	Help    help.Model
	Table   table.Model
	Input   textinput.Model
	Keys    NavigationKeys
	Spinner spinner.Model
	Width   int
	Height  int
}

// NewBaseModel creates a new BaseModel with default values
func NewBaseModel(width, height int) BaseModel {
	return BaseModel{
		Keys:   NewNavigationKeys(),
		Help:   NewDefaultHelpModel(),
		Width:  width,
		Height: height,
	}
}
