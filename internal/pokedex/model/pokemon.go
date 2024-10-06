package model

type Pokemon struct {
	Name           string        `json:"name"`
	BaseExperience int           `json:"base_experience"`
	Weight         int           `json:"weight"`
	Height         int           `json:"height"`
	Stats          []PokemonStat `json:"stats"`
}

type PokemonStat struct {
	BaseStat int `json:"base_stat"`
	Stat     struct {
		Name string `json:"name"`
	} `json:"stat"`
}
