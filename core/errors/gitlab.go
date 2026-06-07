package errors

import "errors"

var (
	ErrInvalidGitlabPersonalAccessToken = errors.New(
		"invalid Gitlab personal access token",
	)
	ErrValidateGitlabPersonalAccessTokenFailed = errors.New(
		"validate Gitlab personal access token failed",
	)
	ErrGetGitlabRepositoryRefComparationFailed = errors.New(
		"get gitlab repository ref comparation failed",
	)
)
