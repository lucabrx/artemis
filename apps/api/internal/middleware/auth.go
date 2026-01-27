package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lukabrkovic/artemis/pkg/apperr"
	"github.com/lukabrkovic/artemis/pkg/token"
)

func Auth(tokenMaker token.Maker) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("authorization")
		if len(authHeader) == 0 {
			c.Error(apperr.Unauthorized("authorization header required"))
			c.Abort()
			return
		}

		fields := strings.Fields(authHeader)
		if len(fields) < 2 {
			c.Error(apperr.Unauthorized("invalid authorization format"))
			c.Abort()
			return
		}

		if strings.ToLower(fields[0]) != "bearer" {
			c.Error(apperr.Unauthorized("unsupported authorization type"))
			c.Abort()
			return
		}

		payload, err := tokenMaker.VerifyAccessToken(fields[1])
		if err != nil {
			c.Error(apperr.Unauthorized(err.Error()))
			c.Abort()
			return
		}

		c.Set("token_payload", payload)
		c.Next()
	}
}
