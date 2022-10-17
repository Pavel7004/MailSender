package main

import (
	"github.com/rs/zerolog/log"

	"github.com/Pavel7004/MailSender/pkg/adapter/db/mongo"
	"github.com/Pavel7004/MailSender/pkg/infra/config"
)

func main() {
	cfg, err := config.Get()
	if err != nil {
		log.Error().Err(err).Msg("Failed to read config")
		return
	}

	db, err := mongo.New(cfg)
	defer func() {
		if err := db.Close(); err != nil {
			log.Error().Err(err).Msg("Failed to close connection to db")
		}
	}()
	if err != nil {
		log.Error().Err(err).Msg("Failed to connect to DB")
		return
	}
}
