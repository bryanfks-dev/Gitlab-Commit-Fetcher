package usecase

import (
	"context"
	"main/core/config"
	"main/data/model"
	"main/infra/logger"
	vaultrepository "main/infra/vault/repository"
	"main/internal/service"
)

type useCase struct {
	logger    logger.Logger
	cfg       *config.Config
	vaultRepo vaultrepository.VaultRepository
	svc       service.Service
}

type UseCase interface {
	SaveGitlabPersonalAccessToken(ctx context.Context, token string) error
	CheckGitlabPersonalAccessTokenExistAndValid(
		ctx context.Context,
	) error
	GetRegisteredGitlabProjects() []model.ProjectModel
	RegisterAndCheckProjectExistAndValid(
		ctx context.Context,
		projectID int,
	) error
	RemoveProject(ctx context.Context, projectID int) error
}

func NewUseCase(
	logger logger.Logger,
	cfg *config.Config,
	vaultRepo vaultrepository.VaultRepository,
	svc service.Service,
) UseCase {
	return &useCase{
		logger:    logger,
		cfg:       cfg,
		vaultRepo: vaultRepo,
		svc:       svc,
	}
}
