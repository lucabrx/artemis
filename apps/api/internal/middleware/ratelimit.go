package middleware

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lukabrkovic/artemis/pkg/apperr"
)

type RateLimiterConfig struct {
	Requests int
	Window   time.Duration
	KeyFunc  func(*gin.Context) string
}

func DefaultRateLimiterConfig() RateLimiterConfig {
	return RateLimiterConfig{
		Requests: 100,
		Window:   time.Minute,
		KeyFunc: func(c *gin.Context) string {
			return c.ClientIP()
		},
	}
}

func AuthRateLimiterConfig() RateLimiterConfig {
	return RateLimiterConfig{
		Requests: 5,
		Window:   time.Minute,
		KeyFunc: func(c *gin.Context) string {
			return c.ClientIP()
		},
	}
}

type bucket struct {
	tokens     int
	lastRefill time.Time
}

type MemoryRateLimiter struct {
	config  RateLimiterConfig
	buckets map[string]*bucket
	mu      sync.RWMutex
}

func NewMemoryRateLimiter(config RateLimiterConfig) *MemoryRateLimiter {
	rl := &MemoryRateLimiter{
		config:  config,
		buckets: make(map[string]*bucket),
	}
	go rl.cleanup()
	return rl
}

func (rl *MemoryRateLimiter) Allow(key string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	b, exists := rl.buckets[key]

	if !exists {
		rl.buckets[key] = &bucket{
			tokens:     rl.config.Requests - 1,
			lastRefill: now,
		}
		return true
	}

	elapsed := now.Sub(b.lastRefill)
	tokensToAdd := int(elapsed/rl.config.Window) * rl.config.Requests

	if tokensToAdd > 0 {
		b.tokens = min(b.tokens+tokensToAdd, rl.config.Requests)
		b.lastRefill = now
	}

	if b.tokens > 0 {
		b.tokens--
		return true
	}

	return false
}

func (rl *MemoryRateLimiter) cleanup() {
	ticker := time.NewTicker(rl.config.Window * 2)
	defer ticker.Stop()

	for range ticker.C {
		rl.mu.Lock()
		now := time.Now()
		for key, b := range rl.buckets {
			if now.Sub(b.lastRefill) > rl.config.Window*2 {
				delete(rl.buckets, key)
			}
		}
		rl.mu.Unlock()
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func RateLimiter(limiter *MemoryRateLimiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := limiter.config.KeyFunc(c)

		if !limiter.Allow(key) {
			c.Error(apperr.New(http.StatusTooManyRequests, "rate limit exceeded"))
			c.Abort()
			return
		}

		c.Next()
	}
}

func RateLimiterByIP(requests int, window time.Duration) gin.HandlerFunc {
	config := RateLimiterConfig{
		Requests: requests,
		Window:   window,
		KeyFunc:  func(c *gin.Context) string { return c.ClientIP() },
	}
	limiter := NewMemoryRateLimiter(config)
	return RateLimiter(limiter)
}

func RateLimiterForAuth() gin.HandlerFunc {
	return RateLimiterByIP(5, time.Minute)
}

func RateLimiterWithUserID(requests int, window time.Duration) gin.HandlerFunc {
	config := RateLimiterConfig{
		Requests: requests,
		Window:   window,
		KeyFunc: func(c *gin.Context) string {
			userID, exists := c.Get("user_id")
			if !exists {
				return c.ClientIP()
			}
			return fmt.Sprintf("user:%v", userID)
		},
	}
	limiter := NewMemoryRateLimiter(config)
	return RateLimiter(limiter)
}
