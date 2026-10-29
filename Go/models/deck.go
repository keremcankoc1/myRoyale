package models

type Deck struct {
	Cards     []Card  `json:"cards"`
	AvgElixir float64 `json:"avgElixir"`
	AvgLevel  float64 `json:"avgLevel"`
}

func (d Deck) HasCard(cardName string) bool {
	for _, card := range d.Cards {
		if card.Name == cardName {
			return true
		}
	}
	return false
}

func (d Deck) CalculateAvgElixir() float64 {
	var sum float64
	for _, card := range d.Cards {
		sum += float64(card.ElixirCost)
	}
	return sum / float64(8.0)
}

func (d Deck) CalculateAvgLevel() float64 {
	var sum float64
	for _, card := range d.Cards {
		sum += float64(card.ActualLevel())
	}
	return sum / float64(8.0)
}
