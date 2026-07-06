package dto

import (
	"context"
	"main/data/model"
)

type GetGitlabRepositoryCommitMessagesDTO struct {
	Ctx       context.Context
	ProjectID int
	LowerRef  string
	UpperRef  string
}

type RefComparationResultDTO struct {
	ComparingSameRef bool
	Commits          []model.CommitModel
}
