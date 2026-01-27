package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lukabrkovic/artemis/internal/service"
	"github.com/lukabrkovic/artemis/pkg/apperr"
)

type WorkspaceHandler struct {
	service service.Workspace
}

func NewWorkspaceHandler(service service.Workspace) *WorkspaceHandler {
	return &WorkspaceHandler{service: service}
}

type createWorkspaceRequest struct {
	Name      string  `json:"name" binding:"required,min=2"`
	AvatarURL *string `json:"avatar_url" binding:"omitempty,url"`
}

type addMemberRequest struct {
	UserID uuid.UUID `json:"user_id" binding:"required"`
	Role   string    `json:"role" binding:"required,oneof=admin member"`
}

// CreateWorkspace godoc
// @Summary      Create workspace
// @Description  Create a new workspace
// @Tags         workspace
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body createWorkspaceRequest true "Create Workspace Request"
// @Success      201  {object}  store.Workspace
// @Failure      400  {object}  apperr.AppError
// @Failure      401  {object}  apperr.AppError
// @Failure      500  {object}  apperr.AppError
// @Router       /workspaces [post]
func (h *WorkspaceHandler) CreateWorkspace(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		c.Error(apperr.Unauthorized(err.Error()))
		return
	}

	var req createWorkspaceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(apperr.BadRequest(err.Error()))
		return
	}

	workspace, err := h.service.CreateWorkspace(c.Request.Context(), userId, service.CreateWorkspaceInput{
		Name:      req.Name,
		AvatarURL: req.AvatarURL,
	})
	if err != nil {
		c.Error(apperr.Internal(err))
		return
	}

	c.JSON(http.StatusCreated, workspace)
}

// ListWorkspaces godoc
// @Summary      List workspaces
// @Description  List all workspaces user is a member of
// @Tags         workspace
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200  {array}   store.WorkspaceWithRole
// @Failure      401  {object}  apperr.AppError
// @Failure      500  {object}  apperr.AppError
// @Router       /workspaces [get]
func (h *WorkspaceHandler) ListWorkspaces(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		c.Error(apperr.Unauthorized(err.Error()))
		return
	}

	workspaces, err := h.service.GetMyWorkspaces(c.Request.Context(), userId)
	if err != nil {
		c.Error(apperr.Internal(err))
		return
	}

	c.JSON(http.StatusOK, workspaces)
}

// GetWorkspace godoc
// @Summary      Get workspace
// @Description  Get a specific workspace by ID
// @Tags         workspace
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      string  true  "Workspace ID"
// @Success      200  {object}  store.Workspace
// @Failure      401  {object}  apperr.AppError
// @Failure      403  {object}  apperr.AppError
// @Failure      404  {object}  apperr.AppError
// @Failure      500  {object}  apperr.AppError
// @Router       /workspaces/{id} [get]
func (h *WorkspaceHandler) GetWorkspace(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		c.Error(apperr.Unauthorized(err.Error()))
		return
	}

	workspaceId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.Error(apperr.BadRequest("invalid workspace id"))
		return
	}

	workspace, err := h.service.GetWorkspace(c.Request.Context(), userId, workspaceId)
	if err != nil {
		if errors.Is(err, service.ErrWorkspaceNotFound) {
			c.Error(apperr.NotFound("workspace not found"))
			return
		}
		if errors.Is(err, service.ErrForbidden) {
			c.Error(apperr.Forbidden("access denied"))
			return
		}
		c.Error(apperr.Internal(err))
		return
	}

	c.JSON(http.StatusOK, workspace)
}

// UpdateWorkspace godoc
// @Summary      Update workspace
// @Description  Update workspace details
// @Tags         workspace
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id      path      string                  true  "Workspace ID"
// @Param        request body      createWorkspaceRequest  true  "Update Workspace Request"
// @Success      200     {object}  store.Workspace
// @Failure      400     {object}  apperr.AppError
// @Failure      401     {object}  apperr.AppError
// @Failure      403     {object}  apperr.AppError
// @Failure      404     {object}  apperr.AppError
// @Failure      500     {object}  apperr.AppError
// @Router       /workspaces/{id} [put]
func (h *WorkspaceHandler) UpdateWorkspace(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		c.Error(apperr.Unauthorized(err.Error()))
		return
	}

	workspaceId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.Error(apperr.BadRequest("invalid workspace id"))
		return
	}

	var req createWorkspaceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(apperr.BadRequest(err.Error()))
		return
	}

	workspace, err := h.service.UpdateWorkspace(c.Request.Context(), userId, workspaceId, req.Name)
	if err != nil {
		if errors.Is(err, service.ErrWorkspaceNotFound) {
			c.Error(apperr.NotFound("workspace not found"))
			return
		}
		if errors.Is(err, service.ErrForbidden) {
			c.Error(apperr.Forbidden("access denied"))
			return
		}
		c.Error(apperr.Internal(err))
		return
	}

	c.JSON(http.StatusOK, workspace)
}

// DeleteWorkspace godoc
// @Summary      Delete workspace
// @Description  Delete a workspace
// @Tags         workspace
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      string  true  "Workspace ID"
// @Success      200  {object}  map[string]string
// @Failure      401  {object}  apperr.AppError
// @Failure      403  {object}  apperr.AppError
// @Failure      404  {object}  apperr.AppError
// @Failure      500  {object}  apperr.AppError
// @Router       /workspaces/{id} [delete]
func (h *WorkspaceHandler) DeleteWorkspace(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		c.Error(apperr.Unauthorized(err.Error()))
		return
	}

	workspaceId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.Error(apperr.BadRequest("invalid workspace id"))
		return
	}

	if err := h.service.DeleteWorkspace(c.Request.Context(), userId, workspaceId); err != nil {
		if errors.Is(err, service.ErrForbidden) {
			c.Error(apperr.Forbidden("access denied"))
			return
		}
		c.Error(apperr.Internal(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "workspace deleted"})
}

// AddMember godoc
// @Summary      Add member
// @Description  Add a user to the workspace
// @Tags         workspace
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id       path      string            true  "Workspace ID"
// @Param        request  body      addMemberRequest  true  "Add Member Request"
// @Success      200      {object}  map[string]string
// @Failure      400      {object}  apperr.AppError
// @Failure      401      {object}  apperr.AppError
// @Failure      403      {object}  apperr.AppError
// @Failure      500      {object}  apperr.AppError
// @Router       /workspaces/{id}/members [post]
func (h *WorkspaceHandler) AddMember(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		c.Error(apperr.Unauthorized(err.Error()))
		return
	}

	workspaceId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.Error(apperr.BadRequest("invalid workspace id"))
		return
	}

	var req addMemberRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(apperr.BadRequest(err.Error()))
		return
	}

	if err := h.service.AddMember(c.Request.Context(), userId, workspaceId, req.UserID, req.Role); err != nil {
		if errors.Is(err, service.ErrForbidden) {
			c.Error(apperr.Forbidden("access denied"))
			return
		}
		if errors.Is(err, service.ErrInvalidRole) {
			c.Error(apperr.BadRequest("invalid role"))
			return
		}
		// Assuming generic db error for duplicate key
		c.Error(apperr.Internal(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "member added"})
}

// RemoveMember godoc
// @Summary      Remove member
// @Description  Remove a user from the workspace or leave
// @Tags         workspace
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id       path      string  true  "Workspace ID"
// @Param        user_id  path      string  true  "User ID to remove"
// @Success      200      {object}  map[string]string
// @Failure      400      {object}  apperr.AppError
// @Failure      401      {object}  apperr.AppError
// @Failure      403      {object}  apperr.AppError
// @Failure      500      {object}  apperr.AppError
// @Router       /workspaces/{id}/members/{user_id} [delete]
func (h *WorkspaceHandler) RemoveMember(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		c.Error(apperr.Unauthorized(err.Error()))
		return
	}

	workspaceId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.Error(apperr.BadRequest("invalid workspace id"))
		return
	}

	targetUserId, err := uuid.Parse(c.Param("user_id"))
	if err != nil {
		c.Error(apperr.BadRequest("invalid user id"))
		return
	}

	if err := h.service.RemoveMember(c.Request.Context(), userId, workspaceId, targetUserId); err != nil {
		if errors.Is(err, service.ErrForbidden) {
			c.Error(apperr.Forbidden("access denied"))
			return
		}
		c.Error(apperr.Internal(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "member removed"})
}

// ListMembers godoc
// @Summary      List members
// @Description  List all members of the workspace
// @Tags         workspace
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      string  true  "Workspace ID"
// @Success      200  {array}   store.WorkspaceMember
// @Failure      401  {object}  apperr.AppError
// @Failure      403  {object}  apperr.AppError
// @Failure      500  {object}  apperr.AppError
// @Router       /workspaces/{id}/members [get]
func (h *WorkspaceHandler) ListMembers(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		c.Error(apperr.Unauthorized(err.Error()))
		return
	}

	workspaceId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.Error(apperr.BadRequest("invalid workspace id"))
		return
	}

	members, err := h.service.GetMembers(c.Request.Context(), userId, workspaceId)
	if err != nil {
		if errors.Is(err, service.ErrForbidden) {
			c.Error(apperr.Forbidden("access denied"))
			return
		}
		c.Error(apperr.Internal(err))
		return
	}

	c.JSON(http.StatusOK, members)
}
