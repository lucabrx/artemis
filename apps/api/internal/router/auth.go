package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lukabrkovic/artemis/internal/handler"
	"github.com/lukabrkovic/artemis/internal/middleware"
)

func RegisterAuthRoutes(r *gin.RouterGroup, h *handler.AuthHandler) {
	auth := r.Group("/auth")
	{
		auth.POST("/register", middleware.RateLimiterForAuth(), h.Register)
		auth.POST("/login", middleware.RateLimiterForAuth(), h.Login)
		auth.POST("/refresh", h.Refresh)
		auth.POST("/logout", h.Logout)
	}
}
