package models

type Card struct {
	Id                int    `json:"id"`
	Name              string `json:"name"`
	Level             int    `json:"level"`
	MaxLevel          int    `json:"maxLevel"`
	ElixirCost        int    `json:"elixirCost"`
	Rarity            string `json:"rarity"`
	EvolutionLevel    int    `json:"evolutionLevel"`    // Oyuncuda açık mı
	MaxEvolutionLevel int    `json:"maxEvolutionLevel"` // Bu kartın evrimi oyunda var mı
	Count             int    `json:"count"`             // Oyuncunun elinde biriken kart adedi (Yükseltme maliyeti için)
	IsInDeck          bool   `json:"isInDeck"`
	CardType          string `json:"card_type"`
	Role              string `json:"role"`
	IsFlying          bool   `json:"is_flying"`
	TargetsAir        bool   `json:"targets_air"`
	TargetsOnlyBuilds bool   `json:"targets_only_builds"`
	IsSplash          bool   `json:"is_splash"`
}

func (c Card) ActualLevel() int {
	if c.MaxLevel == 0 {
		return c.Level // eğer dbden gelmişse
	}
	return c.Level + (16 - c.MaxLevel)
}
