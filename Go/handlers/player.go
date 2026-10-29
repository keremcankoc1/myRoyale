// Package handlers, HTTP isteklerini (Request) karşılayıp istemciye yanıt (Response) dönen katmandır.
package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"myRoyale/database"
	"myRoyale/getters"
	"myRoyale/models"
	"net/http"
	"strings"
)

// PlayerHandler struct'ı: Dependency Injection (Bağımlılık Enjeksiyonu) amacıyla kullanılır.
// Veritabanı bağlantısını (DB) içinde tutar; böylece handler metodları h.DB üzerinden veritabanına erişebilir.
type PlayerHandler struct {
	DB *sql.DB // Canlı veritabanı hattı
}

func (h *PlayerHandler) GetCards(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getCards has worked.")

	tag := r.PathValue("tag")
	tag = strings.ToUpper(tag)
	if len(tag) == 0 {
		fmt.Println("Tag not found.")
		http.Error(w, "Tag not found.", http.StatusBadRequest) // 400 Bad Request
		return
	}

	err := h.SyncPlayer(tag)
	if err != nil {
		fmt.Println(err)
		return
	}

	cards, err := database.GetCardsFromDB(h.DB, tag)
	if err != nil {
		fmt.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cards)
}

// CardHandler: "/player/{tag}/cards" rotasını GetCards fonksiyonuna kaydeder.
func CardHandler(db *sql.DB) {
	handler := &PlayerHandler{DB: db}
	http.HandleFunc("/player/{tag}/cards", handler.GetCards)
	fmt.Println("Server running at http://localhost:8080")
}

func (h *PlayerHandler) GetPlayerStats(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetPlayerStats has worked")

	tag := r.PathValue("tag")
	tag = strings.ToUpper(tag)
	if len(tag) == 0 {
		fmt.Println("Tag not found.")
		http.Error(w, "Tag not found.", http.StatusBadRequest) // 400 Bad Request
		return
	}

	err := h.SyncPlayer(tag)
	if err != nil {
		fmt.Println(err)
		return
	}

	stats, err := database.GetPlayerFromDB(h.DB, tag)
	if err != nil {
		fmt.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stats)

}

// PlayerStatsHandler: Oyuncu istatistikleri endpoint'ini yönlendiriciye (router) bağlar.
func PlayerStatsHandler(db *sql.DB) {
	handler := &PlayerHandler{DB: db}
	http.HandleFunc("/player/{tag}/playerStats", handler.GetPlayerStats)
	fmt.Println("Server running at http://localhost:8080")
}

func (h *PlayerHandler) SyncPlayer(tag string) error {
	fresh, err := database.IsPlayerFresh(h.DB, tag)
	if err != nil {
		return err
	}

	if fresh {
		fmt.Println("Player data is fresh, reading from DB ⚡")
	} else {
		fmt.Println("Data is stale or missing, fetching from Supercell API 🌐")
		stats, err := getters.GetPlayerStats(tag)
		if err != nil {
			return err
		}

		err = database.SavePlayer(h.DB, stats, tag)
		if err != nil {
			return err
		}

		err = database.SaveCards(h.DB, stats.Cards, tag, models.Deck{Cards: stats.CurrentDeck})
		if err != nil {
			return err
		}
	}

	return nil
}

func (h *PlayerHandler) GetDeck(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetDeck has worked.")

	tag := r.PathValue("tag")
	tag = strings.ToUpper(tag)
	if len(tag) == 0 {
		fmt.Println("Tag not found.")
		http.Error(w, "Tag not found.", http.StatusBadRequest) // 400 Bad Request
		return
	}

	err := h.SyncPlayer(tag)
	if err != nil {
		fmt.Println(err)
		return
	}

	deck, err := database.GetDeckFromDB(h.DB, tag)
	if err != nil {
		fmt.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(deck)
}

func DeckHandler(db *sql.DB) {
	handler := &PlayerHandler{DB: db}
	http.HandleFunc("/player/{tag}/deck", handler.GetDeck)
	fmt.Println("Server running at http://localhost:8080")
}
