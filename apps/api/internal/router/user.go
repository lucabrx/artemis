package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lukabrkovic/artemis/internal/handler"
	"github.com/lukabrkovic/artemis/internal/middleware"
	"github.com/lukabrkovic/artemis/pkg/token"
)

func RegisterUserRoutes(r *gin.RouterGroup, h *handler.UserHandler, tokenMaker token.Maker) {
	protected := r.Group("")
	protected.Use(middleware.Auth(tokenMaker))
	{
		protected.GET("/me", h.Me)
		protected.PATCH("/me", h.UpdateProfile)
		protected.POST("/me/avatar", h.UploadAvatar)
		protected.GET("/me/sessions", h.GetSessions)
		protected.DELETE("/me/sessions/:id", h.RevokeSession)
	}
}
