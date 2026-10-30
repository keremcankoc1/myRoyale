package database

import (
	"database/sql"
	"encoding/json"
	"myRoyale/models"
	"os"
)

func SeedCardCatalog(db *sql.DB) error {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM card_catalog").Scan(&count)
	if err != nil {
		return err
	}

	if count > 0 {
		return nil
	}

	byteValue, err := os.ReadFile("database/card_catalog.json")
	if err != nil {
		return err
	}

	var catalog []models.Card
	err = json.Unmarshal(byteValue, &catalog)
	if err != nil {
		return err
	}

	query := `INSERT INTO card_catalog
				(name, card_type, role, is_flying, targets_air, targets_only_builds, is_splash)
			  VALUES(@p1, @p2, @p3, @p4, @p5, @p6, @p7)
			 `

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, card := range catalog {
		_, err := stmt.Exec(card.Name, card.CardType, card.Role, card.IsFlying, card.TargetsAir, card.TargetsOnlyBuilds, card.IsSplash)
		if err != nil {
			return err
		}
	}

	return nil
}
