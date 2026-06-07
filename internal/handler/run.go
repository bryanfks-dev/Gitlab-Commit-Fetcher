package handler

import (
	"context"
	"errors"
	"fmt"
	apperrors "main/errors"
)

func (h *handler) Run(ctx context.Context) error {
	tokenExistAndValid, err := h.svc.CheckGitlabPersonalAccessTokenExistAndValid(
		ctx,
	)
	tokenIsInvalid := errors.Is(err, apperrors.ErrInvalidGitlabPersonalToken)

	if err != nil && !tokenIsInvalid {
		return fmt.Errorf(
			"Error checking Gitlab access personal token existence and validity: %w",
			err,
		)
	}

	if !tokenExistAndValid {
		if tokenIsInvalid {
			h.logger.Printf("Gitlab personal token found but invalid.")
		} else {
			h.logger.Printf("Gitlab personal token not found.")
		}
		personalAccessTokenURL := fmt.Sprintf(
			"%s/-/user_settings/personal_access_tokens",
			h.gitlabBaseURL,
		)
		h.logger.Guidef(
			"Generate a new Gitlab personal access token within `%s` with `api`, `read_api`, `read_user`, and `read_repository` scopes checked.",
			personalAccessTokenURL,
		)
		h.askAndSaveForGitlabPersonalAccessTokenUntilValid(ctx)
	}

	return nil
}
