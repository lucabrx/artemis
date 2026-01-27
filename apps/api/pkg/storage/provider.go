package storage

import (
	"context"
	"io"
	"time"
)

type Provider interface {
	Upload(ctx context.Context, bucketType BucketType, objectName string, reader io.Reader, size int64, contentType string) (string, error)
	GetPublicURL(objectName string) string
	GetPresignedURL(ctx context.Context, objectName string, expiry time.Duration) (string, error)
	Delete(ctx context.Context, bucketType BucketType, objectName string) error
	UploadAvatar(ctx context.Context, userID string, reader io.Reader, size int64, contentType string) (string, error)
	UploadProjectFile(ctx context.Context, projectID, fileName string, reader io.Reader, size int64, contentType string) (string, error)
	GetProjectFileURL(ctx context.Context, projectID, fileName string, expiry time.Duration) (string, error)
}
