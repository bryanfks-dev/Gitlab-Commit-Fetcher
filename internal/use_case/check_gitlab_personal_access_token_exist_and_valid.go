package usecase

import (
	"context"
	"main/core/errors"
	vaultconsts "main/infra/vault/consts"
)

func (uc *useCase) CheckGitlabPersonalAccessTokenExistAndValid(
	ctx context.Context,
) error {
	uc.logger.Infof("Checking if GitLab personal access token exists in vault")
	token, err := uc.vaultRepo.Get(vaultconsts.GitlabPersonalTokenKey)
	if err != nil {
		return err
	}

	if token == nil {
		uc.logger.Infof("GitLab personal access token does not exist in vault")
		return errors.ErrGitLabPersonalAccessTokenNotFound
	}
	tokenStr := string(token)
	if tokenStr == "" {
		uc.logger.Infof("GitLab personal access token is empty")
		return errors.ErrGitLabPersonalAccessTokenNotFound
	}

	uc.logger.Infof("Checking if GitLab personal access token is valid")
	return uc.svc.ValidateGitlabPersonalAccessToken(ctx, tokenStr)
}
