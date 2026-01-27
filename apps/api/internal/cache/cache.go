package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/lukabrkovic/artemis/internal/store"
	"github.com/redis/go-redis/v9"
)

type UserCache interface {
	GetUser(ctx context.Context, id uuid.UUID) (*store.User, error)
	SetUser(ctx context.Context, user *store.User) error
	DeleteUser(ctx context.Context, id uuid.UUID) error
}

type Cache struct {
	client *redis.Client
	ttl    time.Duration
}

func New(client *redis.Client, ttl time.Duration) *Cache {
	return &Cache{
		client: client,
		ttl:    ttl,
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
		return nil, err
	}

	var user store.User
	if err := json.Unmarshal(data, &user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (c *Cache) SetUser(ctx context.Context, user *store.User) error {
	data, err := json.Marshal(user)
	if err != nil {
		return err
	}

	return c.client.Set(ctx, c.userKey(user.ID), data, c.ttl).Err()
}

func (c *Cache) DeleteUser(ctx context.Context, id uuid.UUID) error {
	return c.client.Del(ctx, c.userKey(id)).Err()
}
