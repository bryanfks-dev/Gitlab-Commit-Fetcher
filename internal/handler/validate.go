package handler

import "fmt"

func (h *handler) validateGitlabPersonalToken(token string) error {
	if token == "" {
		return fmt.Errorf("Token cannot be empty, try again")
	}
	return nil
}
