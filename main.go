package main

import (
	"fmt"

	logic_users "github.com/kanchimoe/mal_scraper_go/src/logic/users"
	"github.com/kanchimoe/mal_scraper_go/src/project_init"
	"github.com/rs/zerolog/log"
)

func main() {
	err := project_init.Init_main()
	if err != nil {
		log.Panic().Err(err).Msg("Error initialising project")
	}

	logic_users.Count_up_1000(7354862)

	fmt.Println("end")
}
