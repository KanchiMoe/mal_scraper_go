package project_init

import (
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func load_zerolog(run bool) (err error) {
	// run if true
	if run {
		// set log level
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
		zerolog.DurationFieldUnit = time.Second

		log.Trace().Msg("Zerolog initialised")
		return nil
	} else {
		fmt.Println("Error initialising zerolog")
	}

	return nil

}
