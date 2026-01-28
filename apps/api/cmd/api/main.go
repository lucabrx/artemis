package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/lukabrkovic/artemis/internal/audit"
	"github.com/lukabrkovic/artemis/internal/cache"
	"github.com/lukabrkovic/artemis/internal/config"
	"github.com/lukabrkovic/artemis/internal/database"
	"github.com/lukabrkovic/artemis/internal/events"
	"github.com/lukabrkovic/artemis/internal/router"
	"github.com/lukabrkovic/artemis/internal/store"
	"github.com/lukabrkovic/artemis/pkg/logger"
	"github.com/lukabrkovic/artemis/pkg/storage"
	"github.com/lukabrkovic/artemis/pkg/token"
	"github.com/rs/zerolog"
)

// @title           Artemis API
// @version         1.0
// @description     Project management platform API
// @host            localhost:8080
// @BasePath        /api/v1
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	log := logger.New()

	cfg, err := config.Load()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to load config")
	}

	db, err := database.NewPostgres(cfg.Database)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect to postgres")
	}
	defer db.Close()

	if err := database.RunMigrations(db, cfg.Database, log); err != nil {
		log.Fatal().Err(err).Msg("failed to run migrations")
	}

	keydbClient, err := database.NewRedis(cfg.Redis)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect to keydb")
	}
	defer keydbClient.Close()

	minioClient, err := storage.NewMinIO(cfg.MinIO)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect to minio")
	}

	tokenMaker, err := token.NewPasetoMaker(cfg.Token.SymmetricKey, cfg.Token)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create token maker")
	}

	if cfg.Server.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	st := store.New(db)
	userCache := cache.New(keydbClient, 15*time.Minute, log)

	go collectDBStats(db, log)

	eventBus, err := events.NewBus(cfg.NATS.URL, log)
	if err != nil {
		log.Warn().Err(err).Msg("failed to connect to NATS, continuing without event publishing")
		eventBus = nil
	} else {
		defer eventBus.Close()
	}

	auditLogger := audit.NewLogger(st.AuditLogs, log)

	r := router.New(router.Config{
		Store:                   st,
		Cache:                   userCache,
		Storage:                 minioClient,
		TokenMaker:              tokenMaker,
		TokenConfig:             cfg.Token,
		Logger:                  log,
		Environment:             cfg.Server.Environment,
		EnableOpenAPIValidation: cfg.Server.EnableOpenAPIValidation,
		EventBus:                eventBus,
		AuditLogger:             auditLogger,
	})

	srv := &http.Server{
		Addr:         ":" + cfg.Server.Port,
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		log.Info().Str("port", cfg.Server.Port).Msg("starting server")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Err(err).Msg("server failed")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info().Msg("shutting down server")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal().Err(err).Msg("server forced to shutdown")
	}

	log.Info().Msg("server stopped")
}

func collectDBStats(db *sqlx.DB, log zerolog.Logger) {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		stats := db.Stats()
		log.Debug().
			Int("open_connections", stats.OpenConnections).
			Int("in_use", stats.InUse).
			Int("idle", stats.Idle).
			Int64("wait_count", stats.WaitCount).
			Dur("wait_duration", stats.WaitDuration).
			Int64("max_idle_closed", stats.MaxIdleClosed).
			Int64("max_lifetime_closed", stats.MaxLifetimeClosed).
			Msg("database stats")
	}
}
