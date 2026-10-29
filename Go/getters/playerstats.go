package getters

import (
	"encoding/json"
	"myRoyale/client"
	"myRoyale/models"
	"strings"
)

func GetPlayerStats(playerTag string) (models.Player, error) {
	playerTag = strings.ToUpper(playerTag)
	url := "https://api.clashroyale.com/v1/players/%23" + playerTag

	bodyBytes, err := client.Fetch(url)
	if err != nil {
		return models.Player{}, err
	}

	var player models.Player
	err = json.Unmarshal(bodyBytes, &player)
	if err != nil {
		return models.Player{}, err
	}

	return player, nil
}
