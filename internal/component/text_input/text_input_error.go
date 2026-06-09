package textinputcomponent

import (
	"charm.land/lipgloss/v2"
)

func newTextInputError(err error) string {
	return lipgloss.NewStyle().
		Foreground(lipgloss.Red).
		SetString(err.Error()).
		Render()
}
