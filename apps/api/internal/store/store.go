package store

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type PaginationParams struct {
	Limit  int32 `json:"limit" query:"limit"`
	Offset int32 `json:"offset" query:"offset"`
}

func DefaultPagination() PaginationParams {
	return PaginationParams{
		Limit:  20,
		Offset: 0,
	}
}

func (p *PaginationParams) Normalize() {
	if p.Limit <= 0 {
		p.Limit = 20
	}
	if p.Limit > 100 {
		p.Limit = 100
	}
	if p.Offset < 0 {
		p.Offset = 0
	}
}

type PaginatedResponse[T any] struct {
	Data       []T            `json:"data"`
	Pagination PaginationInfo `json:"pagination"`
}

type PaginationInfo struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
	Total  int64 `json:"total"`
}

type ListSessionsFilter struct {
	UserID uuid.UUID
	PaginationParams
}

type DBTX interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	GetContext(ctx context.Context, dest any, query string, args ...any) error
	SelectContext(ctx context.Context, dest any, query string, args ...any) error
}

type Store struct {
	db         *sqlx.DB
	Users      UserRepository
	Sessions   SessionRepository
	Workspaces WorkspaceRepository
}

func New(db *sqlx.DB) *Store {
	return &Store{
		db:         db,
		Users:      NewUserRepository(db),
		Sessions:   NewSessionRepository(db),
		Workspaces: NewWorkspaceRepository(db),
	}
}

func (s *Store) ExecTx(ctx context.Context, fn func(*Store) error) error {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("begin tx: %w", err)
	}

	txStore := &Store{
		db:         s.db,
		Users:      NewUserRepository(tx),
		Sessions:   NewSessionRepository(tx),
		Workspaces: NewWorkspaceRepository(tx),
	}

	if err := fn(txStore); err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}
