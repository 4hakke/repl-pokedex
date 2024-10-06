package model

type LocationArea struct {
	Name              string `json:"name"`
	PokemonEncounters []struct {
		Pokemon PokemonEncounter `json:"pokemon"`
	} `json:"pokemon_encounters"`
}
