package usecase

import (
	"main/core/errors"
	"main/data/mapper"
	"main/data/model"
)

func (uc *useCase) GetRegisteredGitlabProject(
	projectID int,
) (*model.ProjectModel, error) {
	uc.logger.Infof("Getting registered GitLab project")

	for _, registeredProject := range uc.cfg.RegisteredProjects {
		if registeredProject.ID != projectID {
			continue
		}
		project := mapper.MapProjectConfigToProjectModel(&registeredProject)
		return project, nil
	}
	return nil, errors.ErrProjectNotFound
}
