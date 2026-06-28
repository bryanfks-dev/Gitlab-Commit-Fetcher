package usecase

import "main/data/model"

func (uc *useCase) GetRegisteredGitlabProjects() []model.ProjectModel {
	var projects []model.ProjectModel
	for _, registeredProject := range uc.cfg.RegisteredProjects {
		project := model.NewProjectModel(
			registeredProject.ID,
			registeredProject.Name,
			registeredProject.WebURL,
		)
		projects = append(projects, *project)
	}
	return projects
}
