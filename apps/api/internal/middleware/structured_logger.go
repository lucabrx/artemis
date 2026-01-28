package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type StructuredLogger struct {
	logger zerolog.Logger
}

func NewStructuredLogger(logger zerolog.Logger) *StructuredLogger {
	return &StructuredLogger{logger: logger}
}

func (l *StructuredLogger) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		if raw != "" {
			path = path + "?" + raw
		}

		c.Next()

		duration := time.Since(start)
		requestID, _ := c.Get(RequestIDKey)

		event := l.logger.Info()
		if c.Writer.Status() >= 500 {
			event = l.logger.Error()
		} else if c.Writer.Status() >= 400 {
			event = l.logger.Warn()
		}

		event.
			Str("timestamp", start.Format(time.RFC3339Nano)).
			Str("request_id", requestID.(string)).
			Str("method", c.Request.Method).
			Str("path", path).
			Int("status", c.Writer.Status()).
			Dur("duration_ms", duration).
			Str("client_ip", c.ClientIP()).
			Str("user_agent", c.Request.UserAgent()).
			Int("response_size", c.Writer.Size()).
			Str("error", c.Errors.String()).
			Msg("http_request")
	}
}
