package structs

type Player struct {
	Arena            Arena       `json:"arena"`
	Cards            []Card      `json:"cards"`
	Wins             int         `json:"wins"`
	Losses           int         `json:"losses"`
	KingTowerLevel   int         `json:"kingTowerLevel"`
	BattleCount      int         `json:"battleCount"`
	CollectionLevel  int         `json:"collectionLevel"`
	Trophies         int         `json:"trophies"`
	BestTrophies     int         `json:"bestTrophies"`
	LeagueStatistics SeasonStats `json:"leagueStatistics"`
}
