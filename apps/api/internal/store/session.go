package store

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
)

var ErrSessionNotFound = errors.New("session not found")

type Session struct {
	ID           uuid.UUID `json:"id" db:"id"`
	UserID       uuid.UUID `json:"user_id" db:"user_id"`
	RefreshToken string    `json:"-" db:"refresh_token"`
	IPAddress    string    `json:"ip_address" db:"ip_address"`
	UserAgent    string    `json:"user_agent" db:"user_agent"`
	ExpiresAt    time.Time `json:"expires_at" db:"expires_at"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
}

type CreateSessionParams struct {
	UserID       uuid.UUID
	RefreshToken string
	IPAddress    string
	UserAgent    string
	ExpiresAt    time.Time
}

type SessionRepository interface {
	CreateSession(ctx context.Context, arg CreateSessionParams) (*Session, error)
	GetSessionByID(ctx context.Context, id any) (*Session, error)
	GetSessionByToken(ctx context.Context, refreshToken string) (*Session, error)
	GetSessionsByUserID(ctx context.Context, userID any) ([]Session, error)
	DeleteSession(ctx context.Context, id any) error
	DeleteSessionByToken(ctx context.Context, refreshToken string) error
	DeleteSessionsByUserID(ctx context.Context, userID any) error
	DeleteExpiredSessions(ctx context.Context) (int64, error)
}

type sessionRepository struct {
	db DBTX
}

func NewSessionRepository(db DBTX) SessionRepository {
	return &sessionRepository{db: db}
}

func (r *sessionRepository) CreateSession(ctx context.Context, arg CreateSessionParams) (*Session, error) {
	session := &Session{}
	query := `
		INSERT INTO sessions (user_id, refresh_token, ip_address, user_agent, expires_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, user_id, refresh_token, ip_address, user_agent, expires_at, created_at
	`
	err := r.db.GetContext(ctx, session, query, arg.UserID, arg.RefreshToken, arg.IPAddress, arg.UserAgent, arg.ExpiresAt)
	if err != nil {
		return nil, err
	}
	return session, nil
}

func (r *sessionRepository) GetSessionByID(ctx context.Context, id any) (*Session, error) {
	var session Session
	query := `SELECT * FROM sessions WHERE id = $1`
	err := r.db.GetContext(ctx, &session, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrSessionNotFound
		}
		return nil, err
	}
	return &session, nil
}

func (r *sessionRepository) GetSessionByToken(ctx context.Context, refreshToken string) (*Session, error) {
	var session Session
	query := `SELECT * FROM sessions WHERE refresh_token = $1`
	err := r.db.GetContext(ctx, &session, query, refreshToken)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrSessionNotFound
		}
		return nil, err
	}
	return &session, nil
}

func (r *sessionRepository) GetSessionsByUserID(ctx context.Context, userID any) ([]Session, error) {
	var sessions []Session
	query := `SELECT * FROM sessions WHERE user_id = $1 ORDER BY created_at DESC`
	err := r.db.SelectContext(ctx, &sessions, query, userID)
	if err != nil {
		return nil, err
	}
	return sessions, nil
}

func (r *sessionRepository) DeleteSession(ctx context.Context, id any) error {
	query := `DELETE FROM sessions WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

func (r *sessionRepository) DeleteSessionByToken(ctx context.Context, refreshToken string) error {
	query := `DELETE FROM sessions WHERE refresh_token = $1`
	_, err := r.db.ExecContext(ctx, query, refreshToken)
	return err
}

func (r *sessionRepository) DeleteSessionsByUserID(ctx context.Context, userID any) error {
	query := `DELETE FROM sessions WHERE user_id = $1`
	_, err := r.db.ExecContext(ctx, query, userID)
	return err
}

func (r *sessionRepository) DeleteExpiredSessions(ctx context.Context) (int64, error) {
	query := `DELETE FROM sessions WHERE expires_at < NOW()`
	result, err := r.db.ExecContext(ctx, query)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}
