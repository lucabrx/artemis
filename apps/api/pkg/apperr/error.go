package apperr

import (
	"fmt"
	"net/http"
)

type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Cause   error  `json:"-"`
}

func (e *AppError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Cause)
	}
	return e.Message
}

func (e *AppError) Unwrap() error {
	return e.Cause
}

func New(code int, message string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
	}
}

func Wrap(err error, code int, message string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Cause:   err,
	}
}

func BadRequest(message string) *AppError {
	return New(http.StatusBadRequest, message)
}

func Unauthorized(message string) *AppError {
	return New(http.StatusUnauthorized, message)
}

func Forbidden(message string) *AppError {
	return New(http.StatusForbidden, message)
}

func NotFound(message string) *AppError {
	return New(http.StatusNotFound, message)
}

func Conflict(message string) *AppError {
	return New(http.StatusConflict, message)
}

func Internal(err error) *AppError {
	return Wrap(err, http.StatusInternalServerError, "internal server error")
}
