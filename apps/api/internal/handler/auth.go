package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lukabrkovic/artemis/internal/service"
	"github.com/lukabrkovic/artemis/internal/store"
	"github.com/lukabrkovic/artemis/internal/validator"
	"github.com/lukabrkovic/artemis/pkg/apperr"
	"github.com/lukabrkovic/artemis/pkg/token"
)

type AuthHandler struct {
	service service.Auth
}

func NewAuthHandler(service service.Auth) *AuthHandler {
	return &AuthHandler{service: service}
}

type registerRequest struct {
	Email    *string `json:"email" binding:"omitempty,email"`
	Password string  `json:"password" binding:"required,min=8"`
	Name     string  `json:"name" binding:"required,min=2"`
}

type loginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type refreshRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type logoutRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type authResponse struct {
	User                  *store.User `json:"user"`
	AccessToken           string      `json:"access_token"`
	RefreshToken          string      `json:"refresh_token"`
	AccessTokenExpiresAt  int64       `json:"access_token_expires_at"`
	RefreshTokenExpiresAt int64       `json:"refresh_token_expires_at"`
}

type tokenResponse struct {
	AccessToken           string `json:"access_token"`
	RefreshToken          string `json:"refresh_token"`
	AccessTokenExpiresAt  int64  `json:"access_token_expires_at"`
	RefreshTokenExpiresAt int64  `json:"refresh_token_expires_at"`
}

// Register godoc
// @Summary      Register new user
// @Description  Register a new user account
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request body registerRequest true "Register Request"
// @Success      201  {object}  authResponse
// @Failure      400  {object}  apperr.AppError
// @Failure      409  {object}  apperr.AppError
// @Failure      500  {object}  apperr.AppError
// @Router       /auth/register [post]
func (h *AuthHandler) Register(c *gin.Context) {
	var req registerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(apperr.BadRequest(err.Error()))
		return
	}

	serviceInput := service.RegisterInput{
		Email:    req.Email,
		Password: req.Password,
		Name:     req.Name,
	}
	if err := validator.Struct(&serviceInput); err != nil {
		c.Error(err)
		return
	}

	result, err := h.service.Register(c.Request.Context(), serviceInput, c.ClientIP(), c.GetHeader("User-Agent"))

	if err != nil {
		if errors.Is(err, service.ErrEmailExists) {
			c.Error(apperr.Conflict("email already exists"))
			return
		}
		c.Error(apperr.Internal(err))
		return
	}

	c.Set("user_id", result.User.ID)

	c.JSON(http.StatusCreated, toAuthResponse(result))
}

// Login godoc
// @Summary      Login user
// @Description  Login with email and password
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request body loginRequest true "Login Request"
// @Success      200  {object}  authResponse
// @Failure      400  {object}  apperr.AppError
// @Failure      401  {object}  apperr.AppError
// @Failure      500  {object}  apperr.AppError
// @Router       /auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var req loginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(apperr.BadRequest(err.Error()))
		return
	}

	serviceInput := service.LoginInput{
		Email:    req.Email,
		Password: req.Password,
	}
	if err := validator.Struct(&serviceInput); err != nil {
		c.Error(err)
		return
	}

	result, err := h.service.Login(c.Request.Context(), serviceInput, c.ClientIP(), c.GetHeader("User-Agent"))

	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			c.Error(apperr.Unauthorized("invalid credentials"))
			return
		}
		c.Error(apperr.Internal(err))
		return
	}

	c.Set("user_id", result.User.ID)

	c.JSON(http.StatusOK, toAuthResponse(result))
}

// Refresh godoc
// @Summary      Refresh access token
// @Description  Get a new access token using a refresh token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request body refreshRequest true "Refresh Request"
// @Success      200  {object}  tokenResponse
// @Failure      400  {object}  apperr.AppError
// @Failure      401  {object}  apperr.AppError
// @Failure      500  {object}  apperr.AppError
// @Router       /auth/refresh [post]
func (h *AuthHandler) Refresh(c *gin.Context) {
	var req refreshRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(apperr.BadRequest(err.Error()))
		return
	}

	result, err := h.service.Refresh(c.Request.Context(), req.RefreshToken, c.ClientIP(), c.GetHeader("User-Agent"))
	if err != nil {
		if errors.Is(err, token.ErrExpiredToken) || errors.Is(err, token.ErrInvalidToken) || errors.Is(err, token.ErrInvalidTokenType) {
			c.Error(apperr.Unauthorized("invalid or expired token"))
			return
		}
		c.Error(apperr.Internal(err))
		return
	}

	c.JSON(http.StatusOK, tokenResponse{
		AccessToken:           result.AccessToken,
		RefreshToken:          result.RefreshToken,
		AccessTokenExpiresAt:  result.AccessTokenExpiresAt.Unix(),
		RefreshTokenExpiresAt: result.RefreshTokenExpiresAt.Unix(),
	})
}

// Logout godoc
// @Summary      Logout user
// @Description  Revoke refresh token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request body logoutRequest true "Logout Request"
// @Success      200  {object}  map[string]string
// @Failure      400  {object}  apperr.AppError
// @Failure      500  {object}  apperr.AppError
// @Router       /auth/logout [post]
func (h *AuthHandler) Logout(c *gin.Context) {
	var req logoutRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(apperr.BadRequest(err.Error()))
		return
	}

	if err := h.service.Logout(c.Request.Context(), req.RefreshToken); err != nil {
		c.Error(apperr.Internal(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "logged out"})
}

func toAuthResponse(result *service.AuthResult) authResponse {
	return authResponse{
		User:                  result.User,
		AccessToken:           result.AccessToken,
		RefreshToken:          result.RefreshToken,
		AccessTokenExpiresAt:  result.AccessTokenExpiresAt.Unix(),
		RefreshTokenExpiresAt: result.RefreshTokenExpiresAt.Unix(),
	}
}
