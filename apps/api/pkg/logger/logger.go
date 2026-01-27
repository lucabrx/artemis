package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

func New() zerolog.Logger {
	return zerolog.New(zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
	}).With().Timestamp().Caller().Logger()
}

func NewProduction() zerolog.Logger {
	return zerolog.New(os.Stdout).With().Timestamp().Logger()
}
