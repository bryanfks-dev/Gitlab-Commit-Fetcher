package errors

import "errors"

var (
	ErrGitLabPersonalAccessTokenNotFound = errors.New(
		"GitLab personal access token not found",
	)
	ErrInvalidGitLabPersonalAccessToken = errors.New(
		"GitLab token is invalid",
	)
	ErrValidateGitLabPersonalAccessTokenFailed = errors.New(
		"validate Gitlab personal access token failed",
	)
	ErrGetGitLabRepositoryRefComparationFailed = errors.New(
		"get gitlab repository ref comparation failed",
	)
)
