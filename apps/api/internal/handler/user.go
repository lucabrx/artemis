package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lukabrkovic/artemis/internal/service"
	"github.com/lukabrkovic/artemis/internal/store"
	"github.com/lukabrkovic/artemis/pkg/apperr"
	"github.com/lukabrkovic/artemis/pkg/token"
)

type UserHandler struct {
	service service.User
}

func NewUserHandler(service service.User) *UserHandler {
	return &UserHandler{service: service}
}

type updateProfileRequest struct {
	Name      string  `json:"name" binding:"required,min=2"`
	Email     *string `json:"email" binding:"omitempty,email"`
	AvatarURL *string `json:"avatar_url" binding:"omitempty,url"`
}

// Me godoc
// @Summary      Get current user
// @Description  Get the currently logged-in user's profile
// @Tags         user
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200  {object}  store.User
// @Failure      401  {object}  apperr.AppError
// @Failure      500  {object}  apperr.AppError
// @Router       /users/me [get]
func (h *UserHandler) Me(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		c.Error(apperr.Unauthorized(err.Error()))
		return
	}

	user, err := h.service.GetUser(c.Request.Context(), userId)
	if err != nil {
		if errors.Is(err, store.ErrUserNotFound) {
			c.Error(apperr.NotFound("user not found"))
			return
		}
		c.Error(apperr.Internal(err))
		return
	}

	c.JSON(http.StatusOK, user)
}

// UpdateProfile godoc
// @Summary      Update user profile
// @Description  Update allowed profile fields
// @Tags         user
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body updateProfileRequest true "Update Profile Request"
// @Success      200  {object}  store.User
// @Failure      400  {object}  apperr.AppError
// @Failure      401  {object}  apperr.AppError
// @Failure      409  {object}  apperr.AppError
// @Failure      500  {object}  apperr.AppError
// @Router       /users/profile [patch]
func (h *UserHandler) UpdateProfile(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		c.Error(apperr.Unauthorized(err.Error()))
		return
	}

	var req updateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(apperr.BadRequest(err.Error()))
		return
	}

	user, err := h.service.UpdateProfile(c.Request.Context(), userId, service.UpdateProfileInput{
		Name:      req.Name,
		Email:     req.Email,
		AvatarURL: req.AvatarURL,
	})
	if err != nil {
		if errors.Is(err, service.ErrEmailExists) {
			c.Error(apperr.Conflict("email already exists"))
			return
		}
		c.Error(apperr.Internal(err))
		return
	}

	c.JSON(http.StatusOK, user)
}

// UploadAvatar godoc
// @Summary      Upload avatar
// @Description  Upload a new avatar image
// @Tags         user
// @Accept       multipart/form-data
// @Produce      json
// @Security     BearerAuth
// @Param        avatar formData file true "Avatar image"
// @Success      200  {object}  map[string]string
// @Failure      400  {object}  apperr.AppError
// @Failure      401  {object}  apperr.AppError
// @Failure      500  {object}  apperr.AppError
// @Router       /users/avatar [post]
func (h *UserHandler) UploadAvatar(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		c.Error(apperr.Unauthorized(err.Error()))
		return
	}

	file, header, err := c.Request.FormFile("avatar")
	if err != nil {
		c.Error(apperr.BadRequest("invalid file"))
		return
	}
	defer file.Close()

	if header.Size > 5*1024*1024 {
		c.Error(apperr.BadRequest("file size too large (max 5MB)"))
		return
	}

	contentType := header.Header.Get("Content-Type")
	if contentType != "image/jpeg" && contentType != "image/png" && contentType != "image/webp" {
		c.Error(apperr.BadRequest("invalid file type (allowed: jpeg, png, webp)"))
		return
	}

	avatarURL, err := h.service.UploadAvatar(c.Request.Context(), userId, file, header.Size, contentType)
	if err != nil {
		c.Error(apperr.Internal(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"avatar_url": avatarURL})
}

// GetSessions godoc
// @Summary      List sessions
// @Description  List all active sessions for the user with filtering, sorting, and pagination
// @Tags         user
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        limit    query     int     false  "Limit (default 20, max 100)"
// @Param        offset   query     int     false  "Offset (default 0)"
// @Param        sort_by  query     string  false  "Sort by: created_at, expires_at (default: created_at)"
// @Param        order    query     string  false  "Order: asc, desc (default: desc)"
// @Param        search   query     string  false  "Search in IP address or user agent"
// @Success      200  {object}  store.PaginatedResponse[store.Session]
// @Failure      401  {object}  apperr.AppError
// @Failure      500  {object}  apperr.AppError
// @Router       /users/sessions [get]
func (h *UserHandler) GetSessions(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		c.Error(apperr.Unauthorized(err.Error()))
		return
	}

	filters := store.DefaultFilter()
	if err := c.ShouldBindQuery(&filters); err == nil {
		filters.Normalize()
	}

	sessions, err := h.service.GetSessions(c.Request.Context(), userId, filters)
	if err != nil {
		c.Error(apperr.Internal(err))
		return
	}

	c.JSON(http.StatusOK, sessions)

}

// RevokeSession godoc
// @Summary      Revoke session
// @Description  Revoke a specific session
// @Tags         user
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      string  true  "Session ID"
// @Success      200  {object}  map[string]string
// @Failure      400  {object}  apperr.AppError
// @Failure      401  {object}  apperr.AppError
// @Failure      500  {object}  apperr.AppError
// @Router       /users/sessions/{id} [delete]
func (h *UserHandler) RevokeSession(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		c.Error(apperr.Unauthorized(err.Error()))
		return
	}

	sessionId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.Error(apperr.BadRequest("invalid session id"))
		return
	}

	if err := h.service.RevokeSession(c.Request.Context(), userId, sessionId); err != nil {
		c.Error(apperr.Internal(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "session revoked"})
}

func getUserId(c *gin.Context) (uuid.UUID, error) {
	payload, exists := c.Get("token_payload")
	if !exists {
		return uuid.Nil, errors.New("unauthorized")
	}

	tokenPayload, ok := payload.(*token.Payload)
	if !ok {
		return uuid.Nil, errors.New("invalid token")
	}
	return tokenPayload.UserID, nil
}
