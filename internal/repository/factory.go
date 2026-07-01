package repository

import (
	"main/data/dto"
	"main/infra/logger"
)

type repository struct {
	logger logger.Logger
}

type Repository interface {
	WriteCommitMessageToFile(cmd dto.WriteCommitMessageToFileCommand) error
}

func NewRepository(logger logger.Logger) Repository {
	return &repository{
		logger: logger,
	}
}
