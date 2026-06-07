package errors

import "errors"

var (
	ErrProjectNotFound  = errors.New("project not found")
	ErrGetProjectFailed = errors.New("get project failed")
)
