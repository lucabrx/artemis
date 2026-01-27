package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lukabrkovic/artemis/internal/handler"
)

func RegisterAuthRoutes(r *gin.RouterGroup, h *handler.AuthHandler) {
	auth := r.Group("/auth")
	{
		auth.POST("/register", h.Register)
		auth.POST("/login", h.Login)
		auth.POST("/refresh", h.Refresh)
		auth.POST("/logout", h.Logout)
	}
}
