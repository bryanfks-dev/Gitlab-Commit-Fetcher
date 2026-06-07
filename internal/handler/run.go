package handler

import (
	"context"
	"errors"
	"fmt"
	apperrors "main/core/errors"
)

func (h *handler) Run(ctx context.Context) error {
	tokenExistAndValid, err := h.svc.CheckGitlabPersonalAccessTokenExistAndValid(
		ctx,
	)
	tokenIsInvalid := errors.Is(
		err,
		apperrors.ErrInvalidGitlabPersonalAccessToken,
	)

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
		h.askAndSaveForGitlabPersonalAccessTokenUntilValid(ctx)
	}

	return nil
}
