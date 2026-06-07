package mapper

import (
	"main/data/model"
	"main/data/response"
)

func MapGitlabProjectToProjectModel(
	gitlabProject *response.GitlabProjectResponse,
) *model.ProjectModel {
	return &model.ProjectModel{
		ID:   gitlabProject.ID,
		Name: gitlabProject.Name,
	}
}
