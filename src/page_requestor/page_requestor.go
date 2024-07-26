package page_requestor

import (
	"github.com/antchfx/htmlquery"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/html"
)

func Request_handler(requested_url string) (html *html.Node, err error) {
	log.Debug().Str("url", requested_url).Msg("Requesting url")
	_, html, err = get_url(requested_url)
	if err != nil {
		return nil, err
	}

	return html, err
}

func get_url(requested_url string) (status_code int, html *html.Node, err error) {
	// initialise the client
	mal_client := mal_client()

	// request the url
	response, err := mal_client.Get(requested_url)
	if err != nil {
		log.Error().Err(err).Msg("a could not request url")
		return 0, nil, err
	}
	defer response.Body.Close()

	log.Trace().Int("code", response.StatusCode).Msg("Response code")

	// actions based on response code
	if response.StatusCode == 200 || response.StatusCode == 404 {
		html, err = htmlquery.Parse(response.Body)
		if err != nil {
			log.Error().Err(err).Msg("Error when parsing response body")
		}
	}

	return response.StatusCode, html, nil

}
