package errors

import "errors"

var (
	ErrProjectNotFound          = errors.New("Project not found")
	ErrProjectAlreadyRegistered = errors.New("Project already registered")
	ErrGetProjectFailed         = errors.New("Get project failed")
	ErrCreateProjectFailed      = errors.New("Create project failed")
	ErrRemoveProjectFailed      = errors.New("Remove project failed")
)
