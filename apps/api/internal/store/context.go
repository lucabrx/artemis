package store

import (
	"context"
	"errors"
)

var ErrContextCancelled = errors.New("request was cancelled or timed out")

func CheckContext(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ErrContextCancelled
	default:
		return nil
	}
}
