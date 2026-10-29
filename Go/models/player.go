package models

type Player struct {
	Arena                Arena       `json:"arena"`
	Cards                []Card      `json:"cards"`
	Wins                 int         `json:"wins"`
	Losses               int         `json:"losses"`
	KingTowerLevel       int         `json:"kingTowerLevel"`
	BattleCount          int         `json:"battleCount"`
	CollectionLevel      int         `json:"collectionLevel"`
	Trophies             int         `json:"trophies"`
	BestTrophies         int         `json:"bestTrophies"`
	LeagueStatistics     SeasonStats `json:"leagueStatistics"`
	Name                 string      `json:"name"`
	Tag                  string      `json:"tag"`
	ExpLevel             int         `json:"expLevel"`
	ThreeCrownWins       int         `json:"threeCrownWins"`
	ChallengeMaxWins     int         `json:"challengeMaxWins"`
	TotalDonations       int         `json:"totalDonations"`
	Role                 string      `json:"role"`
	ClanData             Clan        `json:"clan"`
	CurrentDeck          []Card      `json:"currentDeck"`
	CurrentFavouriteCard Card        `json:"currentFavouriteCard"`
}
