package logincomponent

import (
	"image/color"

	"charm.land/lipgloss/v2"
)

func GetInputGuide() string {
	return lipgloss.NewStyle().
		Foreground(color.White).
		MarginTop(2).
		SetString("Press Enter to submit.").
		Render()
}
