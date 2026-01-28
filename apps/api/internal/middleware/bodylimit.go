package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lukabrkovic/artemis/pkg/apperr"
)

func RequestSizeLimiter(maxBytes int64) gin.HandlerFunc {
	return func(c *gin.Context) {
		if maxBytes <= 0 {
			c.Next()
			return
		}

		c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, maxBytes)
		c.Next()
	}
}

func MaxBytesErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			for _, err := range c.Errors {
				if err.Err != nil && err.Err.Error() == "http: request body too large" {
					c.AbortWithStatusJSON(http.StatusRequestEntityTooLarge, gin.H{
						"error": fmt.Sprintf("request body too large, maximum allowed is %d bytes", c.Request.ContentLength),
					})
					return
				}
			}
		}
	}
}

func RequireContentLength(maxBytes int64) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == "POST" || c.Request.Method == "PUT" || c.Request.Method == "PATCH" {
			if c.Request.ContentLength < 0 {
				c.Next()
				return
			}
			if c.Request.ContentLength > maxBytes {
				c.Error(apperr.New(http.StatusRequestEntityTooLarge, "request body too large"))
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
