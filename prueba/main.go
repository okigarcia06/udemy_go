package main

import (
	"log"

	"some-time-api/app"
)

func main() {
	log.Println("Starting server at http://localhost:8080")
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
