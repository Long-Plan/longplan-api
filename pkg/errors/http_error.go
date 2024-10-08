package errors

import "net/http"

type Error struct {
	StatusCode int
	Message    string
}

func (e Error) Error() string {
	return e.Message
}

func NewError(statusCode int, message string) Error {
	return Error{
		StatusCode: statusCode,
		Message:    message,
	}
}

func NewInternalError(message string) Error {
	return NewError(http.StatusInternalServerError, message)
}

func NewNotFoundError(message string) Error {
	return NewError(http.StatusNotFound, message)
}

func NewBadRequestError(message string) Error {
	return NewError(http.StatusBadRequest, message)
}

func NewUnprocessableEntityError(message string) Error {
	return NewError(http.StatusUnprocessableEntity, message)
}

func NewUnauthorizedError(message string) Error {
	return NewError(http.StatusUnauthorized, message)
}

func NewForbiddenError(message string) Error {
	return NewError(http.StatusForbidden, message)
}

func NewConflictError(message string) Error {
	return NewError(http.StatusConflict, message)
}

func NewStatusBadGatewayError(message string) Error {
	return NewError(http.StatusBadGateway, message)
}

func NewGatewayTimeoutError(message string) Error {
	return NewError(http.StatusGatewayTimeout, message)
}

func NewInsufficientStorageError(message string) Error {
	return NewError(http.StatusInsufficientStorage, message)
}

func NewRequestTimeoutError(message string) Error {
	return NewError(http.StatusRequestTimeout, message)
}

func NewUnsupportedMediaTypeError(message string) Error {
	return NewError(http.StatusUnsupportedMediaType, message)
}
