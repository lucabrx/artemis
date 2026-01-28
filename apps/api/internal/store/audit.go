package store

import (
	"context"

	"github.com/google/uuid"
	"github.com/lukabrkovic/artemis/internal/audit"
)

type AuditLogRepository interface {
	Create(ctx context.Context, log *audit.Log) error
	CreateAuditLog(ctx context.Context, log *audit.Log) error
	GetByEntity(ctx context.Context, entityType string, entityID string) ([]audit.Log, error)
	GetByUser(ctx context.Context, userID uuid.UUID, filters FilterParams) ([]audit.Log, int64, error)
}

type auditLogRepository struct {
	db DBTX
}

func NewAuditLogRepository(db DBTX) AuditLogRepository {
	return &auditLogRepository{db: db}
}

func (r *auditLogRepository) Create(ctx context.Context, log *audit.Log) error {
	return r.CreateAuditLog(ctx, log)
}

func (r *auditLogRepository) CreateAuditLog(ctx context.Context, log *audit.Log) error {
	query := `
		INSERT INTO audit_logs (id, user_id, action, entity_type, entity_id, old_value, new_value, ip_address, user_agent, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`
	_, err := r.db.ExecContext(ctx, query,
		log.ID, log.UserID, log.Action, log.EntityType, log.EntityID,
		log.OldValue, log.NewValue, log.IP, log.UserAgent, log.Timestamp,
	)
	return err
}

func (r *auditLogRepository) GetByEntity(ctx context.Context, entityType string, entityID string) ([]audit.Log, error) {
	var logs []audit.Log
	query := `
		SELECT * FROM audit_logs 
		WHERE entity_type = $1 AND entity_id = $2 
		ORDER BY created_at DESC
	`
	err := r.db.SelectContext(ctx, &logs, query, entityType, entityID)
	return logs, err
}

func (r *auditLogRepository) GetByUser(ctx context.Context, userID uuid.UUID, filters FilterParams) ([]audit.Log, int64, error) {
	var logs []audit.Log
	query := `
		SELECT * FROM audit_logs 
		WHERE user_id = $1 
		ORDER BY created_at DESC 
		LIMIT $2 OFFSET $3
	`
	err := r.db.SelectContext(ctx, &logs, query, userID, filters.Limit, filters.Offset)
	if err != nil {
		return nil, 0, err
	}

	var total int64
	countQuery := `SELECT COUNT(*) FROM audit_logs WHERE user_id = $1`
	err = r.db.GetContext(ctx, &total, countQuery, userID)
	if err != nil {
		return nil, 0, err
	}

	return logs, total, nil
}
