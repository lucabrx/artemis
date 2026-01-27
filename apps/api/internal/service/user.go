package service

import (
	"context"
	"errors"
	"io"

	"github.com/google/uuid"
	"github.com/lukabrkovic/artemis/internal/cache"
	"github.com/lukabrkovic/artemis/internal/store"
	pkgstorage "github.com/lukabrkovic/artemis/pkg/storage"
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
	store   *store.Store
	cache   cache.UserCache
	storage pkgstorage.Provider
}

func NewUserService(store *store.Store, cache cache.UserCache, storage pkgstorage.Provider) *UserService {
	return &UserService{
		store:   store,
		cache:   cache,
		storage: storage,
	}
}

var _ User = (*UserService)(nil)

func (s *UserService) GetUser(ctx context.Context, id uuid.UUID) (*store.User, error) {
	user, err := s.cache.GetUser(ctx, id)
	if err == nil {
		return user, nil
	}

	user, err = s.store.Users.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	_ = s.cache.SetUser(ctx, user)

	return user, nil
}

func (s *UserService) UpdateProfile(ctx context.Context, userID uuid.UUID, input UpdateProfileInput) (*store.User, error) {
	current, err := s.store.Users.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	if input.Email != nil && (current.Email == nil || *input.Email != *current.Email) {
		_, err := s.store.Users.GetUserByEmail(ctx, *input.Email)
		if err == nil {
			return nil, ErrEmailExists
		}
		if !errors.Is(err, store.ErrUserNotFound) {
			return nil, err
		}
	}

	updatedUser, err := s.store.Users.UpdateUser(ctx, store.UpdateUserParams{
		ID:           userID,
		Email:        input.Email,
		Name:         input.Name,
		AvatarURL:    input.AvatarURL,
		PasswordHash: current.PasswordHash,
	})
	if err != nil {
		return nil, err
	}

	_ = s.cache.SetUser(ctx, updatedUser)

	return updatedUser, nil
}

func (s *UserService) GetSessions(ctx context.Context, userID uuid.UUID) ([]store.Session, error) {
	return s.store.Sessions.GetSessionsByUserID(ctx, userID)
}

func (s *UserService) RevokeSession(ctx context.Context, userID, sessionID uuid.UUID) error {
	session, err := s.store.Sessions.GetSessionByID(ctx, sessionID)
	if err != nil {
		return err
	}

	if session.UserID != userID {
		return errors.New("unauthorized")
	}

	return s.store.Sessions.DeleteSession(ctx, sessionID)
}

func (s *UserService) UploadAvatar(ctx context.Context, userID uuid.UUID, reader io.Reader, size int64, contentType string) (string, error) {
	user, err := s.store.Users.GetUserByID(ctx, userID)
	if err != nil {
		return "", err
	}

	avatarURL, err := s.storage.UploadAvatar(ctx, userID.String(), reader, size, contentType)
	if err != nil {
		return "", err
	}

	_, err = s.store.Users.UpdateUser(ctx, store.UpdateUserParams{
		ID:           user.ID,
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
