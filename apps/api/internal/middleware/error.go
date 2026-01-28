package middleware

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	v10 "github.com/go-playground/validator/v10"
	appvalidator "github.com/lukabrkovic/artemis/internal/validator"
	"github.com/lukabrkovic/artemis/pkg/apperr"
	"github.com/rs/zerolog"
)

func ErrorHandler(logger zerolog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err

			var appErr *apperr.AppError
			if errors.As(err, &appErr) {
				handleAppError(c, logger, appErr)
				return
			}

			var validationErrs appvalidator.ValidationErrors
			if errors.As(err, &validationErrs) {
				c.JSON(http.StatusBadRequest, gin.H{
					"error":   "validation failed",
					"details": validationErrs,
				})
				return
			}

			var stdValidationErrs v10.ValidationErrors
			if errors.As(err, &stdValidationErrs) {
				c.JSON(http.StatusBadRequest, gin.H{
					"error":   "validation failed",
					"details": formatStdValidationError(stdValidationErrs),
				})
				return
			}

			logger.Error().Err(err).Str("path", c.Request.URL.Path).Msg("unhandled error")
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "internal server error",
			})
		}
	}
}

func handleAppError(c *gin.Context, logger zerolog.Logger, appErr *apperr.AppError) {
	if appErr.Code >= 500 {
		logger.Error().
			Err(appErr).
			Str("path", c.Request.URL.Path).
			Int("code", appErr.Code).
			Msg("server error")
	}

	message := appErr.Message
	if appErr.Code >= 500 {
		message = "internal server error"
	}

	c.JSON(appErr.Code, gin.H{
		"error": message,
	})
}

func formatStdValidationError(errs v10.ValidationErrors) []map[string]string {
	var result []map[string]string
	for _, e := range errs {
		result = append(result, map[string]string{
			"field":   e.Field(),
			"message": getStdErrorMessage(e),
		})
	}
	return result
}

func getStdErrorMessage(e v10.FieldError) string {
	switch e.Tag() {
	case "required":
		return "this field is required"
	case "email":
		return "must be a valid email address"
	case "min":
		return "must be at least " + e.Param() + " characters"
	case "max":
		return "must be at most " + e.Param() + " characters"
	case "uuid":
		return "must be a valid UUID"
	case "url":
		return "must be a valid URL"
	default:
		return "invalid value"
	}
}
