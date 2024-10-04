package repl

import "os"

func (repl *Repl) exitCommand(params []string) error {
	os.Exit(0)
	return nil
}
