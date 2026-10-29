package database

import (
	"database/sql"
	"myRoyale/models"
)

func SavePlayer(db *sql.DB, player models.Player, tag string) error {
	query := `IF EXISTS (SELECT 1 FROM players WHERE player_tag = @p1)
			BEGIN
				UPDATE players 
				SET name = @p2, trophies = @p3, best_trophies = @p4, 
					king_tower_level = @p5, wins = @p6, losses = @p7, 
					battle_count = @p8, exp_level = @p9, three_crown_wins = @p10,
					challenge_max_wins = @p11, total_donations = @p12, role = @p13,
					arena_name = @p14, clan_tag = @p15, clan_name = @p16,
					favourite_card = @p17, last_updated = GETDATE()
				WHERE player_tag = @p1
			END
			ELSE
			BEGIN
				INSERT INTO players (
					player_tag, name, trophies, best_trophies, king_tower_level, wins, losses,
					battle_count, exp_level, three_crown_wins, challenge_max_wins, total_donations,
					role, arena_name, clan_tag, clan_name, favourite_card
				)
				VALUES (
					@p1, @p2, @p3, @p4, @p5, @p6, @p7,
					@p8, @p9, @p10, @p11, @p12, @p13, @p14, @p15, @p16, @p17
				);
			END`

	_, err := db.Exec(
		query,
		tag,                              // @p1
		player.Name,                      // @p2
		player.Trophies,                  // @p3
		player.BestTrophies,              // @p4
		player.KingTowerLevel,            // @p5
		player.Wins,                      // @p6
		player.Losses,                    // @p7
		player.BattleCount,               // @p8
		player.ExpLevel,                  // @p9
		player.ThreeCrownWins,            // @p10
		player.ChallengeMaxWins,          // @p11
		player.TotalDonations,            // @p12
		player.Role,                      // @p13
		player.Arena.Name,                // @p14
		player.ClanData.Tag,              // @p15
		player.ClanData.Name,             // @p16
		player.CurrentFavouriteCard.Name, // @p17
	)

	if err != nil {
		return err
	}

	return nil
}

func IsPlayerFresh(db *sql.DB, tag string) (bool, error) {
	query := `SELECT COUNT(1)
			  FROM players
			  WHERE player_tag = @p1 
				AND last_updated > DATEADD(minute, -5, GETDATE())
			`
	var count int
	err := db.QueryRow(query, tag).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func GetPlayerFromDB(db *sql.DB, tag string) (models.Player, error) {
	query := `SELECT 
				player_tag, name, trophies, best_trophies, king_tower_level, wins, losses,
				battle_count, exp_level, three_crown_wins, challenge_max_wins,
				total_donations, role, arena_name, clan_tag, clan_name, favourite_card
			  FROM players
			  WHERE player_tag = @p1
			 `
	var player models.Player

	err := db.QueryRow(query, tag).Scan(
		&player.Tag,
		&player.Name,
		&player.Trophies,
		&player.BestTrophies,
		&player.KingTowerLevel,
		&player.Wins,
		&player.Losses,
		&player.BattleCount,
		&player.ExpLevel,
		&player.ThreeCrownWins,
		&player.ChallengeMaxWins,
		&player.TotalDonations,
		&player.Role,
		&player.Arena.Name,
		&player.ClanData.Tag,
		&player.ClanData.Name,
		&player.CurrentFavouriteCard.Name,
	)
	if err != nil {
		return models.Player{}, err
	}

	return player, nil
}
