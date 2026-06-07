package dto

import "context"

type GetGitlabRepositoryCommitMessagesDTO struct {
	Ctx       context.Context
	ProjectID int
	LowerRef  string
	UpperRef  string
}
