package structs

type SeasonStats struct {
	CurrentSeason  LeagueStat `json:"currentSeason"`
	PreviousSeason LeagueStat `json:"previousSeason"`
	BestSeason     LeagueStat `json:"bestSeason"`
}

type LeagueStat struct {
	Trophies     int    `json:"trophies"`
	Id           string `json:"id"`
	Rank         int    `json:"rank"`
	BestTrophies int    `json:"bestTrophies"`
}
