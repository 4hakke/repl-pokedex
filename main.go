package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	action      func() error
	name        string
	description string
}

var commands map[string]cliCommand

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
	}

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
	fmt.Println("Usage:")
	fmt.Printf("%s: %s\n", commands["help"].name, commands["help"].description)
	fmt.Printf("%s: %s\n", commands["exit"].name, commands["exit"].description)
	fmt.Println("")
	return nil
}

func exitCommand() error {
	os.Exit(0)
	return nil
}
