package handler

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"
)

func (h *handler) askAndSaveForGitlabPersonalAccessTokenUntilValid(
	ctx context.Context,
) {
	scanner := bufio.NewScanner(os.Stdin)
	var token string

	for {
		personalAccessTokenURL := fmt.Sprintf(
			"%s/-/user_settings/personal_access_tokens",
			h.gitlabBaseURL,
		)
		h.logger.Guidef(
			"Generate a new Gitlab personal access token within `%s` with `api`, `read_api`, `read_user`, and `read_repository` scopes checked.",
			personalAccessTokenURL,
		)
		h.logger.Notef(
			"Your Gitlab personal access token will be stored securely in you local machine secrets.",
		)

		h.logger.Print("Enter your Gitlab Personal Access Token: ")
		if !scanner.Scan() {
			h.logger.Errorf("Error reading token input: %v", scanner.Err())
			h.logger.Printf("Error reading token, try again")
			continue
		}

		token = strings.TrimSpace(scanner.Text())
		if err := h.validateGitlabPersonalToken(token); err != nil {
			h.logger.Printf(err.Error())
			continue
		}
		if err := h.svc.SaveGitlabPersonalAccessToken(ctx, token); err != nil {
			h.logger.Printf("Error saving token: %v\n", err)
			continue
		}
		break
	}
}
