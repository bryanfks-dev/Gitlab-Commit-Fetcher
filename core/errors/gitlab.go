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
		"Validate GitLab personal access token failed",
	)
	ErrInvalidGitLabRefComparation = errors.New(
		"GitLab ref comparation is invalid",
	)
	ErrGetGitLabRepositoryRefComparationFailed = errors.New(
		"Get GitLab repository ref comparation failed",
	)
)
