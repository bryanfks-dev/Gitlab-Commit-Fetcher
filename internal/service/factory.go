package service

import (
	"context"
	"main/infra/logger"
	vaultrepository "main/infra/vault/repository"
	"main/internal/repository"
)

type service struct {
	logger    logger.Logger
	vaultRepo vaultrepository.VaultRepository
	repo      repository.Repository
}

type Service interface {
	SaveGitlabPersonalAccessToken(ctx context.Context, token string) error
	CheckGitlabPersonalAccessTokenExistAndValid(
		ctx context.Context,
	) (bool, error)
}

func NewService(
	logger logger.Logger,
	vaultRepo vaultrepository.VaultRepository,
	repo repository.Repository,
) Service {
	return &service{
		logger:    logger,
		vaultRepo: vaultRepo,
		repo:      repo,
	}
}
