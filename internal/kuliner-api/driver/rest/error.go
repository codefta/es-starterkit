package rest

import (
	"errors"
	"fmt"
	"net/http"
)

type apiError struct {
	StatusCode int
	Err        string
	Message    string
}

func (e *apiError) Error() string {
	return fmt.Sprintf("%v - %v - %v", e.StatusCode, e.Err, e.Message)
}

func (e *apiError) Is(target error) bool {
	var restErr *apiError
	if !errors.As(target, &restErr) {
		return false
	}
	return *e == *restErr
}

func newInternalServerError(msg string) *apiError {
	return &apiError{
		StatusCode: http.StatusInternalServerError,
		Err:        "ERR_INTERNAL_ERROR",
		Message:    msg,
	}
}

func newBadRequestError(msg string) *apiError {
	return &apiError{
		StatusCode: http.StatusBadRequest,
		Err:        "ERR_BAD_REQUEST",
		Message:    msg,
	}
}

func newNotFoundError() *apiError {
	return &apiError{
		StatusCode: http.StatusNotFound,
		Err:        "ERR_NOT_FOUND",
		Message:    "data is not found",
	}
}
