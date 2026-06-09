package di

import (
	"context"
	"main/core/config"
	"main/core/utils"
	httpclient "main/infra/http_client"
	"main/infra/logger"
	"main/infra/vault"
	vaultrepository "main/infra/vault/repository"
	"main/internal/repository"
	"main/internal/service"
	loginview "main/internal/view/login"
	mainview "main/internal/view/main"
	rootview "main/internal/view/root"
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
	RootVM     rootview.RootViewModel
	LoginVM    loginview.LoginViewModel
	MainVM     mainview.MainViewModel
}

func NewContainer(
	ctx context.Context,
	distibutionUseEnabled string,
) *Container {
	forDistributionUse := utils.ForDistributionUse(distibutionUseEnabled)

	cfg := config.LoadFromYAMLFile()
	logger := logger.NewLogger(forDistributionUse)
	vault := vault.NewVault("gitlab-commit-fetcher")
	httpClient := httpclient.NewHttpClient()

	vaultRepo := vaultrepository.NewVaultRepository(logger, vault)
	repo := repository.NewRepository(
		logger,
		httpClient,
		cfg.Gitlab.GetGitlabAPIBaseURL(),
	)
	svc := service.NewService(logger, cfg, vaultRepo, repo)

	loginVM := loginview.NewLoginViewModel(ctx, cfg.Gitlab.BaseURL, svc)
	mainVM := mainview.NewMainViewModel(ctx, svc)
	rootVM := rootview.NewRootViewModel(ctx, svc, loginVM, mainVM)

	return &Container{
		Ctx:        ctx,
		Cfg:        cfg,
		Logger:     logger,
		Vault:      vault,
		HttpClient: httpClient,
		VaultRepo:  vaultRepo,
		Repo:       repo,
		Svc:        svc,
		RootVM:     rootVM,
		LoginVM:    loginVM,
		MainVM:     mainVM,
	}
}
