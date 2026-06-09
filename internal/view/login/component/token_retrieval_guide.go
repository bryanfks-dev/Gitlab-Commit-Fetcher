package logincomponent

import (
	"fmt"

	"charm.land/lipgloss/v2"
)

func GetTokenRetrievalGuide(gitlabBaseURL string) string {
	tokenURL := fmt.Sprintf(
		"%s/-/user_settings/personal_access_tokens",
		gitlabBaseURL,
	)
	msg := fmt.Sprintf(
		"Generate a new Gitlab personal access token within %s with api, read_api, read_user, and read_repository scopes checked.",
		tokenURL,
	)
	msg = lipgloss.JoinVertical(lipgloss.Left, "[GUIDE]", msg)
	return lipgloss.NewStyle().
		Foreground(lipgloss.Yellow).
		MarginBottom(1).
		SetString(msg).
		Render()
}
