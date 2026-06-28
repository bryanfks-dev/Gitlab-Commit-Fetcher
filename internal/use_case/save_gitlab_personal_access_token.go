package usecase

import (
	"context"
	vaultconsts "main/infra/vault/consts"
)

func (uc *useCase) SaveGitlabPersonalAccessToken(
	ctx context.Context,
	token string,
) error {
	uc.logger.Infof("Setting GitLab personal access token in vault")
	tokenByte := []byte(token)
	err := uc.svc.ValidateGitlabPersonalAccessToken(ctx, token)
	if err != nil {
		return err
	}
	return uc.vaultRepo.Set(vaultconsts.GitlabPersonalTokenKey, tokenByte)
}
