package main

import (
	"fmt"

	"github.com/kanchimoe/mal_scraper_go/src/project_init"
	"github.com/rs/zerolog/log"
)

func main() {
	project_config()
	err := project_init.Init_main()
	if err != nil {
		log.Panic().Err(err).Msg("Error initialising project")
	}

	fmt.Println("dssfg")

}

func project_config() {

}
