package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/4hakke/repl-pokedex/internal/pokedex/model"
)

type PokedexProviderInterface interface {
	GetLocationArea(name string) (model.LocationArea, error)
	LocationsNext() ([]model.Location, error)
	LocationsPrevious() ([]model.Location, error)
	Catch(pokemonName string) (bool, error)
	Inspect(pokemonName string) (model.Pokemon, error)
	CaughtPokemons() ([]model.Pokemon, error)
}

func NewRepl(pokedexProvider PokedexProviderInterface) *Repl {
	return &Repl{pokedexProvider: pokedexProvider}
}

type Repl struct {
	pokedexProvider PokedexProviderInterface
}

func (repl *Repl) Start() {
	buildCommands(repl)
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

		var commandErr error
		if len(enteredCommand) > 1 {
			commandErr = command.action(enteredCommand[1:])
		} else {
			commandErr = command.action(make([]string, 0))
		}

		if commandErr != nil {
			fmt.Println(commandErr)
		}
	}
}

func buildCommands(repl *Repl) {
	commands = map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			action:      repl.helpCommand,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			action:      repl.exitCommand,
		},
		"map": {
			name:        "map",
			description: "Iterate through list of locations",
			action:      repl.mapCommand,
		},
		"mapb": {
			name:        "mapb",
			description: "Iterate through list of locations (backwards)",
			action:      repl.mapBCommand,
		},
		"explore": {
			name:        "explore",
			description: "Explore the pokemons in specific area. Area name should be provided as an argument",
			action:      repl.exploreCommand,
		},
		"catch": {
			name:        "catch",
			description: "Attempt to catch a pokemon. Pokemon name should be provided as an argument",
			action:      repl.catchCommand,
		},
		"inspect": {
			name:        "inspect",
			description: "Check basic information about caught pokemon. Pokemon name should be provided as an argument",
			action:      repl.inspectCommand,
		},
		"pokedex": {
			name:        "pokedex",
			description: "List all pokemons you caught",
			action:      repl.pokedexCommand,
		},
	}
}
