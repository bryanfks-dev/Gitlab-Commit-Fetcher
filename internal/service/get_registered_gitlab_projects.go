package service

import "main/data/model"

func (s *service) GetRegisteredGitlabProjects() []model.ProjectModel {
	var projects []model.ProjectModel
	for _, registeredProject := range s.cfg.RegisteredProjects {
		project := model.NewProjectModel(
			registeredProject.ID,
			registeredProject.Name,
			registeredProject.WebURL,
		)
		projects = append(projects, *project)
	}
	return projects
}
