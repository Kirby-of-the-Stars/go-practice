package errs

import "net/http"

type AppError struct {
	Code    int `json:",omitempty"`
	Message string
}

func (err AppError) Error() string {
	return err.Message
}

func (err AppError) AsText() AppError {
	return AppError{
		Message: err.Message,
	}
}

func NewStatusNotFoundError(message string) AppError {
	return AppError{
		http.StatusNotFound,
		message,
	}
}

func NewStatusInternalServerError(message string) AppError {
	return AppError{
		http.StatusInternalServerError,
		message,
	}
}
