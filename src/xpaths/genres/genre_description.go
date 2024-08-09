package xpaths_genres

import (
	"github.com/antchfx/htmlquery"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/html"
)

func get_genre_description(html_node *html.Node) (genre_description string, err error) {
	log.Trace().Msg("Getting genre description")
	genre_description_node, err := htmlquery.Query(html_node, genre_description_xpath)
	if err != nil {
		log.Error().Err(err).Str("location", error_location).Msg("Error getting genre description xpath")
	}

	// if the description block is not on the page
	if genre_description_node == nil {
		log.Warn().Msg("Genre does not have a description.")
		genre_description = "[No Description]"

	} else {
		// process if it does have a description block
		genre_description_inner_text := htmlquery.InnerText(genre_description_node)
		log.Trace().Str("description", genre_description_inner_text).Msg("Got genre description")
		return genre_description_inner_text, nil

	}

	return "", nil
}
