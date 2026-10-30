package main

import (
	"fmt"
	"log"
	"myRoyale/database"
	"myRoyale/handlers"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found.")
	}

	db, err := database.ConnectDb()
	if err != nil {
		log.Println("Veritabani bağlanti hatasi:", err)
	} else {
		fmt.Println("MSSQL Veritabanina başariyla bağlanildi! 🎉")
		defer db.Close()
	}

	handlers.CardHandler(db)
	handlers.PlayerStatsHandler(db)
	handlers.DeckHandler(db)
	err = database.SeedCardCatalog(db)
	if err != nil {
		fmt.Println("SeedCardCatalog has fault.", err)
	}

	http.ListenAndServe(":8080", nil)
}
