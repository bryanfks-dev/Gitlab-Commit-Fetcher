package service

import (
	"context"
	"main/core/errors"
	vaultconsts "main/infra/vault/consts"
)

func (s *service) SaveGitlabPersonalAccessToken(
	ctx context.Context,
	token string,
) error {
	s.logger.Infof("Setting GitLab personal access token in vault")
	tokenByte := []byte(token)
	ok, err := s.repo.ValidateGitlabPersonalAccessToken(ctx, token)
	if err != nil {
		return err
	}
	if !ok {
		s.logger.Infof("GitLab personal access token is invalid")
		return errors.ErrInvalidGitlabPersonalAccessToken
	}
	return s.vaultRepo.Set(vaultconsts.GitlabPersonalTokenKey, tokenByte)
}
