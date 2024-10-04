package repl

type cliCommand struct {
	action      func(params []string) error
	name        string
	description string
}

var commands map[string]cliCommand
