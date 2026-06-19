package service

import (
	"context"
	"main/core/errors"
)

func (s *service) RemoveProject(ctx context.Context, projectID int) error {
	s.logger.Infof("Removing project with ID %d", projectID)

	if !s.cfg.CheckProjectExists(projectID) {
		s.logger.Infof("Project with ID %d not found", projectID)
		return errors.ErrProjectNotFound
	}

	if err := s.cfg.UnregisterProject(projectID); err != nil {
		s.logger.Errorf(
			"Failed to remove project with ID %d: %v",
			projectID,
			err,
		)
		return errors.ErrRemoveProjectFailed
	}
	return nil
}
