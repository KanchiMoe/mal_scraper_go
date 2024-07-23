package project_init

import (
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func load_dotenv(run bool) (err error) {
	// if run is true
	if run {
		log.Trace().Msg("Attempting to load .env file")

		err = godotenv.Load()
		if err != nil {
			log.Err(err).Msg("Could not load .env file")
			return err

		} else {
			log.Trace().Msg("Was able to load .env file")
			return nil
		}
	} else {
		log.Trace().Msg("Run .env set to false, skipping")
		return nil
	}
}
