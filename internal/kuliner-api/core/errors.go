package core

import "errors"

var (
	ErrMissingName        = errors.New("ERR_MISSING_NAME")
	ErrMissingDescription = errors.New("ERR_MISSING_DESCRIPTION")
	ErrMissingID          = errors.New("ERR_MISSING_ID")
	ErrNotFound           = errors.New("ERR_NOT_FOUND")
)
