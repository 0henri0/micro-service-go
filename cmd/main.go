package main

import (
	"log"
	"micro-service-go/cmd/api/server"
)

func main() {
	app, err := server.NewApp()

	if err != nil {
		log.Println(err)
	}

	if err := app.Run(); err != nil {
		log.Println(err)
	}
}
