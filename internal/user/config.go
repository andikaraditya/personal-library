package user

import (
	"os"

	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type config struct {
	JWTSecret string `envconfig:"jwt_secret"`
}

var cfg config

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	var err error

	if err = envconfig.Process("", &cfg); err != nil {
		log.Fatal().Err(err).Msg("")
	}

	log.Info().
		Interface("config", cfg).
		Msg("initialize user")
}
