package main

import (
	"log"
	handlers "myRoyale/Handlers"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found.")
	}

	handlers.CardHandler()
	handlers.PlayerStatsHandler()

	http.ListenAndServe(":8080", nil)
}
