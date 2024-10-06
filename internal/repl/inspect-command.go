package repl

import (
	"errors"
	"fmt"
)

func (repl *Repl) inspectCommand(params []string) error {
	if len(params) == 0 {
		return errors.New("Pokemon name is not provided as an argument")
	}
	pokemonName := params[0]
	pokemon, err := repl.pokedexProvider.Inspect(pokemonName)
	if err != nil {
		return err
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  - %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("  - %s\n", t.Type.Name)
	}

	return nil
}
