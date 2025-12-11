package response

import (
	"fmt"
	"runtime/debug"
)

type AppError struct {
	Status  string
	Message string
	Track   any
}

func NewAppError(message string) *AppError {
	if message == "" {
		message = "ERROR"
	}

	stackTrace := string(debug.Stack())

	return &AppError{
		Status:  "Error",
		Message: message,
		Track:   stackTrace,
	}
}

func (e *AppError) Error() string {
	return fmt.Sprintf("Error %s", e.Message)
}
