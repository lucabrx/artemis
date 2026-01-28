package events

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
	"github.com/rs/zerolog"
)

type EventType string

const (
	EventUserRegistered     EventType = "user.registered"
	EventUserLoggedIn       EventType = "user.logged_in"
	EventUserUpdated        EventType = "user.updated"
	EventWorkspaceCreated   EventType = "workspace.created"
	EventWorkspaceUpdated   EventType = "workspace.updated"
	EventWorkspaceDeleted   EventType = "workspace.deleted"
	EventMemberAdded        EventType = "member.added"
	EventMemberRemoved      EventType = "member.removed"
	EventEmailSendRequested EventType = "email.send_requested"
)

type Event struct {
	ID        string    `json:"id"`
	Type      EventType `json:"type"`
	Timestamp time.Time `json:"timestamp"`
	UserID    uuid.UUID `json:"user_id,omitempty"`
	Payload   any       `json:"payload"`
}

type Bus struct {
	conn   *nats.Conn
	logger zerolog.Logger
}

func NewBus(natsURL string, logger zerolog.Logger) (*Bus, error) {
	conn, err := nats.Connect(natsURL, nats.Timeout(5*time.Second))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to NATS: %w", err)
	}

	return &Bus{
		conn:   conn,
		logger: logger.With().Str("component", "event_bus").Logger(),
	}, nil
}

func (b *Bus) Close() {
	b.conn.Close()
}

func (b *Bus) Publish(ctx context.Context, eventType EventType, userID uuid.UUID, payload any) error {
	if b == nil || b.conn == nil {
		return nil
	}

	event := Event{
		ID:        uuid.New().String(),
		Type:      eventType,
		Timestamp: time.Now().UTC(),
		UserID:    userID,
		Payload:   payload,
	}

	data, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("failed to marshal event: %w", err)
	}

	subject := fmt.Sprintf("artemis.%s", eventType)
	if err := b.conn.Publish(subject, data); err != nil {
		return fmt.Errorf("failed to publish event: %w", err)
	}

	b.logger.Debug().
		Str("event_id", event.ID).
		Str("type", string(eventType)).
		Str("subject", subject).
		Msg("event published")

	return nil
}

func (b *Bus) Subscribe(eventType EventType, handler func(*Event) error) (*nats.Subscription, error) {
	if b == nil || b.conn == nil {
		return nil, fmt.Errorf("event bus not connected")
	}

	subject := fmt.Sprintf("artemis.%s", eventType)
	sub, err := b.conn.Subscribe(subject, func(msg *nats.Msg) {
		var event Event
		if err := json.Unmarshal(msg.Data, &event); err != nil {
			b.logger.Error().Err(err).Msg("failed to unmarshal event")
			return
		}

		if err := handler(&event); err != nil {
			b.logger.Error().
				Err(err).
				Str("event_id", event.ID).
				Str("type", string(event.Type)).
				Msg("event handler failed")
		}
	})

	if err != nil {
		return nil, fmt.Errorf("failed to subscribe: %w", err)
	}

	b.logger.Info().
		Str("subject", subject).
		Msg("subscribed to events")

	return sub, nil
}
