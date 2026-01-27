package service

import (
	"context"
	"errors"
	"time"

	"github.com/lukabrkovic/artemis/internal/cache"
	"github.com/lukabrkovic/artemis/internal/config"
	"github.com/lukabrkovic/artemis/internal/store"
	"github.com/lukabrkovic/artemis/pkg/token"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrEmailExists        = errors.New("email already exists")
)

type Auth interface {
	Register(ctx context.Context, input RegisterInput, ip, userAgent string) (*AuthResult, error)
	Login(ctx context.Context, input LoginInput, ip, userAgent string) (*AuthResult, error)
	Refresh(ctx context.Context, refreshToken, ip, userAgent string) (*TokenResult, error)
	Logout(ctx context.Context, refreshToken string) error
}

type AuthService struct {
	store       *store.Store
	cache       cache.UserCache
	tokenMaker  token.Maker
	tokenConfig config.TokenConfig
}

func NewAuthService(store *store.Store, cache cache.UserCache, tokenMaker token.Maker, tokenConfig config.TokenConfig) *AuthService {
	return &AuthService{
		store:       store,
		cache:       cache,
		tokenMaker:  tokenMaker,
		tokenConfig: tokenConfig,
	}
}

var _ Auth = (*AuthService)(nil)

type RegisterInput struct {
	Email    *string
	Password string
	Name     string
}

type LoginInput struct {
	Email    string
	Password string
}

type UpdateProfileInput struct {
	Name      string
	Email     *string
	AvatarURL *string
}

type AuthResult struct {
	User         *store.User
	AccessToken  string
	RefreshToken string
}

type TokenResult struct {
	AccessToken  string
	RefreshToken string
}

func (s *AuthService) Register(ctx context.Context, input RegisterInput, ip, userAgent string) (*AuthResult, error) {
	if input.Email != nil {
		_, err := s.store.Users.GetUserByEmail(ctx, *input.Email)
		if err == nil {
			return nil, ErrEmailExists
		}
		if !errors.Is(err, store.ErrUserNotFound) {
			return nil, err
		}
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	var result *AuthResult
	err = s.store.ExecTx(ctx, func(tx *store.Store) error {
		user, err := tx.Users.CreateUser(ctx, store.CreateUserParams{
			Email:        input.Email,
			PasswordHash: string(hashedPassword),
			Name:         input.Name,
		})
		if err != nil {
			return err
		}

		_ = s.cache.SetUser(ctx, user)

		result, err = s.createAuthResult(ctx, tx, user, ip, userAgent)
		return err
	})

	return result, err
}

func (s *AuthService) Login(ctx context.Context, input LoginInput, ip, userAgent string) (*AuthResult, error) {
	user, err := s.store.Users.GetUserByEmail(ctx, input.Email)
	if err != nil {
		if errors.Is(err, store.ErrUserNotFound) {
			return nil, ErrInvalidCredentials
		}
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.Password)); err != nil {
		return nil, ErrInvalidCredentials
	}

	_ = s.cache.SetUser(ctx, user)

	var result *AuthResult
	err = s.store.ExecTx(ctx, func(tx *store.Store) error {
		result, err = s.createAuthResult(ctx, tx, user, ip, userAgent)
		return err
	})

	return result, err
}

func (s *AuthService) Refresh(ctx context.Context, refreshToken, ip, userAgent string) (*TokenResult, error) {
	payload, err := s.tokenMaker.VerifyRefreshToken(refreshToken)
	if err != nil {
		return nil, err
	}

	session, err := s.store.Sessions.GetSessionByToken(ctx, refreshToken)
	if err != nil {
		return nil, err
	}

	if session.UserID != payload.UserID {
		return nil, errors.New("session mismatch")
	}

	if time.Now().After(session.ExpiresAt) {
		_ = s.store.Sessions.DeleteSession(ctx, session.ID)
		return nil, token.ErrExpiredToken
	}

	if err := s.store.Sessions.DeleteSession(ctx, session.ID); err != nil {
		return nil, err
	}

	accessToken, _, err := s.tokenMaker.CreateAccessToken(payload.UserID)
	if err != nil {
		return nil, err
	}

	newRefreshToken, refreshPayload, err := s.tokenMaker.CreateRefreshToken(payload.UserID)
	if err != nil {
		return nil, err
	}

	_, err = s.store.Sessions.CreateSession(ctx, store.CreateSessionParams{
		UserID:       payload.UserID,
		RefreshToken: newRefreshToken,
		IPAddress:    ip,
		UserAgent:    userAgent,
		ExpiresAt:    refreshPayload.ExpiredAt,
	})
	if err != nil {
		return nil, err
	}

	return &TokenResult{
		AccessToken:  accessToken,
		RefreshToken: newRefreshToken,
	}, nil
}

func (s *AuthService) Logout(ctx context.Context, refreshToken string) error {
	return s.store.Sessions.DeleteSessionByToken(ctx, refreshToken)
}

func (s *AuthService) createAuthResult(ctx context.Context, tx *store.Store, user *store.User, ip, userAgent string) (*AuthResult, error) {
	accessToken, _, err := s.tokenMaker.CreateAccessToken(user.ID)
	if err != nil {
		return nil, err
	}

	refreshToken, refreshPayload, err := s.tokenMaker.CreateRefreshToken(user.ID)
	if err != nil {
		return nil, err
	}

	_, err = tx.Sessions.CreateSession(ctx, store.CreateSessionParams{
		UserID:       user.ID,
		RefreshToken: refreshToken,
		IPAddress:    ip,
		UserAgent:    userAgent,
		ExpiresAt:    refreshPayload.ExpiredAt,
	})
	if err != nil {
		return nil, err
	}

	return &AuthResult{
		User:         user,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
