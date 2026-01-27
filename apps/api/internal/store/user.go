package store

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
)

var ErrUserNotFound = errors.New("user not found")

type User struct {
	ID           uuid.UUID `json:"id" db:"id"`
	Email        *string   `json:"email" db:"email"`
	PasswordHash string    `json:"-" db:"password_hash"`
	Name         string    `json:"name" db:"name"`
	AvatarURL    *string   `json:"avatar_url" db:"avatar_url"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}

type CreateUserParams struct {
	Email        *string
	PasswordHash string
	Name         string
	AvatarURL    *string
}

type UpdateUserParams struct {
	ID           uuid.UUID
	Email        *string
	Name         string
	AvatarURL    *string
	PasswordHash string
}

type UserRepository interface {
	CreateUser(ctx context.Context, arg CreateUserParams) (*User, error)
	GetUserByID(ctx context.Context, id any) (*User, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (*User, error)
	DeleteUser(ctx context.Context, id any) error
}

type userRepository struct {
	db DBTX
}

func NewUserRepository(db DBTX) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(ctx context.Context, arg CreateUserParams) (*User, error) {
	user := &User{}
	query := `
		INSERT INTO users (email, password_hash, name, avatar_url)
		VALUES ($1, $2, $3, $4)
		RETURNING id, email, password_hash, name, avatar_url, created_at, updated_at
	`
	err := r.db.GetContext(ctx, user, query, arg.Email, arg.PasswordHash, arg.Name, arg.AvatarURL)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) GetUserByID(ctx context.Context, id any) (*User, error) {
	var user User
	query := `SELECT * FROM users WHERE id = $1`
	err := r.db.GetContext(ctx, &user, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	var user User
	query := `SELECT * FROM users WHERE email = $1`
	err := r.db.GetContext(ctx, &user, query, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) UpdateUser(ctx context.Context, arg UpdateUserParams) (*User, error) {
	var user User
	query := `
		UPDATE users 
		SET email = $1, name = $2, avatar_url = $3, password_hash = $4, updated_at = NOW()
		WHERE id = $5
		RETURNING id, email, password_hash, name, avatar_url, created_at, updated_at
	`
	err := r.db.GetContext(ctx, &user, query, arg.Email, arg.Name, arg.AvatarURL, arg.PasswordHash, arg.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) DeleteUser(ctx context.Context, id any) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
