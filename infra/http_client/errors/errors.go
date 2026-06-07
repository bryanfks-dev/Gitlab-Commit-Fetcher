package httpclienterrors

import "errors"

var (
	ErrCreateRequestFailed = errors.New("failed to create HTTP request")
	ErrSendRequestFailed   = errors.New("failed to send HTTP request")
)
