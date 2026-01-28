package storage

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/lukabrkovic/artemis/internal/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type BucketType string

const (
	BucketPublic  BucketType = "public"
	BucketPrivate BucketType = "private"
)

type MinIO struct {
	client        *minio.Client
	publicBucket  string
	privateBucket string
	endpoint      string
	useSSL        bool
}

func NewMinIO(cfg config.MinIOConfig) (*MinIO, error) {
	client, err := minio.New(cfg.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.AccessKeyID, cfg.SecretAccessKey, ""),
		Secure: cfg.UseSSL,
	})
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	publicBucket := cfg.BucketName + "-public"
	privateBucket := cfg.BucketName + "-private"

	if err := ensureBucket(ctx, client, publicBucket); err != nil {
		return nil, err
	}

	if err := setPublicPolicy(ctx, client, publicBucket); err != nil {
		return nil, err
	}

	if err := ensureBucket(ctx, client, privateBucket); err != nil {
		return nil, err
	}

	return &MinIO{
		client:        client,
		publicBucket:  publicBucket,
		privateBucket: privateBucket,
		endpoint:      cfg.Endpoint,
		useSSL:        cfg.UseSSL,
	}, nil
}

func ensureBucket(ctx context.Context, client *minio.Client, bucketName string) error {
	exists, err := client.BucketExists(ctx, bucketName)
	if err != nil {
		return err
	}
	if !exists {
		return client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
	}
	return nil
}

func setPublicPolicy(ctx context.Context, client *minio.Client, bucketName string) error {
	policy := fmt.Sprintf(`{
		"Version": "2012-10-17",
		"Statement": [{
			"Effect": "Allow",
			"Principal": {"AWS": ["*"]},
			"Action": ["s3:GetObject"],
			"Resource": ["arn:aws:s3:::%s/*"]
		}]
	}`, bucketName)

	return client.SetBucketPolicy(ctx, bucketName, policy)
}

func (m *MinIO) bucket(bucketType BucketType) string {
	if bucketType == BucketPublic {
		return m.publicBucket
	}
	return m.privateBucket
}

func (m *MinIO) Upload(ctx context.Context, bucketType BucketType, objectName string, reader io.Reader, size int64, contentType string) (string, error) {
	bucket := m.bucket(bucketType)
	_, err := m.client.PutObject(ctx, bucket, objectName, reader, size, minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		return "", err
	}
	return objectName, nil
}

func (m *MinIO) GetPublicURL(objectName string) string {
	scheme := "http"
	if m.useSSL {
		scheme = "https"
	}
	return fmt.Sprintf("%s://%s/%s/%s", scheme, m.endpoint, m.publicBucket, objectName)
}

func (m *MinIO) GetPresignedURL(ctx context.Context, objectName string, expiry time.Duration) (string, error) {
	url, err := m.client.PresignedGetObject(ctx, m.privateBucket, objectName, expiry, nil)
	if err != nil {
		return "", err
	}
	return url.String(), nil
}

func (m *MinIO) Delete(ctx context.Context, bucketType BucketType, objectName string) error {
	return m.client.RemoveObject(ctx, m.bucket(bucketType), objectName, minio.RemoveObjectOptions{})
}

func (m *MinIO) UploadAvatar(ctx context.Context, workspaceID string, reader io.Reader, size int64, contentType string) (string, error) {
	objectName := fmt.Sprintf("avatars/%s", workspaceID)
	_, err := m.Upload(ctx, BucketPublic, objectName, reader, size, contentType)
	if err != nil {
		return "", err
	}
	return m.GetPublicURL(objectName), nil
}

func (m *MinIO) DeleteAvatar(ctx context.Context, workspaceID string) error {
	objectName := fmt.Sprintf("avatars/%s", workspaceID)
	return m.Delete(ctx, BucketPublic, objectName)
}

func (m *MinIO) UploadProjectFile(ctx context.Context, projectID, fileName string, reader io.Reader, size int64, contentType string) (string, error) {
	objectName := fmt.Sprintf("projects/%s/%s", projectID, fileName)
	return m.Upload(ctx, BucketPrivate, objectName, reader, size, contentType)
}

func (m *MinIO) GetProjectFileURL(ctx context.Context, projectID, fileName string, expiry time.Duration) (string, error) {
	objectName := fmt.Sprintf("projects/%s/%s", projectID, fileName)
	return m.GetPresignedURL(ctx, objectName, expiry)
}
