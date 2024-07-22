package project_init

import (
	"github.com/rs/zerolog/log"
)

func Init_main() (err error) {
	err = load_zerolog(true)
	if err != nil {
		log.Error().Err(err).Msg("Could not load zerolog")
		return err
	}

	err = load_dotenv(true)
	if err != nil {
		log.Error().Err(err).Msg("Could not initialised dotenv")
		return err
	}
	return err
}
