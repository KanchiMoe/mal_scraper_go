package sql_users

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kanchimoe/mal_scraper_go/src"
	"github.com/kanchimoe/mal_scraper_go/src/project_structs"
	changetracking "github.com/kanchimoe/mal_scraper_go/src/sql/change_tracking"
	"github.com/rs/zerolog/log"
)

const error_location string = "sql/users"

func Get_id_from_username(db_connection *pgxpool.Pool, username string) (query_results project_structs.Username_and_id) {
	log.Trace().Str("username", username).Msg("SQL: search for username")

	// query
	var sql_query string = "SELECT id, username FROM users WHERE username = $1"

	// query results
	err := db_connection.QueryRow(context.Background(), sql_query, username).Scan(
		&query_results.Id, &query_results.Username,
	)

	// no results
	if err != nil {
		log.Debug().Str("username", username).Msg("SQL results: Username not in the database")
		query_results.In_db = false
		return query_results
	}

	// are results
	log.Debug().Str("username", query_results.Username).Int("id", query_results.Id).Msg("SQL results: username is in the database")
	query_results.In_db = true
	return query_results
}

func Get_username_from_id(db_connection *pgxpool.Pool, user_id int) (query_results project_structs.Username_and_id) {
	log.Trace().Int("id", user_id).Msg("SQL: search for id")

	// query
	var sql_query string = "SELECT id, username FROM users WHERE id = $1"

	// query results
	err := db_connection.QueryRow(context.Background(), sql_query, user_id).Scan(
		&query_results.Id, &query_results.Username,
	)

	// no results
	if err != nil {
		query_results.In_db = false
		log.Debug().Msg("SQL results: ID not in the database")
		return query_results
	}

	// are results
	log.Debug().Int("id", query_results.Id).Str("username", query_results.Username).Msg("SQL results: ID is in the database")
	query_results.In_db = true
	return query_results
}

func Add_user_id_to_db(db_connection *pgxpool.Pool, id int, username string) (err error) {
	log.Trace().Int("id", id).Str("username", username).Msg("SQL: add user to database")

	timestamp, err := src.Generate_timestamp()
	if err != nil {
		log.Error().Err(err).Str("location", error_location).Msg("Error calling generate timestamp")
	}

	// query
	const sql_insert_new_user_id string = "INSERT INTO users (id, username, last_interacted) VALUES ($1, $2, $3);"

	// query results
	_, err = db_connection.Exec(context.Background(), sql_insert_new_user_id, id, username, timestamp)
	if err != nil {
		log.Error().Err(err).Msg("error query") //////////// TO DO, INVESTIGATE WHY NOT PROPOGATE -- TO TEST, REMOVE TIMESTAMP FIELDS
		return err                              ////// ALSO!!! CHANGE TRACKING TABLE, TIMESTAMP FIELD IS UUID MUST CHANGE

	} else {
		log.Info().Str("username", username).Int("id", id).Msg("Username added")
		err = changetracking.Add_to_change_tracking(db_connection, "USERS", id, "ID/USERNAME", "NULL", username)
		if err != nil {
			log.Error().Err(err).Str("location", error_location).Msg("Error passing to change tracking")
		}
	}

	return nil
}

func Update_username_from_id(db_connection *pgxpool.Pool, new_username string, id int, old_username string) (err error) {
	log.Trace().Int("id", id).Str("username", new_username).Msg("SQL: update username")

	// query
	var sql_query string = "UPDATE users SET username = $1 WHERE id = $2;"

	// safeguard: check to make sure username is not just "[Deleted]"
	if new_username == "[Deleted]" {
		err = errors.New("username is '[Deleted]'")
		log.Error().Err(err).Msg("Username is '[Deleted]'")
		return err
	}

	// query results
	_, err = db_connection.Exec(context.Background(), sql_query, new_username, id)
	if err != nil {
		return nil

	} else {
		log.Info().Str("new", new_username).Int("id", id).Msg("Username updated")
		err = changetracking.Add_to_change_tracking(db_connection, "USERS", id, "ID/USERNAME", old_username, new_username)
		if err != nil {
			return err
		}
	}

	return nil
}
