package mapper

import (
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
