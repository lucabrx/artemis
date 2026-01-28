package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

const RequestIDKey = "request_id"
const LoggerKey = "logger"

func RequestID(logger zerolog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		rid := c.GetHeader("X-Request-ID")
		if rid == "" {
			rid = uuid.New().String()
		}

		requestLogger := logger.With().Str("request_id", rid).Logger()

		c.Set(RequestIDKey, rid)
		c.Set(LoggerKey, requestLogger)
		c.Header("X-Request-ID", rid)
		c.Next()
	}
}

func GetLogger(c *gin.Context) zerolog.Logger {
	logger, exists := c.Get(LoggerKey)
	if !exists {
		return zerolog.Logger{}
	}
	return logger.(zerolog.Logger)
}
