package errors

import "errors"

var (
	ErrInvalidGitlabPersonalToken = errors.New(
		"invalid Gitlab personal token",
	)
	ErrValidateGitlabPersonalTokenFailed = errors.New(
		"validate Gitlab personal token failed",
	)
)
