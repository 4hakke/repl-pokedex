package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/4hakke/repl-pokedex/internal/pokeapi"
)

type cliCommand struct {
	action      func() error
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
	}
	locationsIterator = pokeapi.LocationsIterator{Limit: 20}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanSuccess := scanner.Scan()

		if !scanSuccess {
			return
		}
		enteredCommand := scanner.Text()
		command := commands[enteredCommand]

		if command.name != enteredCommand {
			continue
		}

		command.action()
	}
}

func helpCommand() error {
	fmt.Println("")
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("")
	fmt.Println("Usage:")
	fmt.Printf("%s: %s\n", commands["help"].name, commands["help"].description)
	fmt.Printf("%s: %s\n", commands["exit"].name, commands["exit"].description)
	fmt.Printf("%s: %s\n", commands["map"].name, commands["map"].description)
	fmt.Printf("%s: %s\n", commands["mapb"].name, commands["mapb"].description)
	fmt.Println("")
	return nil
}

func exitCommand() error {
	os.Exit(0)
	return nil
}

func mapCommand() error {
	locations, err := locationsIterator.Next()
	if err != nil {
		return err
	}
	for _, location := range locations {
		fmt.Println(location.Name)
	}
	return nil
}

func mapBCommand() error {
	locations, err := locationsIterator.Previous()
	if err != nil {
		return err
	}
	for _, location := range locations {
		fmt.Println(location.Name)
	}
	return nil
}
