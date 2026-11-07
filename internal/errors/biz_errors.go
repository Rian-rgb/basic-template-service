package errors

import "errors"

var (
	ErrAlreadyExists       = errors.New("already exists")
	ErrInternalServerError = errors.New("an unexpected errors occurred. Please try again later")
)
