package audit

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Action string

const (
	ActionCreate Action = "create"
	ActionUpdate Action = "update"
	ActionDelete Action = "delete"
	ActionLogin  Action = "login"
	ActionLogout Action = "logout"
)

type Log struct {
	ID         uuid.UUID       `json:"id" db:"id"`
	UserID     *uuid.UUID      `json:"user_id" db:"user_id"`
	Action     Action          `json:"action" db:"action"`
	EntityType string          `json:"entity_type" db:"entity_type"`
	EntityID   string          `json:"entity_id" db:"entity_id"`
	OldValue   json.RawMessage `json:"old_value,omitempty" db:"old_value"`
	NewValue   json.RawMessage `json:"new_value,omitempty" db:"new_value"`
	IP         string          `json:"ip" db:"ip_address"`
	UserAgent  string          `json:"user_agent" db:"user_agent"`
	Timestamp  time.Time       `json:"timestamp" db:"created_at"`
}
