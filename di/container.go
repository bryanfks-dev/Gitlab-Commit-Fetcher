package di

import (
	"context"
	"main/core/config"
	httpclient "main/infra/http_client"
	"main/infra/logger"
	"main/infra/vault"
	vaultrepository "main/infra/vault/repository"
	"main/internal/delivery"
	"main/internal/repository"
	"main/internal/service"
)

type Container struct {
	Ctx        context.Context
	Cfg        *config.Config
	Logger     logger.Logger
	Vault      *vault.Vault
	HttpClient httpclient.HttpClient
	VaultRepo  vaultrepository.VaultRepository
	Repo       repository.Repository
	Svc        service.Service
	App        delivery.App
}

func NewContainer(
	ctx context.Context,
) *Container {
	cfg := config.LoadFromYAMLFile()
	logger := logger.NewLogger()
	vault := vault.NewVault("gitlab-commit-fetcher")
	httpClient := httpclient.NewHttpClient()

	vaultRepo := vaultrepository.NewVaultRepository(logger, vault)
	repo := repository.NewRepository(
		logger,
		httpClient,
		cfg.Gitlab.GetGitlabAPIBaseURL(),
	)
	svc := service.NewService(logger, cfg, vaultRepo, repo)
	app := delivery.NewApp(ctx, svc, cfg)

	return &Container{
		Ctx:        ctx,
		Cfg:        cfg,
		Logger:     logger,
		Vault:      vault,
		HttpClient: httpClient,
		VaultRepo:  vaultRepo,
		Repo:       repo,
		Svc:        svc,
		App:        app,
	}
}
