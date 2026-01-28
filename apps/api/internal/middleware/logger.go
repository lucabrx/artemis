package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func Logger(baseLogger zerolog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		c.Next()

		param := gin.LogFormatterParams{
			Request: c.Request,
			Keys:    c.Keys,
		}

		end := time.Now()
		latency := end.Sub(start)

		if raw != "" {
			path = path + "?" + raw
		}

		logger := GetLogger(c)
		if logger.GetLevel() == zerolog.Disabled {
			logger = baseLogger
		}

		requestID, _ := c.Get(RequestIDKey)

		msg := "request"
		if len(c.Errors) > 0 {
			msg = c.Errors.String()
		}

		event := logger.Info()
		if c.Writer.Status() >= 500 {
			event = logger.Error()
		} else if c.Writer.Status() >= 400 {
			event = logger.Warn()
		}

		event.
			Int("status", c.Writer.Status()).
			Str("method", param.Method).
			Str("path", path).
			Str("ip", c.ClientIP()).
			Dur("latency", latency).
			Interface("request_id", requestID).
			Msg(msg)
	}
}
