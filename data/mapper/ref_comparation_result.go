package mapper

import (
	"main/data/dto"
	"main/data/model"
	"main/data/response"
)

func MapGitlabRefComparationToRefComparationResultDTO(
	gitlabComparation *response.GitlabRefComparationResponse,
) *dto.RefComparationResultDTO {
	commits := []model.CommitModel{}
	for _, c := range gitlabComparation.Commits {
		commit := MapGitlabCommitToCommitModel(&c)
		commits = append(commits, *commit)
	}
	return &dto.RefComparationResultDTO{
		ComparingSameRef: gitlabComparation.CompareSameRef,
		Commits:          commits,
	}
}
