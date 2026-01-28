package database

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/jmoiron/sqlx"
	"github.com/lukabrkovic/artemis/internal/config"
	"github.com/pressly/goose/v3"
	"github.com/rs/zerolog"
)

func RunMigrations(db *sqlx.DB, cfg config.DatabaseConfig, logger zerolog.Logger) error {
	migrationsDir := getMigrationsPath()
	logger.Debug().Str("path", migrationsDir).Msg("loading migrations")

	goose.SetLogger(&gooseLogger{logger: logger})
	goose.SetDialect("postgres")

	if err := goose.Up(db.DB, migrationsDir); err != nil {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	version, err := goose.GetDBVersion(db.DB)
	if err != nil {
		return fmt.Errorf("failed to get migration version: %w", err)
	}

	logger.Info().
		Int64("version", version).
		Msg("database migrations completed")

	return nil
}

func getMigrationsPath() string {
	_, b, _, _ := runtime.Caller(0)
	basePath := filepath.Dir(b)
	migrationsDir := filepath.Join(basePath, "..", "..", "migrations")

	if _, err := os.Stat(migrationsDir); os.IsNotExist(err) {
		return "migrations"
	}

	absPath, _ := filepath.Abs(migrationsDir)
	return absPath
}

type gooseLogger struct {
	logger zerolog.Logger
}

func (g *gooseLogger) Printf(format string, v ...interface{}) {
	g.logger.Info().Msg(fmt.Sprintf(format, v...))
}

func (g *gooseLogger) Fatalf(format string, v ...interface{}) {
	g.logger.Fatal().Msg(fmt.Sprintf(format, v...))
}
