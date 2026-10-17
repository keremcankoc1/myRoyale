package getters

import (
	"encoding/json"
	client "myRoyale/Client"
	structs "myRoyale/Structs/DataStruct"
	"strings"
)

func GetPlayerStats(playerTag string) (structs.Player, error) {
	playerTag = strings.ToUpper(playerTag)
	url := "https://api.clashroyale.com/v1/players/%23" + playerTag

	bodyBytes, err := client.Fetch(url)
	if err != nil {
		return structs.Player{}, err
	}

	var player structs.Player
	err = json.Unmarshal(bodyBytes, &player)
	if err != nil {
		return structs.Player{}, err
	}

	return player, nil
}
