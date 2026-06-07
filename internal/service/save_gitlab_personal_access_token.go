package service

import (
	"context"
	"main/errors"
	vaultconsts "main/infra/vault/consts"
)

func (s *service) SaveGitlabPersonalAccessToken(
	ctx context.Context,
	token string,
) error {
	s.logger.Infof("Setting GitLab Personal Token in vault")
	tokenByte := []byte(token)
	ok, err := s.repo.ValidateGitlabPersonalToken(ctx, token)
	if err != nil {
		return err
	}
	if !ok {
		s.logger.Infof("GitLab Personal Token is invalid")
		return errors.ErrInvalidGitlabPersonalToken
	}
	return s.vaultRepo.Set(vaultconsts.GitlabPersonalTokenKey, tokenByte)
}
