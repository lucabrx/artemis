package validator

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

func init() {
	Validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 1)[0]
		if name == "-" {
			return ""
		}
		return name
	})
}

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
	Value   any    `json:"value,omitempty"`
}

type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
	var messages []string
	for _, e := range v {
		messages = append(messages, fmt.Sprintf("%s: %s", e.Field, e.Message))
	}
	return strings.Join(messages, "; ")
}

func Struct(s any) error {
	if err := Validate.Struct(s); err != nil {
		var validationErrs validator.ValidationErrors
		if errors.As(err, &validationErrs) {
			return formatValidationErrors(validationErrs)
		}
		return err
	}
	return nil
}

func Var(field any, tag string) error {
	return Validate.Var(field, tag)
}

func formatValidationErrors(errs validator.ValidationErrors) ValidationErrors {
	var result ValidationErrors
	for _, e := range errs {
		result = append(result, ValidationError{
			Field:   e.Field(),
			Message: getErrorMessage(e),
			Value:   e.Value(),
		})
	}
	return result
}

func getErrorMessage(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return "this field is required"
	case "email":
		return "must be a valid email address"
	case "min":
		if e.Kind() == reflect.String {
			return fmt.Sprintf("must be at least %s characters", e.Param())
		}
		return fmt.Sprintf("must be at least %s", e.Param())
	case "max":
		if e.Kind() == reflect.String {
			return fmt.Sprintf("must be at most %s characters", e.Param())
		}
		return fmt.Sprintf("must be at most %s", e.Param())
	case "len":
		return fmt.Sprintf("must be exactly %s characters", e.Param())
	case "uuid":
		return "must be a valid UUID"
	case "url":
		return "must be a valid URL"
	case "oneof":
		return fmt.Sprintf("must be one of: %s", strings.ReplaceAll(e.Param(), " ", ", "))
	case "alphanum":
		return "must contain only letters and numbers"
	case "gte":
		return fmt.Sprintf("must be greater than or equal to %s", e.Param())
	case "lte":
		return fmt.Sprintf("must be less than or equal to %s", e.Param())
	case "gt":
		return fmt.Sprintf("must be greater than %s", e.Param())
	case "lt":
		return fmt.Sprintf("must be less than %s", e.Param())
	case "numeric":
		return "must be numeric"
	case "ascii":
		return "must contain only ASCII characters"
	case "printascii":
		return "must contain only printable ASCII characters"
	default:
		return fmt.Sprintf("failed validation: %s", e.Tag())
	}
}
