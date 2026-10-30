package database

import (
	"database/sql"
	"myRoyale/models"
)

func SaveCards(db *sql.DB, cards []models.Card, tag string, currentDeck models.Deck) error {
	queryDelete := `DELETE FROM player_cards WHERE player_tag = @p1`
	_, err := db.Exec(queryDelete, tag)
	if err != nil {
		return err
	}

	queryInsert := `INSERT INTO player_cards
						(player_tag, card_name, real_level, max_level, elixir_cost, rarity, evolution_level, max_evolution_level, is_in_deck)
					VALUES (@p1, @p2, @p3, @p4, @p5, @p6, @p7, @p8, @p9)`
	stmt, err := db.Prepare(queryInsert)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, card := range cards {
		_, err := stmt.Exec(tag, card.Name, card.ActualLevel(), card.MaxLevel, card.ElixirCost, card.Rarity, card.EvolutionLevel, card.MaxEvolutionLevel, currentDeck.HasCard(card.Name))
		if err != nil {
			return err
		}
	}

	return nil
}

func GetCardsFromDB(db *sql.DB, tag string) ([]models.Card, error) {
	query := `SELECT pc.card_name, pc.real_level, pc.elixir_cost, pc.rarity, pc.evolution_level, pc.max_evolution_level, pc.is_in_deck,
				cc.card_type, cc.role, cc.is_flying, cc.targets_air, cc.targets_only_builds, cc.is_splash
			  FROM player_cards pc
			  INNER JOIN card_catalog cc ON pc.card_name = cc.name
			  WHERE pc.player_tag = @p1
			 `
	rows, err := db.Query(query, tag)
	if err != nil {
		return []models.Card{}, err
	}
	defer rows.Close()

	var cards []models.Card

	for rows.Next() {
		var name string
		var level int
		var elixir_cost int
		var rarity string
		var evolution_level int
		var max_evolution_level int
		var is_in_deck bool
		var card_type string
		var role string
		var is_flying bool
		var targets_air bool
		var targets_only_build bool
		var is_splash bool
		rows.Scan(&name, &level, &elixir_cost, &rarity, &evolution_level, &max_evolution_level, &is_in_deck,
			&card_type, &role, &is_flying, &targets_air, &targets_only_build, &is_splash)
		card := models.Card{
			Name:              name,
			Level:             level,
			ElixirCost:        elixir_cost,
			Rarity:            rarity,
			EvolutionLevel:    evolution_level,
			MaxEvolutionLevel: max_evolution_level,
			IsInDeck:          is_in_deck,
			CardType:          card_type,
			Role:              role,
			IsFlying:          is_flying,
			TargetsAir:        targets_air,
			TargetsOnlyBuilds: targets_only_build,
			IsSplash:          is_splash,
		}
		cards = append(cards, card)
	}

	return cards, nil
}

func GetDeckFromDB(db *sql.DB, tag string) (models.Deck, error) {
	query := `SELECT pc.card_name, pc.real_level, pc.elixir_cost, pc.rarity, pc.evolution_level, pc.max_evolution_level, pc.is_in_deck,
				cc.card_type, cc.role, cc.is_flying, cc.targets_air, cc.targets_only_builds, cc.is_splash
			  FROM player_cards pc
			  INNER JOIN card_catalog cc ON pc.card_name = cc.name
			  WHERE pc.player_tag = @p1 AND is_in_deck = 1
			 `
	rows, err := db.Query(query, tag)
	if err != nil {
		return models.Deck{}, err
	}
	defer rows.Close()

	var cards []models.Card
	var deck models.Deck

	for rows.Next() {
		var name string
		var real_level int
		var elixir_cost int
		var rarity string
		var evolution_level int
		var max_evolution_level int
		var is_in_deck bool
		var card_type string
		var role string
		var is_flying bool
		var targets_air bool
		var targets_only_build bool
		var is_splash bool
		rows.Scan(&name, &real_level, &elixir_cost, &rarity, &evolution_level, &max_evolution_level, &is_in_deck,
			&card_type, &role, &is_flying, &targets_air, &targets_only_build, &is_splash)
		card := models.Card{
			Name:              name,
			Level:             real_level,
			ElixirCost:        elixir_cost,
			Rarity:            rarity,
			EvolutionLevel:    evolution_level,
			MaxEvolutionLevel: max_evolution_level,
			IsInDeck:          is_in_deck,
			CardType:          card_type,
			Role:              role,
			IsFlying:          is_flying,
			TargetsAir:        targets_air,
			TargetsOnlyBuilds: targets_only_build,
			IsSplash:          is_splash,
		}
		cards = append(cards, card)
		deck = models.Deck{
			Cards: cards,
		}
	}
	deck.AvgElixir = deck.CalculateAvgElixir()
	deck.AvgLevel = deck.CalculateAvgLevel()

	return deck, nil
}
