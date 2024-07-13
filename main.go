package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/antchfx/htmlquery"
	"github.com/kanchimoe/mal_scraper_go/src/xpaths"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	project_config()

	var requested_url string = "https://myanimelist.net/comments.php?id=7354854"

	resp, err := http.Get(requested_url)
	if err != nil {
		panic("stop1")
	}
	defer resp.Body.Close()

	html_node, err := htmlquery.Parse(resp.Body)
	if err != nil {
		panic("html node empty")
	}
	foo, err := xpaths.Xpath_username_from_userpage_comments(html_node)
	if err != nil {
		panic("at the disco")
	}
	fmt.Println(foo)

	panic("stop here main")

}

func project_config() {
	// set log level
	zerolog.SetGlobalLevel(zerolog.TraceLevel)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	zerolog.DurationFieldUnit = time.Second
}
