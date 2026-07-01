package mapper

import (
	"main/data/model"
	"main/data/response"
)

func MapGitlabCommitToCommitModel(
	gitlabCommit *response.GitlabCommitResponse,
) *model.CommitModel {
	return model.NewCommitModel(
		gitlabCommit.ID,
		gitlabCommit.Message,
		gitlabCommit.AuthorName,
		gitlabCommit.ParentIds,
	)
}
