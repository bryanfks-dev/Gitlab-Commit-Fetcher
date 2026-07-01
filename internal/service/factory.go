package service

import (
	"context"
	"main/data/dto"
	"main/data/model"
	httpclient "main/infra/http_client"
	"main/infra/logger"
)

type service struct {
	logger     logger.Logger
	httpClient httpclient.HttpClient

	gitlabAPIBaseURL string
}

type Service interface {
	ValidateGitlabPersonalAccessToken(ctx context.Context, token string) error
	GetGitlabRefComparation(
		query *dto.GetGitlabRepositoryRefComparationQuery,
	) (*dto.RefComparationResultDTO, error)
	GetGitlabProject(
		ctx context.Context,
		projectID int,
		token string,
	) (*model.ProjectModel, error)

	CopyTextToClipboard(ctx context.Context, text string)
}

func NewService(
	logger logger.Logger,
	httpClient httpclient.HttpClient,
	gitlabAPIBaseURL string,
) Service {
	return &service{
		logger:           logger,
		httpClient:       httpClient,
		gitlabAPIBaseURL: gitlabAPIBaseURL,
	}
}
