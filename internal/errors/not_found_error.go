package errors

import (
	"errors"
	"fmt"
)

type notFoundError struct {
	errorMessage error
}

func (e *notFoundError) Error() string {
	return fmt.Sprintf("%v", e.errorMessage)
}

func IsNotFoundError(err error) bool {
	_, ok := err.(*notFoundError)
	return ok
}

func NotFoundError(errorMessage string) *notFoundError {
	return &notFoundError{
		errorMessage: errors.New(errorMessage),
	}
}
