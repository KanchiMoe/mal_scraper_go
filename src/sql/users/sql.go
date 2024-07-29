package sql_users

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kanchimoe/mal_scraper_go/src/project_structs"
	"github.com/rs/zerolog/log"
)

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

	// query
	const sql_insert_new_user_id string = "INSERT INTO users (id, username) VALUES ($1, $2)"

	// query results
	_, err = db_connection.Exec(context.Background(), sql_insert_new_user_id, id, username)
	if err != nil {
		return err

	} else {
		log.Info().Str("username", username).Int("id", id).Msg("Username added")
	}

	return nil
}

func Update_username_from_id(db_connection *pgxpool.Pool, username string, id int) (err error) {
	log.Trace().Int("id", id).Str("username", username).Msg("SQL: update username")

	// query
	var sql_query string = "UPDATE users SET username = $1 WHERE id = $2;"

	// safeguard: check to make sure username is not just "[Deleted]"
	if username == "[Deleted]" {
		err = errors.New("username is '[Deleted]'")
		log.Error().Err(err).Msg("Username is '[Deleted]'")
		return err
	}

	// query results
	_, err = db_connection.Exec(context.Background(), sql_query, username, id)
	if err != nil {
		return nil

	} else {
		log.Info().Str("new username", username).Int("id", id).Msg("Username updated")
	}

	return nil
}
