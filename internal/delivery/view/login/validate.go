package loginview

import (
	"errors"
	"strings"
)

func (v *loginView) ValidateGitLabPersonalAccessToken(token string) error {
	token = strings.TrimSpace(token)
	if token == "" {
		return errors.New("Token cannot be empty")
	}
	return nil
}
