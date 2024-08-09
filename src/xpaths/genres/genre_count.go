package xpaths_genres

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/antchfx/htmlquery"
	"github.com/kanchimoe/mal_scraper_go/src"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/html"
)

const genre_count_xpath string = "/html/body/div[1]/div/div[3]/div[2]/div[3]/span/span"

func get_genre_count(html_node *html.Node) (count int, err error) {
	log.Trace().Msg("Getting genre count")
	genre_count_node, err := htmlquery.Query(html_node, genre_count_xpath)
	if err != nil {
		log.Error().Err(err).Str("location", error_location).Msg("Error getting genre count xpath")
		return -1, err
	}

	// get count string
	genre_count_innter_text := htmlquery.InnerText(genre_count_node)

	// remove any commas from the string
	genre_count_innter_text = strings.ReplaceAll(genre_count_innter_text, ",", "")

	// remove brackets from string
	re := regexp.MustCompile(src.REGEX_INT_IN_BRACKETS)
	extracted_number := re.FindStringSubmatch(genre_count_innter_text)

	// cast to int
	extracted_int, err := strconv.Atoi(extracted_number[1])
	if err != nil {
		log.Error().Err(err).Str("location", error_location).Msg("Error converting string to int")
		return -1, err
	}

	log.Trace().Int("count", extracted_int).Msg("Got genre count")
	return extracted_int, err
}
