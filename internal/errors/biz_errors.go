package errors

import "errors"

var (
	ErrNotFound            = errors.New("not found")
	ErrAlreadyExists       = errors.New("already exists")
	ErrInternalServerError = errors.New("an unexpected errors occurred. Please try again later")
)
