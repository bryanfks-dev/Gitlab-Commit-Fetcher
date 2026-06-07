package mapper

import (
	"main/data/model"
	"main/data/response"
)

func MapGitlabCommitToCommitModel(
	gitlabCommit *response.GitlabCommitResponse,
) *model.CommitModel {
	return &model.CommitModel{
		ID:      gitlabCommit.ID,
		Message: gitlabCommit.Message,
	}
}
