package repository

import (
	"context"
	"main/data/dto"
	"main/data/model"
	httpclient "main/infra/http_client"
	"main/infra/logger"
)

type repository struct {
	logger     logger.Logger
	httpClient httpclient.HttpClient

	gitlabAPIBaseURL string
}

type Repository interface {
	ValidateGitlabPersonalAccessToken(
		ctx context.Context,
		token string,
	) (bool, error)
	GetGitlabRefComparation(
		query *dto.GetGitlabRepositoryRefComparationQuery,
	) (*dto.RefComparationResultDTO, error)
	GetGitlabProject(
		ctx context.Context,
		projectID int,
		token string,
	) (*model.ProjectModel, error)
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
