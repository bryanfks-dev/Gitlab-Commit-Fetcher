package service

import (
	"context"
	"main/core/config"
	"main/infra/logger"
	vaultrepository "main/infra/vault/repository"
	"main/internal/repository"
)

type service struct {
	logger    logger.Logger
	cfg       *config.Config
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
	cfg *config.Config,
	vaultRepo vaultrepository.VaultRepository,
	repo repository.Repository,
) Service {
	return &service{
		logger:    logger,
		cfg:       cfg,
		vaultRepo: vaultRepo,
		repo:      repo,
	}
}
