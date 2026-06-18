package loginview

import (
	"fmt"
	"main/core/consts"
	"strings"

	"github.com/mitchellh/go-wordwrap"
)

func (v *loginView) newGenerateGitLabPersonalAccessTokenGuideMessage() string {
	gitlabPersonalAccessTokenURL := fmt.Sprintf(
		"%s/-/user_settings/personal_access_tokens",
		v.gitlabBaseURL,
	)
	scopes := []string{
		"api",
		"read_api",
		"read_user",
		"read_repository",
	}
	var scopesStr strings.Builder
	for i, scope := range scopes {
		if i == len(scopes)-1 {
			fmt.Fprintf(&scopesStr, "and `%s`", scope)
			continue
		}
		fmt.Fprintf(&scopesStr, "`%s`, ", scope)
	}

	msg := fmt.Sprintf(
		"Generate a new Gitlab personal access token within `%s` with %s scopes checked.",
		gitlabPersonalAccessTokenURL,
		scopesStr.String(),
	)
	return wordwrap.WrapString(msg, consts.WindowMaxWidth)
}
