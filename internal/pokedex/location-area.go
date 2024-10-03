package pokedex

type LocationArea struct {
	Name              string `json:"name"`
	PokemonEncounters []struct {
		Pokemon Pokemon `json:"pokemon"`
	} `json:"pokemon_encounters"`
}
