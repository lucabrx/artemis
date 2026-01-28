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

type Store struct {
	db         *sqlx.DB
	Users      UserRepository
	Sessions   SessionRepository
	Workspaces WorkspaceRepository
	AuditLogs  AuditLogRepository
}

func New(db *sqlx.DB) *Store {
	return &Store{
		db:         db,
		Users:      NewUserRepository(db),
		Sessions:   NewSessionRepository(db),
		Workspaces: NewWorkspaceRepository(db),
		AuditLogs:  NewAuditLogRepository(db),
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
		AuditLogs:  NewAuditLogRepository(tx),
	}

	if err := fn(txStore); err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}
