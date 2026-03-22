package components

import (
	"github.com/charmbracelet/bubbles/help"
)

func NewHelpModel(styles HelpStyles) help.Model {
	m := help.New()
	s := m.Styles
	s.ShortDesc = styles.DescStyle
	s.FullDesc = styles.DescStyle
	s.FullKey = styles.KeyStyle
	s.ShortKey = styles.KeyStyle
	s.Ellipsis = styles.EllipsisStyle
	s.FullSeparator = styles.EllipsisStyle
	s.ShortSeparator = styles.EllipsisStyle
	m.Styles = s
	return m
}

func NewDefaultHelpModel() help.Model {
	return NewHelpModel(DefaultHelpStyles())
}
