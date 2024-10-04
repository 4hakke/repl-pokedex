package main

import (
	"time"

	"github.com/4hakke/repl-pokedex/internal/cache"
	networkclient "github.com/4hakke/repl-pokedex/internal/network-client"
	"github.com/4hakke/repl-pokedex/internal/pokedex"
	"github.com/4hakke/repl-pokedex/internal/repl"
)

func main() {
	cache := cache.NewCache(20 * time.Second)
	networkClient := networkclient.NewClient(cache)
	pokedexProvider := pokedex.NewProvider(networkClient)

	repl := repl.NewRepl(pokedexProvider)
	repl.Start()
}
