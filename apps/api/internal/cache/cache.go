package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/lukabrkovic/artemis/internal/store"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
)

type UserCache interface {
	GetUser(ctx context.Context, id uuid.UUID) (*store.User, error)
	SetUser(ctx context.Context, user *store.User) error
	DeleteUser(ctx context.Context, id uuid.UUID) error
}

type Cache struct {
	client *redis.Client
	ttl    time.Duration
	logger zerolog.Logger
}

func New(client *redis.Client, ttl time.Duration, logger zerolog.Logger) *Cache {
	return &Cache{
		client: client,
		ttl:    ttl,
		logger: logger.With().Str("component", "cache").Logger(),
	}
}

var _ UserCache = (*Cache)(nil)

func (c *Cache) userKey(id uuid.UUID) string {
	return fmt.Sprintf("user:%s", id.String())
}

func (c *Cache) GetUser(ctx context.Context, id uuid.UUID) (*store.User, error) {
	data, err := c.client.Get(ctx, c.userKey(id)).Bytes()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		c.logger.Warn().Err(err).Str("user_id", id.String()).Msg("failed to get user from cache")
		return nil, err
	}

	var user store.User
	if err := json.Unmarshal(data, &user); err != nil {
		c.logger.Error().Err(err).Str("user_id", id.String()).Msg("failed to unmarshal user from cache")
		return nil, err
	}

	return &user, nil
}

func (c *Cache) SetUser(ctx context.Context, user *store.User) error {
	data, err := json.Marshal(user)
	if err != nil {
		c.logger.Error().Err(err).Str("user_id", user.ID.String()).Msg("failed to marshal user for cache")
		return err
	}

	if err := c.client.Set(ctx, c.userKey(user.ID), data, c.ttl).Err(); err != nil {
		c.logger.Warn().Err(err).Str("user_id", user.ID.String()).Msg("failed to set user in cache")
		return err
	}

	c.logger.Debug().Str("user_id", user.ID.String()).Msg("user cached successfully")
	return nil
}

func (c *Cache) DeleteUser(ctx context.Context, id uuid.UUID) error {
	if err := c.client.Del(ctx, c.userKey(id)).Err(); err != nil {
		c.logger.Warn().Err(err).Str("user_id", id.String()).Msg("failed to delete user from cache")
		return err
	}
	c.logger.Debug().Str("user_id", id.String()).Msg("user removed from cache")
	return nil
}
