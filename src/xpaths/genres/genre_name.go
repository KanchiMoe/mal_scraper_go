package xpaths_genres

import (
	"strings"

	"github.com/antchfx/htmlquery"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/html"
)

const genre_name_xpath string = "/html/body/div[1]/div/div[3]/div[1]/h1"

func get_genre_name(html_node *html.Node) (genre_name string, err error) {
	genre_name_node, err := htmlquery.Query(html_node, genre_name_xpath)
	if err != nil {
		log.Error().Err(err).Str("location", error_location).Msg("Error getting genre name xpath")
		return "", err
	}

	log.Trace().Msg("Getting genre header string")

	// get genre header string
	genre_innter_text := htmlquery.InnerText(genre_name_node)

	// check if it has been delete or does not exist
	if genre_innter_text == "404 Not Found" {
		genre_name := "[404]"
		log.Debug().Str("setting name", genre_name).Msg("Genre does not exist or deleted.")
		return genre_name, nil
	}

	// remove "anime" from the name string
	genre_name = strings.TrimSuffix(genre_innter_text, genre_name_suffix)

	log.Trace().Str("name", genre_name).Msg("Got genre name")
	return genre_name, nil
}
