package logincomponent

import (
	"charm.land/lipgloss/v2"
)

func GetInformation() string {
	return lipgloss.NewStyle().
		Foreground(lipgloss.Green).
		MarginTop(1).
		SetString("Your Gitlab personal access token will be stored securely in you local machine secrets.").
		Render()
}
