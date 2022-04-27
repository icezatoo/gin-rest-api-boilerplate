package errors

import (
	"errors"
	"fmt"
)

type alredyExistsError struct {
	errorMessage error
}

func (e *alredyExistsError) Error() string {
	return fmt.Sprintf("%v", e.errorMessage)
}

func IsAlredyExistsError(err error) bool {
	_, ok := err.(*alredyExistsError)
	return ok
}

func AlredyExistsError(errorMessage string) *alredyExistsError {
	return &alredyExistsError{
		errorMessage: errors.New(errorMessage),
	}
}
