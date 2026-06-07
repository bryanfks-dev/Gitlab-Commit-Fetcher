package repository

import (
	"context"
	httpclient "main/infra/http_client"
	"main/infra/logger"
)

type repository struct {
	logger     logger.Logger
	httpClient httpclient.HttpClient

	gitlabAPIBaseURL string
}

type Repository interface {
	ValidateGitlabPersonalToken(ctx context.Context, token string) (bool, error)
}

func NewRepository(
	logger logger.Logger,
	httpClient httpclient.HttpClient,
	gitlabAPIBaseURL string,
) Repository {
	return &repository{
		logger:           logger,
		httpClient:       httpClient,
		gitlabAPIBaseURL: gitlabAPIBaseURL,
	}
}
