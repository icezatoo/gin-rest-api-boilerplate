package errors

import (
	"errors"
	"fmt"
)

type authFailedError struct {
	errorMessage error
}

func (e *authFailedError) Error() string {
	return fmt.Sprintf("%v", e.errorMessage)
}

func IsAuthFailedError(err error) bool {
	_, ok := err.(*authFailedError)
	return ok
}

func AuthFailedError(errorMessage string) *authFailedError {
	return &authFailedError{
		errorMessage: errors.New(errorMessage),
	}
}
