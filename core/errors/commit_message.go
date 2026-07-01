package errors

import "errors"

var (
	ErrWriteCommitMessageFailed      = errors.New("Write commit message failed")
	ErrCreateCommitMessageFileFailed = errors.New(
		"Create commit message file failed",
	)
)
