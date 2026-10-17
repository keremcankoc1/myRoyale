package getters

import (
	"encoding/json"
	client "myRoyale/Client"
	structs "myRoyale/Structs/DataStruct"
	"strings"
)

func GetCardData(playerTag string) (map[string]int, error) {
	playerTag = strings.ToUpper(playerTag)
	url := "https://api.clashroyale.com/v1/players/%23" + playerTag

	bodyBytes, err := client.Fetch(url)
	if err != nil {
		return map[string]int{}, err
	}

	var player structs.Player
	err = json.Unmarshal(bodyBytes, &player)
	if err != nil {
		return map[string]int{}, err
	}

	cardData := make(map[string]int)
	for _, card := range player.Cards {
		cardData[card.Name] = card.Level + (16 - card.MaxLevel)
	}

	return cardData, nil
}
