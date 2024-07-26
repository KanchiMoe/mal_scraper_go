package xpaths_users

import (
	"strings"

	"github.com/antchfx/htmlquery"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/html"
)

// When given a url like https://myanimelist.net/comments.php?id=1234567
// this will return the MAL username for that ID
func Username_from_user_comments_page(html_node *html.Node) (username string, err error) {
	log.Trace().Msg("Processing xpath from userpage comments")
	const username_xpath string = "/html/body/div[1]/div[2]/div[3]/div[1]/h1"
	const suffix string = "'s Comments"

	username_node, err := htmlquery.Query(html_node, username_xpath)
	if err != nil {
		log.Error().Err(err).Msg("Error getting username xpath")
		return "", err
	}

	// get username header string
	username_innter_text := htmlquery.InnerText(username_node)

	// remove "'s comments" from the username string
	username = strings.TrimSuffix(username_innter_text, suffix)
	log.Debug().Str("username", username).Msg("Returning username from xpaths")

	return username, nil
}
