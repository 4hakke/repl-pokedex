package main

import (
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
			name:   "name",
			action: helpCommand,
		},
		"exit": {
			name:   "exit",
			action: exitCommand,
		},
	}

	fmt.Println("Hello world")
}

func helpCommand() error {
	return nil
}

func exitCommand() error {
	os.Exit(0)
	return nil
}
