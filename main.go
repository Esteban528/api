package main

import (
	// "estebandev_api/api"
	// "estebandev_api/db"
	"log"

	"github.com/nxtgo/env"
)

func main() {
	err := env.Load("./.env")

	if err != nil {
		log.Panic(".env file missing")
	}
	//db.Load()
	//api.Load()

}
