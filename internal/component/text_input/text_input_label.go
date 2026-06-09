package textinputcomponent

import "charm.land/lipgloss/v2"

func newTextInputLabel(label string) string {
	return lipgloss.NewStyle().
		Foreground(lipgloss.White).
		SetString(label).
		Render()
}
