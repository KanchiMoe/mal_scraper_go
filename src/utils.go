package src

import (
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

// A utility function that will generate a UUID string. UUIDs are used a lot in SQL queries, so rather than repeatdly
// use type this code, this function serves as a way to generate a UUID by calling it when needed.
func Generate_uuid() (generated_uuid string) {
	generated_uuid = uuid.New().String()
	log.Trace().Str("uuid", generated_uuid).Msg("Created UUID")
	return generated_uuid
}

func Generate_timestamp() (timestamp string, err error) {
	location, err := time.LoadLocation("Europe/London")
	if err != nil {
		log.Error().Err(err).Msg("Error with timestamp location")
		return "", err
	}
	current_time := time.Now().In(location)

	const layout = "2006-01-02 15:04:05.999999-07"
	formatted := current_time.Format(layout)
	log.Trace().Str("timestamp", formatted).Msg("Created timestamp")

	return formatted, err
}
