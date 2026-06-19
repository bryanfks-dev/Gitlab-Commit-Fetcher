package service

import (
	"context"
	"main/core/errors"
	"main/data/mapper"
	vaultconsts "main/infra/vault/consts"
)

func (s *service) RegisterAndCheckProjectExistAndValid(
	ctx context.Context,
	projectID int,
) error {
	s.logger.Infof(
		"Registering and checking if project exists and is valid with ID %d",
		projectID,
	)

	// IDK, this is not always be valid, but checking if project exists in
	// config first to avoid unnecessary API calls to GitLab
	if s.cfg.CheckProjectExists(projectID) {
		s.logger.Infof("Project with ID %d already exists in config", projectID)
		return errors.ErrProjectAlreadyRegistered
	}

	token, err := s.vaultRepo.Get(vaultconsts.GitlabPersonalTokenKey)
	if err != nil {
		s.logger.Infof("Gitlab personal access token does not exist in vault")
		return err
	}

	tokenStr := string(token)
	project, err := s.repo.GetGitlabProject(ctx, projectID, tokenStr)
	if err != nil {
		return err
	}

	if project == nil {
		s.logger.Infof("Project with ID %d does not exist", projectID)
		return errors.ErrProjectNotFound
	}

	projectCfg := mapper.MapProjectModelToProjectConfig(project)
	if err := s.cfg.RegisterProject(*projectCfg); err != nil {
		s.logger.Errorf(
			"Failed to register project with ID %d: %v",
			projectID,
			err,
		)
		return err
	}
	return nil
}
