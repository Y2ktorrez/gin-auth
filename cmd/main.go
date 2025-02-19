package main

import (
	"log"

	"github.com/mutinho/cmd/api"
	"github.com/mutinho/config"
)

func main() {
	config.Load()
	r := api.SetupRouter()
	log.Println("Server started on port", config.ServerPort)
	r.Run(":" + config.ServerPort)

}
