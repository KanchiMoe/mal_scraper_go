package logic_users

import (
	"fmt"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kanchimoe/mal_scraper_go/src"
	"github.com/kanchimoe/mal_scraper_go/src/database"
	"github.com/kanchimoe/mal_scraper_go/src/page_requestor"
	"github.com/kanchimoe/mal_scraper_go/src/project_structs"
	sql_users "github.com/kanchimoe/mal_scraper_go/src/sql/users"
	xpaths_users "github.com/kanchimoe/mal_scraper_go/src/xpaths/users"
	"github.com/rs/zerolog/log"
)

// we take id
// grab the username from mal
// check if id is in db (we get id and username)

func Count_up_1000(start_id int) (err error) {
	log.Debug().Int("starting id", start_id).Msg("User count up 1000 from ID")

	// db connection
	db_connection := database.Database_pool()

	// init struct
	var from_mal project_structs.From_mal

	// calculate upper values
	var id_1000_up int = start_id + 1000
	var current_id int = start_id + 1

	for current_id <= id_1000_up {
		// assign current id to comparison struct
		from_mal.Id = current_id

		log.Trace().Msg("......")
		log.Trace().Msg("......")
		log.Trace().Msg("......")
		log.Trace().Msg("......")

		// create url
		var url string = src.ROOT_URL + src.COMMENTS_SLUG + strconv.Itoa(current_id)
		log.Trace().Str("url", url).Msg("Constructed url")

		// request url
		html, err := page_requestor.Request_handler(url)
		if err != nil {
			return err
		}

		// get username from the page
		xpath_username, err := xpaths_users.Username_from_user_comments_page(html)
		if err != nil {
			return err
		}

		//put username in struct
		from_mal.Username = xpath_username

		// check if username is in db

		checks(db_connection, from_mal)

		current_id++
		time.Sleep(1 * time.Second)

	}

	return nil
}

func checks(db_connection *pgxpool.Pool, from_mal project_structs.From_mal) (err error) {

	// Check to see if username is just [Deleted]
	if from_mal.Username == "[Deleted]" {
		from_mal.Username = from_mal.Username + "-" + strconv.Itoa(from_mal.Id)
		log.Warn().Str("amended username", from_mal.Username).Msg("Username is just '[Deleted]', appending id as suffix")
	}

	log.Debug().Str("username", from_mal.Username).Int("id", from_mal.Id).Msg("Conducting username checks with data from MAL")

	// check if username is in the db already
	err = check_if_username_is_in_db(db_connection, from_mal)
	if err != nil {
		return err
	}
	err = check_if_id_is_in_db(db_connection, from_mal)
	if err != nil {
		return err
	}

	return err
}

func check_if_username_is_in_db(db_connection *pgxpool.Pool, from_mal project_structs.From_mal) (err error) {
	log.Trace().Str("username", from_mal.Username).Msg("Checking if username is in database")
	id_from_username_data := sql_users.Get_id_from_username(db_connection, from_mal.Username)

	if !id_from_username_data.In_db {
		log.Info().Str("username", from_mal.Username).Msg("Username is NOT in database")

	} else if id_from_username_data.In_db {
		log.Info().Str("username", from_mal.Username).Msg("Username already exists in database")
		log.Debug().Msg("Checking to see if the ID from MAL matches the database")

		if id_from_username_data.Id == from_mal.Id {
			log.Info().Int("mal", from_mal.Id).Int("db", id_from_username_data.Id).Msg("IDs match. No action needed.")

		} else if id_from_username_data.Id != from_mal.Id {
			log.Warn().Int("db", id_from_username_data.Id).Int("mal", from_mal.Id).Msg("The ID from DB does not match MAL")
			err := ids_dont_match_update_them(db_connection, id_from_username_data.Id)
			if err != nil {
				return err
			}
		}
	}
	return err
}

func check_if_id_is_in_db(db_connection *pgxpool.Pool, from_mal project_structs.From_mal) (err error) {
	log.Debug().Int("id", from_mal.Id).Msg("Checking if ID is in database")
	username_from_id_data := sql_users.Get_username_from_id(db_connection, from_mal.Id)

	// .In_db = false = the id is not in the db
	// .In_db = true = the id is not in the db

	if !username_from_id_data.In_db {
		log.Info().Int("id", from_mal.Id).Msg("ID NOT in database")
		err = sql_users.Add_user_id_to_db(db_connection, from_mal.Id, from_mal.Username)
		if err != nil {
			return err
		}

	} else if username_from_id_data.In_db {
		log.Info().Int("id", username_from_id_data.Id).Msg("ID is already in the database")
		log.Debug().Msg("Checking to see if the usernames from MAL matches the database")

		if username_from_id_data.Username == from_mal.Username {
			log.Info().Str("mal", from_mal.Username).Str("db", username_from_id_data.Username).Msg("Usernames match. No action needed.")

		} else if username_from_id_data.Username != from_mal.Username {
			fmt.Println("USERNAME DOES NOT MATCH DB. update needed")
			log.Warn().Str("mal", from_mal.Username).Str("db", username_from_id_data.Username).Msg("Username from DB does not match MAL")
			sql_users.Update_username_from_id(db_connection, from_mal.Username, username_from_id_data.Id)
		}

	}
	return err
}

func ids_dont_match_update_them(db_connection *pgxpool.Pool, id int) (err error) {
	log.Info().Int("id", id).Msg("Getting data to update username for")

	// construct url
	var url string = src.ROOT_URL + src.COMMENTS_SLUG + strconv.Itoa(id)
	log.Trace().Str("url", url).Msg("Constructed url")

	// request url
	html, err := page_requestor.Request_handler(url)
	if err != nil {
		log.Error().Err(err).Msg("Error trying to request url")
		return err
	}

	// get username from page
	xpaths_username, err := xpaths_users.Username_from_user_comments_page(html)
	if err != nil {
		return err
	}

	err = sql_users.Update_username_from_id(db_connection, xpaths_username, id)
	if err != nil {
		return err
	}

	return nil
}
