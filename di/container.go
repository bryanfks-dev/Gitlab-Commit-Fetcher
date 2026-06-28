package di

import (
	"context"
	"main/core/config"
	httpclient "main/infra/http_client"
	"main/infra/logger"
	"main/infra/vault"
	vaultrepository "main/infra/vault/repository"
	"main/internal/delivery"
	"main/internal/service"
	usecase "main/internal/use_case"
)

type Container struct {
	Ctx        context.Context
	Cfg        *config.Config
	Logger     logger.Logger
	Vault      *vault.Vault
	HttpClient httpclient.HttpClient
	VaultRepo  vaultrepository.VaultRepository
	UC         usecase.UseCase
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
	svc := service.NewService(
		logger,
		httpClient,
		cfg.Gitlab.GetGitlabAPIBaseURL(),
	)
	uc := usecase.NewUseCase(logger, cfg, vaultRepo, svc)
	app := delivery.NewApp(ctx, uc, cfg)

	return &Container{
		Ctx:        ctx,
		Cfg:        cfg,
		Logger:     logger,
		Vault:      vault,
		HttpClient: httpClient,
		VaultRepo:  vaultRepo,
		UC:         uc,
		Svc:        svc,
		App:        app,
	}
}
