package usecase

import (
	"context"
	"main/core/errors"
)

func (uc *useCase) RemoveProject(ctx context.Context, projectID int) error {
	uc.logger.Infof("Removing project with ID %d", projectID)

	if !uc.cfg.CheckProjectExists(projectID) {
		uc.logger.Infof("Project with ID %d not found", projectID)
		return errors.ErrProjectNotFound
	}

	if err := uc.cfg.UnregisterProject(projectID); err != nil {
		uc.logger.Errorf(
			"Failed to remove project with ID %d: %v",
			projectID,
			err,
		)
		return errors.ErrRemoveProjectFailed
	}
	return nil
}
