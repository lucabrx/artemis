package service

import (
	"context"
	"errors"
	"io"

	"github.com/google/uuid"
	"github.com/lukabrkovic/artemis/internal/cache"
	"github.com/lukabrkovic/artemis/internal/store"
)

type User interface {
	GetUser(ctx context.Context, id uuid.UUID) (*store.User, error)
	UpdateProfile(ctx context.Context, userID uuid.UUID, input UpdateProfileInput) (*store.User, error)
	GetSessions(ctx context.Context, userID uuid.UUID) ([]store.Session, error)
	RevokeSession(ctx context.Context, userID, sessionID uuid.UUID) error
	UploadAvatar(ctx context.Context, userID uuid.UUID, reader io.Reader, size int64, contentType string) (string, error)
}

type FileStorage interface {
	UploadAvatar(ctx context.Context, userID string, reader io.Reader, size int64, contentType string) (string, error)
}

type UserService struct {
	store   store.Repository
	cache   cache.UserCache
	storage FileStorage
}

func NewUserService(store store.Repository, cache cache.UserCache, storage FileStorage) *UserService {
	return &UserService{
		store:   store,
		cache:   cache,
		storage: storage,
	}
}

var _ User = (*UserService)(nil)

func (s *UserService) GetUser(ctx context.Context, id uuid.UUID) (*store.User, error) {
	cached, err := s.cache.GetUser(ctx, id)
	if err == nil && cached != nil {
		return cached, nil
	}

	user, err := s.store.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	_ = s.cache.SetUser(ctx, user)

	return user, nil
}

func (s *UserService) UpdateProfile(ctx context.Context, userID uuid.UUID, input UpdateProfileInput) (*store.User, error) {
	user, err := s.store.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	if input.Email != nil && (user.Email == nil || *input.Email != *user.Email) {
		_, err := s.store.GetUserByEmail(ctx, *input.Email)
		if err == nil {
			return nil, ErrEmailExists
		}
		if !errors.Is(err, store.ErrUserNotFound) {
			return nil, err
		}
	}

	updatedUser, err := s.store.UpdateUser(ctx, store.UpdateUserParams{
		ID:           userID,
		Email:        input.Email,
		Name:         input.Name,
		AvatarURL:    input.AvatarURL,
		PasswordHash: user.PasswordHash,
	})
	if err != nil {
		return nil, err
	}

	_ = s.cache.SetUser(ctx, updatedUser)

	return updatedUser, nil
}

func (s *UserService) GetSessions(ctx context.Context, userID uuid.UUID) ([]store.Session, error) {
	return s.store.GetSessionsByUserID(ctx, userID)
}

func (s *UserService) RevokeSession(ctx context.Context, userID, sessionID uuid.UUID) error {
	session, err := s.store.GetSessionByID(ctx, sessionID)
	if err != nil {
		return err
	}

	if session.UserID != userID {
		return errors.New("unauthorized")
	}

	return s.store.DeleteSession(ctx, sessionID)
}

func (s *UserService) UploadAvatar(ctx context.Context, userID uuid.UUID, reader io.Reader, size int64, contentType string) (string, error) {
	user, err := s.store.GetUserByID(ctx, userID)
	if err != nil {
		return "", err
	}

	avatarURL, err := s.storage.UploadAvatar(ctx, userID.String(), reader, size, contentType)
	if err != nil {
		return "", err
	}

	_, err = s.store.UpdateUser(ctx, store.UpdateUserParams{
		ID:           userID,
		Email:        user.Email,
		Name:         user.Name,
		AvatarURL:    &avatarURL,
		PasswordHash: user.PasswordHash,
	})
	if err != nil {
		return "", err
	}

	user.AvatarURL = &avatarURL
	_ = s.cache.SetUser(ctx, user)

	return avatarURL, nil
}
