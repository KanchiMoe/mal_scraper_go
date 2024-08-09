package xpaths_genres

import (
	"github.com/kanchimoe/mal_scraper_go/src/project_structs"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/html"
)

const error_location string = "xpath/genre"
const genre_description_xpath string = "/html/body/div[1]/div/div[3]/div[2]/div[4]/p"
const genre_name_suffix string = " Anime"

func Blah(html_node *html.Node) (genre_data project_structs.Genre_xpath, err error) {
	log.Debug().Msg("Processing genres xpaths")

	genre_name, err := get_genre_name(html_node)
	if err != nil {
		log.Error().Err(err).Str("location", error_location).Msg("Error passing to get genre name")
		return genre_data, err
	}
	genre_data.Name = genre_name

	// if genre name is "404", then we want to skip getting the count
	if genre_data.Name == "[404]" {
		genre_data.Count = -1
	} else {
		genre_count, err := get_genre_count(html_node)
		if err != nil {
			log.Error().Err(err).Str("location", error_location).Msg("Error passing to get genre count")
			return genre_data, err
		}
		genre_data.Count = genre_count
	}

	// if genre name is "404", then we want to skip getting the description
	if genre_data.Name == "[404]" {
		genre_data.Description = "[No data]"
	} else {
		genre_description, err := get_genre_description(html_node)
		if err != nil {
			log.Error().Err(err).Str("location", error_location).Msg("Error passing to get genre description")
			return genre_data, err
		}
		genre_data.Description = genre_description
	}

	log.Debug().Str("name", genre_data.Name).Int("count", genre_data.Count).Str("description", genre_data.Description).Msg("Genre xpath data")
	return genre_data, nil
}
