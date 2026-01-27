package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lukabrkovic/artemis/internal/handler"
	"github.com/lukabrkovic/artemis/internal/middleware"
	"github.com/lukabrkovic/artemis/pkg/token"
)

func RegisterWorkspaceRoutes(r *gin.RouterGroup, h *handler.WorkspaceHandler, tokenMaker token.Maker) {
	protected := r.Group("/workspaces")
	protected.Use(middleware.Auth(tokenMaker))
	{
		protected.POST("", h.CreateWorkspace)
		protected.GET("", h.ListWorkspaces)
		protected.GET("/:id", h.GetWorkspace)
		protected.PUT("/:id", h.UpdateWorkspace)
		protected.DELETE("/:id", h.DeleteWorkspace)

		protected.POST("/:id/members", h.AddMember)
		protected.GET("/:id/members", h.ListMembers)
		protected.DELETE("/:id/members/:user_id", h.RemoveMember)
		protected.POST("/:id/avatar", h.UploadAvatar)
	}
}
