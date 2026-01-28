package store

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
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
	GetSessionByID(ctx context.Context, id uuid.UUID) (*Session, error)
	GetSessionByToken(ctx context.Context, refreshToken string) (*Session, error)
	GetSessionsByUserID(ctx context.Context, userID uuid.UUID, filters FilterParams) ([]Session, int64, error)
	DeleteSession(ctx context.Context, id uuid.UUID) error
	DeleteSessionByToken(ctx context.Context, refreshToken string) error
	DeleteSessionsByUserID(ctx context.Context, userID uuid.UUID) error
	DeleteExpiredSessions(ctx context.Context) (int64, error)
	CountSessionsByUserID(ctx context.Context, userID uuid.UUID) (int64, error)
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

func (r *sessionRepository) GetSessionByID(ctx context.Context, id uuid.UUID) (*Session, error) {
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

func (r *sessionRepository) GetSessionsByUserID(ctx context.Context, userID uuid.UUID, filters FilterParams) ([]Session, int64, error) {
	var sessions []Session
	var args []any
	argPos := 1

	baseQuery := `SELECT * FROM sessions WHERE user_id = $` + fmt.Sprintf("%d", argPos)
	args = append(args, userID)
	argPos++

	if filters.HasSearch() {
		baseQuery += fmt.Sprintf(` AND (ip_address ILIKE $%d OR user_agent ILIKE $%d)`, argPos, argPos)
		args = append(args, filters.GetSearchPattern())
		argPos++
	}

	baseQuery += " " + filters.GetSortClause("created_at")

	baseQuery += fmt.Sprintf(` LIMIT $%d OFFSET $%d`, argPos, argPos+1)
	args = append(args, filters.Limit, filters.Offset)

	err := r.db.SelectContext(ctx, &sessions, baseQuery, args...)
	if err != nil {
		return nil, 0, err
	}

	var totalArgs []any
	totalArgPos := 1
	countQuery := `SELECT COUNT(*) FROM sessions WHERE user_id = $` + fmt.Sprintf("%d", totalArgPos)
	totalArgs = append(totalArgs, userID)
	totalArgPos++

	if filters.HasSearch() {
		countQuery += fmt.Sprintf(` AND (ip_address ILIKE $%d OR user_agent ILIKE $%d)`, totalArgPos, totalArgPos)
		totalArgs = append(totalArgs, filters.GetSearchPattern())
	}

	var total int64
	err = r.db.GetContext(ctx, &total, countQuery, totalArgs...)
	if err != nil {
		return nil, 0, err
	}

	return sessions, total, nil
}

func (r *sessionRepository) CountSessionsByUserID(ctx context.Context, userID uuid.UUID) (int64, error) {
	var count int64
	query := `SELECT COUNT(*) FROM sessions WHERE user_id = $1`
	err := r.db.GetContext(ctx, &count, query, userID)
	return count, err
}

func (r *sessionRepository) DeleteSession(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM sessions WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

func (r *sessionRepository) DeleteSessionByToken(ctx context.Context, refreshToken string) error {
	query := `DELETE FROM sessions WHERE refresh_token = $1`
	_, err := r.db.ExecContext(ctx, query, refreshToken)
	return err
}

func (r *sessionRepository) DeleteSessionsByUserID(ctx context.Context, userID uuid.UUID) error {
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
