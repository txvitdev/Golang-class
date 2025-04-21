package exceptions

import "net/http"

type HttpError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (err HttpError) Error() string {
	return err.Message
}

func NewNotFound(message string) HttpError {
	if message == "" {
		message = "Entity not found"
	}
	return HttpError{
		Code: http.StatusNotFound,
		Message: message,
	}
}

func NewInternal(message string) HttpError {
	if message == "" {
		message = "Internal server error"
	}

	return HttpError{
		Code: http.StatusInternalServerError,
		Message: message,
	}
}

func NewConflict(message string) HttpError {
	if message == "" {
		message = "Conflict"
	}

	return HttpError{
		Code: http.StatusInternalServerError,
		Message: message,
	}
}

func NewUnProcessableEntity(message string) HttpError {
	if message == "" {
		message = "UnprocessableEntity"
	}

	return HttpError{
		Code: http.StatusUnprocessableEntity,
		Message: message,
	}
}

func NewUnauthorized(message string) HttpError {
	if message == "" {
		message = "Unauthorized"
	}

	return HttpError{
		Code: http.StatusUnauthorized,
		Message: message,
	}
}