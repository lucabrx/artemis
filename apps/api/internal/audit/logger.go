package audit

import (
	"context"
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

type Store interface {
	CreateAuditLog(ctx context.Context, log *Log) error
}

type Logger struct {
	store  Store
	logger zerolog.Logger
}

func NewLogger(store Store, logger zerolog.Logger) *Logger {
	return &Logger{
		store:  store,
		logger: logger.With().Str("component", "audit").Logger(),
	}
}

func (l *Logger) Log(ctx context.Context, userID *uuid.UUID, action Action, entityType string, entityID string, oldVal, newVal any, ip, userAgent string) error {
	var oldBytes, newBytes json.RawMessage
	var err error

	if oldVal != nil {
		oldBytes, err = json.Marshal(oldVal)
		if err != nil {
			l.logger.Warn().Err(err).Msg("failed to marshal old value")
		}
	}

	if newVal != nil {
		newBytes, err = json.Marshal(newVal)
		if err != nil {
			l.logger.Warn().Err(err).Msg("failed to marshal new value")
		}
	}

	log := &Log{
		ID:         uuid.New(),
		UserID:     userID,
		Action:     action,
		EntityType: entityType,
		EntityID:   entityID,
		OldValue:   oldBytes,
		NewValue:   newBytes,
		IP:         ip,
		UserAgent:  userAgent,
		Timestamp:  time.Now().UTC(),
	}

	if err := l.store.CreateAuditLog(ctx, log); err != nil {
		l.logger.Error().Err(err).Msg("failed to write audit log")
		return err
	}

	return nil
}
