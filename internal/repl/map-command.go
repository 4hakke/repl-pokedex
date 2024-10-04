package repl

import "fmt"

func (repl *Repl) mapCommand(params []string) error {
	locations, err := repl.pokedexProvider.Locations(0, 20)
	if err != nil {
		fmt.Println(err)
		return err
	}
	for _, location := range locations {
		fmt.Println(location.Name)
	}
	return nil
}

func (repl *Repl) mapBCommand(params []string) error {
	locations, err := repl.pokedexProvider.Locations(0, 20)
	if err != nil {
		fmt.Println(err)
		return err
	}
	for _, location := range locations {
		fmt.Println(location.Name)
	}
	return nil
}
