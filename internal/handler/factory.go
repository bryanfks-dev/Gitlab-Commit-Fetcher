package handler

import (
	"context"
	"main/infra/logger"
	"main/internal/service"
)

type handler struct {
	logger logger.Logger
	svc    service.Service

	gitlabBaseURL string
}

type Handler interface {
	Run(ctx context.Context) error
}

func NewHandler(
	logger logger.Logger,
	svc service.Service,
	gitlabBaseURL string,
) Handler {
	return &handler{
		logger:        logger,
		svc:           svc,
		gitlabBaseURL: gitlabBaseURL,
	}
}
