package db

import (
	"os"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type config struct {
	// MaxConn, IddleConn can be set here.
	// Refer to https://pkg.go.dev/github.com/jackc/pgx/v4@v4.13.0/pgxpool#ParseConfig
	DSN          string        `envconfig:"dsn" required:"true"`
	PingInterval time.Duration `envconfig:"db_ping_interval" default:"15m"`
}

var cfg config

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatal().Err(err).Msg("")
	}

	log.Info().
		Interface("config", cfg).
		Msg("initialize postgres")
}
