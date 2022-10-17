package main

import (
	"flag"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/Pavel7004/MailSender/pkg/infra/config"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})

	filename := flag.String("filename", "config.yaml", "Path to write file")

	flag.Parse()

	if _, err := config.Get(); err != nil {
		log.Error().Err(err).Msg("Failed to read config")
		return
	}

	if err := config.WriteToFile(*filename); err != nil {
		log.Error().Err(err).Msg("Failed to write config")
		return
	}

	log.Info().Msgf("Written %q", *filename)
}
