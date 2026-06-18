package service

import (
	"context"
	"main/core/errors"
	vaultconsts "main/infra/vault/consts"
)

func (s *service) CheckGitlabPersonalAccessTokenExistAndValid(
	ctx context.Context,
) error {
	s.logger.Infof("Checking if GitLab personal access token exists in vault")
	token, err := s.vaultRepo.Get(vaultconsts.GitlabPersonalTokenKey)
	if err != nil {
		return err
	}

	if token == nil {
		s.logger.Infof("GitLab personal access token does not exist in vault")
		return errors.ErrGitLabPersonalAccessTokenNotFound
	}
	tokenStr := string(token)
	if tokenStr == "" {
		s.logger.Infof("GitLab personal access token is empty")
		return errors.ErrGitLabPersonalAccessTokenNotFound
	}

	s.logger.Infof("Checking if GitLab personal access token is valid")
	return s.repo.ValidateGitlabPersonalAccessToken(ctx, tokenStr)
}
