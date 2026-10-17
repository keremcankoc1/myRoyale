package structs

type Card struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Level    int    `json:"level"`
	MaxLevel int    `json:"maxLevel"`
}
