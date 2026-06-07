package vaulterrors

import "errors"

var (
	ErrItemNotFound      = errors.New("item not found")
	ErrItemAlreadyExists = errors.New("item already exists")
	ErrGetItemFailed     = errors.New("failed to get item")
	ErrSetItemFailed     = errors.New("failed to set item")
	ErrDeleteItemFailed  = errors.New("failed to delete item")
)
