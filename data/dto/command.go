package dto

import "context"

type GetGitlabRepositoryRefComparationQuery struct {
	Ctx       context.Context
	ProjectID int
	LowerRef  string
	UpperRef  string
	Token     string
}
