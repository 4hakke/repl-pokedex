package repl

import "fmt"

func (repl *Repl) exploreCommand(params []string) error {
	if len(params) == 0 {
		fmt.Println("Please enter an area to explore as a parameter")
	}
	area := params[0]
	fmt.Printf("Exploring %s...\n", area)

	locationArea, err := repl.pokedexProvider.GetLocationArea(area)
	if err != nil {
		fmt.Println(err)
	}

	for _, encounters := range locationArea.PokemonEncounters {
		fmt.Printf("- %s\n", encounters.Pokemon.Name)
	}
	return nil
}
