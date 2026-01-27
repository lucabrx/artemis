package store

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type DBTX interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	GetContext(ctx context.Context, dest any, query string, args ...any) error
	SelectContext(ctx context.Context, dest any, query string, args ...any) error
}

type Queries struct {
	db DBTX
}

type Store struct {
	*Queries
	db *sqlx.DB
}

func New(db *sqlx.DB) *Store {
	return &Store{
		db:      db,
		Queries: &Queries{db: db},
	}
}

func (s *Store) ExecTx(ctx context.Context, fn func(Querier) error) error {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("begin tx: %w", err)
	}

	q := &Queries{db: tx}
	if err := fn(q); err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

type Repository interface {
	Querier
	ExecTx(ctx context.Context, fn func(Querier) error) error
}

type Querier interface {
	CreateUser(ctx context.Context, arg CreateUserParams) (*User, error)
	GetUserByID(ctx context.Context, id any) (*User, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (*User, error)
	DeleteUser(ctx context.Context, id any) error

	CreateSession(ctx context.Context, arg CreateSessionParams) (*Session, error)
	GetSessionByID(ctx context.Context, id any) (*Session, error)
	GetSessionByToken(ctx context.Context, refreshToken string) (*Session, error)
	GetSessionsByUserID(ctx context.Context, userID any) ([]Session, error)
	DeleteSession(ctx context.Context, id any) error
	DeleteSessionByToken(ctx context.Context, refreshToken string) error
	DeleteSessionsByUserID(ctx context.Context, userID any) error
	DeleteExpiredSessions(ctx context.Context) (int64, error)
}

var _ Querier = (*Queries)(nil)
