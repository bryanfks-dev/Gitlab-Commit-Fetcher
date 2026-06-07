package service

import (
	"context"
	vaultconsts "main/infra/vault/consts"
)

func (s *service) CheckGitlabPersonalAccessTokenExistAndValid(
	ctx context.Context,
) (bool, error) {
	s.logger.Infof("Checking if GitLab personal access token exists in vault")
	val, err := s.vaultRepo.Get(vaultconsts.GitlabPersonalTokenKey)
	if err != nil || val == nil {
		return false, err
	}
	token := string(val)
	if token == "" {
		return false, nil
	}

	s.logger.Infof("Checking if GitLab personal access token is valid")
	return s.repo.ValidateGitlabPersonalToken(ctx, token)
}
