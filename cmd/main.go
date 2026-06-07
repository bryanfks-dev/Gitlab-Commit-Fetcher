package main

import (
	"bufio"
	"context"
	"main/config"
	httpclient "main/infra/http_client"
	"main/infra/logger"
	"main/infra/vault"
	vaultrepository "main/infra/vault/repository"
	"main/internal/handler"
	"main/internal/repository"
	"main/internal/service"
	"main/utils"
	"os"
)

var (
	DistributionUse string = "false"
)

func main() {
	cfg := config.LoadFromYAMLFile()
	logger := logger.NewLogger(utils.ForDistributionUse(DistributionUse))
	vault := vault.NewVault("gitlab-commit-fetcher")
	httpClient := httpclient.NewHttpClient()

	vaultRepo := vaultrepository.NewVaultRepository(logger, vault)
	repo := repository.NewRepository(
		logger,
		httpClient,
		cfg.Gitlab.GetGitlabAPIBaseURL(),
	)
	svc := service.NewService(logger, vaultRepo, repo)
	handler := handler.NewHandler(logger, svc, cfg.Gitlab.BaseURL)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := handler.Run(ctx); err != nil {
		logger.Printf(err.Error())
		logger.Printf("\nPress Enter to exit...")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
	}
}
