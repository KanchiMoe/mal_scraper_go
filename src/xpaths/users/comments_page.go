package xpaths_users

import (
	"fmt"
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

	// check if user has deleted account
	if username_innter_text == "404 Not Found" {
		username = "[Deleted]"
		log.Debug().Str("username", username).Msg("User account has been deleted.")
		return username, nil
	}

	// remove "'s comments" from the username string
	username = strings.TrimSuffix(username_innter_text, suffix)

	// check if username is MALnewbie
	if username == "MALnewbie" {
		const profile_link_xpath string = "html/body/div[1]/div[2]/div[3]/div[2]/div[1]/div[1]/a"

		// find the link to the profile page
		profile_link_node, err := htmlquery.Query(html_node, profile_link_xpath)
		if err != nil {
			log.Error().Err(err).Str("location", "xpath/users").Msg("Error getting profile link xpath")
		}

		// extract the profile link
		profile_link := htmlquery.SelectAttr(profile_link_node, "href")

		// Extract everything after "profile/"
		const profile_prefix string = "/profile/"
		profile_id := strings.TrimPrefix(profile_link, profile_prefix)
		fmt.Println(profile_id)

		// reconstruct username
		username = username + "-" + profile_id
		log.Debug().Str("username", username).Msg("Newbie user account")

		return username, err

	}

	log.Debug().Str("username", username).Msg("Returning username from xpaths")

	return username, nil
}
