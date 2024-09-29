package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/4hakke/repl-pokedex/internal/pokeapi"
)

type cliCommand struct {
	action      func(params []string) error
	name        string
	description string
}

var (
	commands          map[string]cliCommand
	locationsIterator pokeapi.LocationsIterator
)

func main() {
	commands = map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			action:      helpCommand,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			action:      exitCommand,
		},
		"map": {
			name:        "map",
			description: "Iterate through list of locations",
			action:      mapCommand,
		},
		"mapb": {
			name:        "mapb",
			description: "Iterate through list of locations (backwards)",
			action:      mapBCommand,
		},
		"explore": {
			name:        "explore",
			description: "Explore the pokemons on an area",
			action:      exploreCommand,
		},
	}
	locationsIterator = pokeapi.LocationsIterator{Limit: 20}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanSuccess := scanner.Scan()

		if !scanSuccess {
			return
		}
		enteredCommand := strings.Split(scanner.Text(), " ")
		command := commands[enteredCommand[0]]

		if command.name != enteredCommand[0] {
			continue
		}
		if len(enteredCommand) > 1 {
			command.action(enteredCommand[1:])
		} else {
			command.action(make([]string, 0))
		}
	}
}

func helpCommand(params []string) error {
	fmt.Println("")
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("")
	fmt.Println("Usage:")
	fmt.Printf("%s: %s\n", commands["help"].name, commands["help"].description)
	fmt.Printf("%s: %s\n", commands["exit"].name, commands["exit"].description)
	fmt.Printf("%s: %s\n", commands["map"].name, commands["map"].description)
	fmt.Printf("%s: %s\n", commands["mapb"].name, commands["mapb"].description)
	fmt.Printf("%s: %s\n", commands["explore"].name, commands["explore"].description)
	fmt.Println("")
	return nil
}

func exitCommand(params []string) error {
	os.Exit(0)
	return nil
}

func mapCommand(params []string) error {
	locations, err := locationsIterator.Next()
	if err != nil {
		fmt.Println(err)
		return err
	}
	for _, location := range locations {
		fmt.Println(location.Name)
	}
	return nil
}

func mapBCommand(params []string) error {
	locations, err := locationsIterator.Previous()
	if err != nil {
		fmt.Println(err)
		return err
	}
	for _, location := range locations {
		fmt.Println(location.Name)
	}
	return nil
}

func exploreCommand(params []string) error {
	if len(params) == 0 {
		fmt.Println("Please enter an area to explore as a parameter")
	}
	area := params[0]
	fmt.Printf("Exploring %s...\n", area)

	locationArea, err := pokeapi.GetLocationArea(area)
	if err != nil {
		fmt.Println(err)
	}

	for _, encounters := range locationArea.PokemonEncounters {
		fmt.Printf("- %s\n", encounters.Pokemon.Name)
	}
	return nil
}
