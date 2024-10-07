package repl

import (
	"errors"
	"fmt"
)

func (repl *Repl) pokedexCommand(params []string) error {
	caughtPokemons, err := repl.pokedexProvider.CaughtPokemons()
	if err != nil {
		return err
	}
	if len(caughtPokemons) == 0 {
		return errors.New("You haven't caught any pokemon yet")
	}

	fmt.Println("Your Pokedex:")

	for _, pokemon := range caughtPokemons {
		fmt.Printf("  - %s\n", pokemon.Name)
	}

	return nil
}
