package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	action func() error
	name   string
}

func main() {
	commands := map[string]cliCommand{
		"help": {
			name:   "help",
			action: helpCommand,
		},
		"exit": {
			name:   "exit",
			action: exitCommand,
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

		fmt.Println(enteredCommand)
		fmt.Println(command.name)

		if command.name != enteredCommand {
			continue
		}

		command.action()
	}
}

func helpCommand() error {
	return nil
}

func exitCommand() error {
	os.Exit(0)
	return nil
}
