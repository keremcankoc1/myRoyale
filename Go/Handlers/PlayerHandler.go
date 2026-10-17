package handlers

import (
	"encoding/json"
	"fmt"
	errorstruct "myRoyale/Errors/ErrorStruct"
	getters "myRoyale/Getters"
	"net/http"
)

// / -------------- Card Process --------------
func GetCards(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getCards has worked.")
	tag := r.PathValue("tag")
	if len(tag) == 0 {
		fmt.Println("Tag not found.")
		http.Error(w, "Tag not found.", http.StatusBadRequest)
		return
	}

	cards, err := getters.GetCardData(tag)
	if err != nil {
		fmt.Println(err)
		if statErr, ok := err.(*errorstruct.StatCode); ok {
			http.Error(w, statErr.Error(), statErr.Code)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cards)
}

func CardHandler() {
	http.HandleFunc("/player/{tag}/cards", GetCards)
	fmt.Println("Server running at http://localhost:8080")
}

/// -------------- Player Stats --------------

func GetPlayerStats(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetPlayerStats has worked")
	tag := r.PathValue("tag")
	if len(tag) == 0 {
		fmt.Println("Tag not found.")
		http.Error(w, "Tag not found.", http.StatusBadRequest)
		return
	}

	stats, err := getters.GetPlayerStats(tag)
	if err != nil {
		fmt.Println(err)
		if statErr, ok := err.(*errorstruct.StatCode); ok {
			http.Error(w, statErr.Error(), statErr.Code)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stats)
}

func PlayerStatsHandler() {
	http.HandleFunc("/player/{tag}/playerStats", GetPlayerStats)
	fmt.Println("Server running at http://localhost:8080")
}
