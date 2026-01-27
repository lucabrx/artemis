package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/lukabrkovic/artemis/pkg/apperr"
	"github.com/rs/zerolog"
)

func ErrorHandler(logger zerolog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err
			var appErr *apperr.AppError
			var validationErrs validator.ValidationErrors

			if errors.As(err, &appErr) {
				c.JSON(appErr.Code, gin.H{
					"error": appErr.Message,
				})
				if appErr.Code >= 500 {
					logger.Error().Err(err).Msg("internal server error")
				}
				return
			}

			if errors.As(err, &validationErrs) {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": formatValidationError(validationErrs),
				})
				return
			}

			logger.Error().Err(err).Msg("unhandled error")
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "internal server error",
			})
		}
	}
}

func formatValidationError(errs validator.ValidationErrors) string {
	var messages []string
	for _, e := range errs {
		field := strings.ToLower(e.Field())
		switch e.Tag() {
		case "required":
			messages = append(messages, fmt.Sprintf("%s is required", field))
		case "email":
			messages = append(messages, fmt.Sprintf("%s must be a valid email", field))
		case "min":
			messages = append(messages, fmt.Sprintf("%s must be at least %s characters", field, e.Param()))
		case "url":
			messages = append(messages, fmt.Sprintf("%s must be a valid URL", field))
		default:
			messages = append(messages, fmt.Sprintf("%s is invalid", field))
		}
	}
	return strings.Join(messages, ", ")
}
