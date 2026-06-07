package service

import (
	"context"
	vaultconsts "main/infra/vault/consts"
)

func (s *service) CheckGitlabPersonalAccessTokenExistAndValid(
	ctx context.Context,
) (bool, error) {
	s.logger.Infof("Checking if GitLab personal access token exists in vault")
	token, err := s.vaultRepo.Get(vaultconsts.GitlabPersonalTokenKey)
	if err != nil || token == nil {
		s.logger.Infof("GitLab personal access token does not exist in vault")
		return false, err
	}
	tokenStr := string(token)
	if tokenStr == "" {
		s.logger.Infof("GitLab personal access token is empty")
		return false, nil
	}

	s.logger.Infof("Checking if GitLab personal access token is valid")
	return s.repo.ValidateGitlabPersonalAccessToken(ctx, tokenStr)
}
