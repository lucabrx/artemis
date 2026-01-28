package config

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Redis    RedisConfig
	MinIO    MinIOConfig
	NATS     NATSConfig
	Token    TokenConfig
}

type ServerConfig struct {
	Port                    string
	Environment             string
	EnableOpenAPIValidation bool
}

type DatabaseConfig struct {
	Host            string
	Port            string
	User            string
	Password        string
	DBName          string
	SSLMode         string
	MaxConns        int32
	MinConns        int32
	MaxConnLifetime time.Duration
	MaxConnIdleTime time.Duration
}

func (c DatabaseConfig) DSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Host, c.Port, c.User, c.Password, c.DBName, c.SSLMode)
}

type RedisConfig struct {
	Host     string
	Port     string
	Password string
	DB       int
}

func (c RedisConfig) Addr() string {
	return c.Host + ":" + c.Port
}

type MinIOConfig struct {
	Endpoint        string
	AccessKeyID     string
	SecretAccessKey string
	UseSSL          bool
	BucketName      string
}

type TokenConfig struct {
	SymmetricKey         string
	AccessTokenDuration  time.Duration
	RefreshTokenDuration time.Duration
}

type NATSConfig struct {
	URL string
}

func Load() (*Config, error) {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AddConfigPath("../..")
	viper.AutomaticEnv()

	viper.SetDefault("SERVER_PORT", "8080")
	viper.SetDefault("SERVER_ENVIRONMENT", "development")
	viper.SetDefault("ENABLE_OPENAPI_VALIDATION", false)
	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_PORT", "5432")
	viper.SetDefault("DB_USER", "artemis")
	viper.SetDefault("DB_PASSWORD", "artemis")
	viper.SetDefault("DB_NAME", "artemis")
	viper.SetDefault("DB_SSLMODE", "disable")
	viper.SetDefault("DB_MAX_CONNS", 25)
	viper.SetDefault("DB_MIN_CONNS", 5)
	viper.SetDefault("DB_MAX_CONN_LIFETIME", "1h")
	viper.SetDefault("DB_MAX_CONN_IDLE_TIME", "30m")
	viper.SetDefault("REDIS_HOST", "localhost")
	viper.SetDefault("REDIS_PORT", "6379")
	viper.SetDefault("REDIS_PASSWORD", "")
	viper.SetDefault("REDIS_DB", 0)
	viper.SetDefault("MINIO_ENDPOINT", "localhost:9000")
	viper.SetDefault("MINIO_ACCESS_KEY", "minioadmin")
	viper.SetDefault("MINIO_SECRET_KEY", "minioadmin")
	viper.SetDefault("MINIO_USE_SSL", false)
	viper.SetDefault("MINIO_BUCKET", "artemis")
	viper.SetDefault("TOKEN_SYMMETRIC_KEY", "12345678901234567890123456789012")
	viper.SetDefault("ACCESS_TOKEN_DURATION", "15m")
	viper.SetDefault("REFRESH_TOKEN_DURATION", "168h")
	viper.SetDefault("NATS_URL", "nats://localhost:4222")

	_ = viper.ReadInConfig()

	accessDuration, err := time.ParseDuration(viper.GetString("ACCESS_TOKEN_DURATION"))
	if err != nil {
		accessDuration = 15 * time.Minute
	}

	refreshDuration, err := time.ParseDuration(viper.GetString("REFRESH_TOKEN_DURATION"))
	if err != nil {
		refreshDuration = 7 * 24 * time.Hour
	}

	dbMaxConnLifetime, err := time.ParseDuration(viper.GetString("DB_MAX_CONN_LIFETIME"))
	if err != nil {
		dbMaxConnLifetime = time.Hour
	}

	dbMaxConnIdleTime, err := time.ParseDuration(viper.GetString("DB_MAX_CONN_IDLE_TIME"))
	if err != nil {
		dbMaxConnIdleTime = 30 * time.Minute
	}

	cfg := &Config{
		Server: ServerConfig{
			Port:                    viper.GetString("SERVER_PORT"),
			Environment:             viper.GetString("SERVER_ENVIRONMENT"),
			EnableOpenAPIValidation: viper.GetBool("ENABLE_OPENAPI_VALIDATION"),
		},
		Database: DatabaseConfig{
			Host:            viper.GetString("DB_HOST"),
			Port:            viper.GetString("DB_PORT"),
			User:            viper.GetString("DB_USER"),
			Password:        viper.GetString("DB_PASSWORD"),
			DBName:          viper.GetString("DB_NAME"),
			SSLMode:         viper.GetString("DB_SSLMODE"),
			MaxConns:        viper.GetInt32("DB_MAX_CONNS"),
			MinConns:        viper.GetInt32("DB_MIN_CONNS"),
			MaxConnLifetime: dbMaxConnLifetime,
			MaxConnIdleTime: dbMaxConnIdleTime,
		},
		Redis: RedisConfig{
			Host:     viper.GetString("REDIS_HOST"),
			Port:     viper.GetString("REDIS_PORT"),
			Password: viper.GetString("REDIS_PASSWORD"),
			DB:       viper.GetInt("REDIS_DB"),
		},
		MinIO: MinIOConfig{
			Endpoint:        viper.GetString("MINIO_ENDPOINT"),
			AccessKeyID:     viper.GetString("MINIO_ACCESS_KEY"),
			SecretAccessKey: viper.GetString("MINIO_SECRET_KEY"),
			UseSSL:          viper.GetBool("MINIO_USE_SSL"),
			BucketName:      viper.GetString("MINIO_BUCKET"),
		},
		Token: TokenConfig{
			SymmetricKey:         viper.GetString("TOKEN_SYMMETRIC_KEY"),
			AccessTokenDuration:  accessDuration,
			RefreshTokenDuration: refreshDuration,
		},
		NATS: NATSConfig{
			URL: viper.GetString("NATS_URL"),
		},
	}

	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	return cfg, nil
}

func (c *Config) Validate() error {
	var missing []string

	if c.Token.SymmetricKey == "" {
		missing = append(missing, "TOKEN_SYMMETRIC_KEY")
	} else if len(c.Token.SymmetricKey) != 32 {
		return errors.New("TOKEN_SYMMETRIC_KEY must be exactly 32 bytes")
	}

	if c.Server.Environment == "production" {
		if c.Token.SymmetricKey == "12345678901234567890123456789012" {
			return errors.New("TOKEN_SYMMETRIC_KEY must be changed in production")
		}
		if c.Database.Password == "artemis" {
			return errors.New("DB_PASSWORD must be changed in production")
		}
	}

	if c.Database.Host == "" {
		missing = append(missing, "DB_HOST")
	}
	if c.Redis.Host == "" {
		missing = append(missing, "REDIS_HOST")
	}
	if c.MinIO.Endpoint == "" {
		missing = append(missing, "MINIO_ENDPOINT")
	}

	if len(missing) > 0 {
		return fmt.Errorf("missing required environment variables: %s", strings.Join(missing, ", "))
	}

	return nil
}
