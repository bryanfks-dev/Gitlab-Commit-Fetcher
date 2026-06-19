package mapper

import (
	"main/core/config"
	"main/data/model"
	"main/data/response"
	"main/data/viewmodel"
)

func MapGitlabProjectToProjectModel(
	gitlabProject *response.GitlabProjectResponse,
) *model.ProjectModel {
	return model.NewProjectModel(
		gitlabProject.ID,
		gitlabProject.Name,
		gitlabProject.WebURL,
	)
}

func MapProjectModelToProjectViewModel(
	projectModel *model.ProjectModel,
) *viewmodel.ProjectViewModel {
	return viewmodel.NewProjectViewModel(
		projectModel.ID,
		projectModel.Name,
		projectModel.WebURL,
	)
}

func MapProjectModelToProjectConfig(
	projectModel *model.ProjectModel,
) *config.ProjectConfig {
	return &config.ProjectConfig{
		ID:     projectModel.ID,
		Name:   projectModel.Name,
		WebURL: projectModel.WebURL,
	}
}
