package service

import (
	"context"
	vaultconsts "main/infra/vault/consts"
)

func (s *service) SaveGitlabPersonalAccessToken(
	ctx context.Context,
	token string,
) error {
	s.logger.Infof("Setting GitLab personal access token in vault")
	tokenByte := []byte(token)
	err := s.repo.ValidateGitlabPersonalAccessToken(ctx, token)
	if err != nil {
		return err
	}
	return s.vaultRepo.Set(vaultconsts.GitlabPersonalTokenKey, tokenByte)
}
