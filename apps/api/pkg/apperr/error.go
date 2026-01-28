package apperr

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	ErrUnauthorized       = New(http.StatusUnauthorized, "unauthorized")
	ErrForbidden          = New(http.StatusForbidden, "forbidden")
	ErrInvalidToken       = New(http.StatusUnauthorized, "invalid or expired token")
	ErrInvalidCredentials = New(http.StatusUnauthorized, "invalid credentials")
	ErrNotFound           = New(http.StatusNotFound, "resource not found")
	ErrConflict           = New(http.StatusConflict, "resource already exists")
	ErrValidationFailed   = New(http.StatusBadRequest, "validation failed")
	ErrInternal           = New(http.StatusInternalServerError, "internal server error")
	ErrServiceUnavailable = New(http.StatusServiceUnavailable, "service temporarily unavailable")
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

func (e *AppError) Is(target error) bool {
	if t, ok := target.(*AppError); ok {
		return e.Code == t.Code && e.Message == t.Message
	}
	return errors.Is(e.Cause, target)
}

func (e *AppError) WithMessage(message string) *AppError {
	return &AppError{
		Code:    e.Code,
		Message: message,
		Cause:   e.Cause,
	}
}

func (e *AppError) WithCause(err error) *AppError {
	return &AppError{
		Code:    e.Code,
		Message: e.Message,
		Cause:   err,
	}
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

func Validation(message string) *AppError {
	return New(http.StatusBadRequest, fmt.Sprintf("validation failed: %s", message))
}

func Unauthorized(message string) *AppError {
	return New(http.StatusUnauthorized, message)
}

func Forbidden(message string) *AppError {
	return New(http.StatusForbidden, message)
}

func NotFound(resource string) *AppError {
	return New(http.StatusNotFound, fmt.Sprintf("%s not found", resource))
}

func Conflict(message string) *AppError {
	return New(http.StatusConflict, message)
}

func Internal(err error) *AppError {
	return Wrap(err, http.StatusInternalServerError, "internal server error")
}

func IsNotFound(err error) bool {
	var appErr *AppError
	if errors.As(err, &appErr) {
		return appErr.Code == http.StatusNotFound
	}
	return false
}

func IsConflict(err error) bool {
	var appErr *AppError
	if errors.As(err, &appErr) {
		return appErr.Code == http.StatusConflict
	}
	return false
}

func IsValidation(err error) bool {
	var appErr *AppError
	if errors.As(err, &appErr) {
		return appErr.Code == http.StatusBadRequest
	}
	return false
}
