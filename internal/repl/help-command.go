package repl

import "fmt"

func (repl *Repl) helpCommand(params []string) error {
	fmt.Println("")
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("")
	fmt.Println("Usage:")
	fmt.Printf("%s: %s\n", commands["help"].name, commands["help"].description)
	fmt.Printf("%s: %s\n", commands["exit"].name, commands["exit"].description)
	fmt.Printf("%s: %s\n", commands["map"].name, commands["map"].description)
	fmt.Printf("%s: %s\n", commands["mapb"].name, commands["mapb"].description)
	fmt.Printf("%s: %s\n", commands["explore"].name, commands["explore"].description)
	fmt.Printf("%s: %s\n", commands["inspect"].name, commands["inspect"].description)
	fmt.Printf("%s: %s\n", commands["pokedex"].name, commands["pokedex"].description)
	fmt.Println("")
	return nil
}
