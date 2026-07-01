package usecase

import (
	"main/data/mapper"
	"main/data/model"
)

func (uc *useCase) GetRegisteredGitlabProjects() []model.ProjectModel {
	uc.logger.Infof("Getting registered Gitlab projects")

	var projects []model.ProjectModel
	for _, registeredProject := range uc.cfg.RegisteredProjects {
		project := mapper.MapProjectConfigToProjectModel(&registeredProject)
		projects = append(projects, *project)
	}
	return projects
}
