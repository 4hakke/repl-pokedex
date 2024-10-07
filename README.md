### Overview

Purpose of the repo is to practice some Go learnings obtained through joy, sweat and tears at [boot.dev](https://www.boot.dev).

The program itself is yet another pokemon cli that allows to complete some basic operations:

- iterate through locations
- check for pokemons in a location area
- catch pokemons
- and inspect basic stats of caught pokemons

#### Structure

The program is split in 4 internal packages:

- repl - read command input and coordinate to corresponding supported command
- pokedex - contains some domain knowledge about pokemons and actions
- network-client - generic layer to make GET request
- cache - generic layer for caching
