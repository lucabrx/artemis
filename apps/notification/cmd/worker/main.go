package main

import (
	"encoding/json"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/rs/zerolog"
)

type Event struct {
	ID        string    `json:"id"`
	Type      string    `json:"type"`
	Timestamp time.Time `json:"timestamp"`
	UserID    string    `json:"user_id,omitempty"`
	Payload   any       `json:"payload"`
}

func main() {
	log := zerolog.New(os.Stdout).With().Timestamp().Logger()

	natsURL := os.Getenv("NATS_URL")
	if natsURL == "" {
		natsURL = nats.DefaultURL
	}

	log.Info().Str("url", natsURL).Msg("connecting to NATS")

	nc, err := nats.Connect(natsURL, nats.Timeout(10*time.Second))
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect to NATS")
	}
	defer nc.Close()

	log.Info().Msg("connected to NATS")

	subjects := []string{
		"artemis.user.registered",
		"artemis.user.logged_in",
		"artemis.workspace.created",
		"artemis.workspace.updated",
		"artemis.workspace.deleted",
		"artemis.member.added",
		"artemis.member.removed",
		"artemis.email.send_requested",
	}

	var subs []*nats.Subscription

	for _, subject := range subjects {
		sub, err := nc.Subscribe(subject, func(msg *nats.Msg) {
			handleEvent(log, msg)
		})
		if err != nil {
			log.Fatal().Err(err).Str("subject", subject).Msg("failed to subscribe")
		}
		subs = append(subs, sub)
		log.Info().Str("subject", subject).Msg("subscribed")
	}

	log.Info().Msg("notification worker started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info().Msg("shutting down")

	for _, sub := range subs {
		sub.Unsubscribe()
	}
}

func handleEvent(log zerolog.Logger, msg *nats.Msg) {
	var event Event
	if err := json.Unmarshal(msg.Data, &event); err != nil {
		log.Error().Err(err).Str("subject", msg.Subject).Msg("failed to unmarshal event")
		return
	}

	logger := log.With().
		Str("event_id", event.ID).
		Str("type", event.Type).
		Str("user_id", event.UserID).
		Str("subject", msg.Subject).
		Logger()

	switch event.Type {
	case "user.registered":
		logger.Info().Interface("payload", event.Payload).Msg("user registered - would send welcome email")
	case "user.logged_in":
		logger.Info().Interface("payload", event.Payload).Msg("user logged in")
	case "workspace.created":
		logger.Info().Interface("payload", event.Payload).Msg("workspace created")
	case "workspace.updated":
		logger.Info().Interface("payload", event.Payload).Msg("workspace updated")
	case "workspace.deleted":
		logger.Info().Interface("payload", event.Payload).Msg("workspace deleted")
	case "member.added":
		logger.Info().Interface("payload", event.Payload).Msg("member added - would send invitation email")
	case "member.removed":
		logger.Info().Interface("payload", event.Payload).Msg("member removed")
	case "email.send_requested":
		logger.Info().Interface("payload", event.Payload).Msg("email send requested")
	default:
		logger.Info().Interface("payload", event.Payload).Msg("received event")
	}
}
