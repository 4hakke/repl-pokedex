package repl

import (
	"errors"
	"fmt"
)

func (repl *Repl) catchCommand(params []string) error {
	if len(params) == 0 {
		return errors.New("Pokemon name is not provided as an argument")
	}

	if len(params) > 1 {
		return errors.New("Hey, slow down, you can't catch more then one pokemon at once")
	}
	pokemonName := params[0]

	fmt.Printf("Throwing a Pokeball at %s\n", pokemonName)

	isCatch, err := repl.pokedexProvider.Catch(pokemonName)
	if err != nil {
		return err
	}

	if isCatch {
		fmt.Printf("%s was caught!\n", pokemonName)
	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
	}

	return nil
}
