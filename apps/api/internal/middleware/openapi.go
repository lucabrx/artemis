package middleware

import (
	"bytes"
	"io"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/getkin/kin-openapi/routers"
	"github.com/getkin/kin-openapi/routers/gorillamux"
	"github.com/gin-gonic/gin"
	"github.com/lukabrkovic/artemis/pkg/apperr"
	"github.com/rs/zerolog"
)

type OpenAPIValidator struct {
	router routers.Router
	logger zerolog.Logger
}

func NewOpenAPIValidator(specPath string, logger zerolog.Logger) (*OpenAPIValidator, error) {
	loader := openapi3.NewLoader()
	doc, err := loader.LoadFromFile(specPath)
	if err != nil {
		return nil, err
	}

	if err := doc.Validate(loader.Context); err != nil {
		return nil, err
	}

	router, err := gorillamux.NewRouter(doc)
	if err != nil {
		return nil, err
	}

	return &OpenAPIValidator{
		router: router,
		logger: logger.With().Str("component", "openapi_validator").Logger(),
	}, nil
}

func (v *OpenAPIValidator) Validate() gin.HandlerFunc {
	return func(c *gin.Context) {
		if shouldSkipValidation(c.Request.URL.Path) {
			c.Next()
			return
		}

		route, pathParams, err := v.router.FindRoute(c.Request)
		if err != nil {
			v.logger.Debug().
				Str("path", c.Request.URL.Path).
				Str("method", c.Request.Method).
				Msg("route not found in OpenAPI spec")
			c.Next()
			return
		}

		var bodyBytes []byte
		if c.Request.Body != nil {
			bodyBytes, _ = io.ReadAll(c.Request.Body)
			c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		}

		requestInput := &openapi3filter.RequestValidationInput{
			Request:    c.Request,
			PathParams: pathParams,
			Route:      route,
		}

		if err := openapi3filter.ValidateRequest(c.Request.Context(), requestInput); err != nil {
			v.logger.Debug().
				Err(err).
				Str("path", c.Request.URL.Path).
				Str("method", c.Request.Method).
				Msg("OpenAPI validation failed")

			msg := formatOpenAPIError(err)
			c.Error(apperr.BadRequest(msg))
			c.Abort()
			return
		}

		if len(bodyBytes) > 0 {
			c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		}

		c.Next()
	}
}

func (v *OpenAPIValidator) ValidateResponse() gin.HandlerFunc {
	return func(c *gin.Context) {
		if shouldSkipValidation(c.Request.URL.Path) {
			c.Next()
			return
		}

		route, pathParams, err := v.router.FindRoute(c.Request)
		if err != nil {
			c.Next()
			return
		}

		blw := &bodyLogWriter{body: &bytes.Buffer{}, ResponseWriter: c.Writer}
		c.Writer = blw

		c.Next()

		responseInput := &openapi3filter.ResponseValidationInput{
			RequestValidationInput: &openapi3filter.RequestValidationInput{
				Request:    c.Request,
				PathParams: pathParams,
				Route:      route,
			},
			Status: c.Writer.Status(),
			Header: blw.Header(),
			Body:   io.NopCloser(bytes.NewBuffer(blw.body.Bytes())),
		}

		if err := openapi3filter.ValidateResponse(c.Request.Context(), responseInput); err != nil {
			v.logger.Warn().
				Err(err).
				Str("path", c.Request.URL.Path).
				Int("status", c.Writer.Status()).
				Msg("OpenAPI response validation failed")
		}
	}
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func shouldSkipValidation(path string) bool {
	skipped := []string{
		"/health",
		"/swagger",
		"/docs",
		"/api/docs",
	}

	for _, prefix := range skipped {
		if strings.HasPrefix(path, prefix) {
			return true
		}
	}
	return false
}

func formatOpenAPIError(err error) string {
	if err == nil {
		return "validation failed"
	}

	errStr := err.Error()

	replacements := map[string]string{
		"property":      "field",
		"is required":   "is required",
		"invalid type":  "has invalid type",
		"expected type": "expected",
		"value must be": "must be",
		"minimum":       "must be at least",
		"maximum":       "must be at most",
		"minLength":     "must have at least",
		"maxLength":     "must have at most",
		"pattern":       "has invalid format",
		"enum":          "must be one of",
		"format":        "has invalid format",
	}

	for old, new := range replacements {
		errStr = strings.ReplaceAll(errStr, old, new)
	}

	if len(errStr) > 200 {
		if idx := strings.Index(errStr, "Error at\""); idx != -1 {
			endIdx := idx + 100
			if endIdx > len(errStr) {
				endIdx = len(errStr)
			}
			return "validation failed: " + errStr[idx:endIdx]
		}
		return "request validation failed - check your input"
	}

	return errStr
}
