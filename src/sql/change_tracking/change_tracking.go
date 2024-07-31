package changetracking

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kanchimoe/mal_scraper_go/src"
	"github.com/rs/zerolog/log"
)

const error_location string = "sql/change_tracking"

func Add_to_change_tracking(db_connection *pgxpool.Pool, category string, item_id int, field string, old_value string, new_value string) (err error) {
	// vars
	uuid := src.Generate_uuid()
	timestamp, err := src.Generate_timestamp()
	if err != nil {
		log.Error().Err(err).Str("location", error_location).Msg("Error with timestamp")
		return err
	}

	// category validation

	// query
	const sql_query string = `
		INSERT INTO change_tracking
		(uuid, category, item_id, field, old_value, new_value, timestamp)
		VALUES ($1, $2, $3, $4, $5, $6, $7)`

	// debug
	log.Trace().
		Str("uuid", uuid).
		Str("timestamp", timestamp).
		Str("category", category).
		Int("item_id", item_id).
		Str("field", field).
		Str("old_value", old_value).
		Str("new_value", new_value).
		Msg("Change tracking")

	// query results
	_, err = db_connection.Exec(context.Background(), sql_query, uuid, category, item_id, field, old_value, new_value, timestamp)
	if err != nil {
		log.Error().Err(err).Str("location", error_location)
		return err
	}

	return nil
}

// func category_validation(category string) (err error) {

// }
