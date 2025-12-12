package main

import (
	"estebandev_api/api"
	"estebandev_api/db"
	"estebandev_api/webhooks"
	"log"

	"github.com/nxtgo/env"
)

func main() {
	err := env.Load("./.env")

	if err != nil {
		log.Panic(".env file missing")
	}

	webhooks.LoadEvents()
	db.Load()
	api.Load()

}
