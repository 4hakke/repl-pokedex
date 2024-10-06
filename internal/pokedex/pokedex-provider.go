package pokedex

import (
	"errors"
	"fmt"
	"maps"
	"math/rand"
	"slices"

	"github.com/4hakke/repl-pokedex/internal/pokedex/model"
)

func NewProvider(networkClient NetworkClientInterface) *PokedexProvider {
	return &PokedexProvider{networkClient: networkClient, state: &State{caughtPokemons: make(map[string]model.Pokemon)}}
}

type NetworkClientInterface interface {
	Get(url string, resultedObject any) error
}

type PokedexProvider struct {
	networkClient NetworkClientInterface
	state         *State
}

type State struct {
	locationsResult *model.LocationsResult
	caughtPokemons  map[string]model.Pokemon
}

const baseUrl = "https://pokeapi.co/api/v2/"

func (provider *PokedexProvider) GetLocationArea(name string) (model.LocationArea, error) {
	fullUrl := fmt.Sprintf("%s/location-area/%s", baseUrl, name)
	locationArea := model.LocationArea{}
	err := provider.networkClient.Get(fullUrl, &locationArea)
	if err != nil {
		return model.LocationArea{}, err
	}

	return locationArea, nil
}

func (provider *PokedexProvider) LocationsNext() ([]model.Location, error) {
	if provider.state.locationsResult == nil {
		url := fmt.Sprintf("%s/location/?offset=%d&limit=%d", baseUrl, 0, 20)
		return provider.locations(url)
	}

	if provider.state.locationsResult.Next != nil {
		return provider.locations(*provider.state.locationsResult.Next)
	} else {
		return []model.Location{}, errors.New("You reached the end of the locations list")
	}
}

func (provider *PokedexProvider) LocationsPrevious() ([]model.Location, error) {
	if provider.state.locationsResult == nil {
		return []model.Location{}, errors.New("You reached the beginning of the locations list")
	}

	if provider.state.locationsResult.Previous != nil {
		return provider.locations(*provider.state.locationsResult.Previous)
	} else {
		return []model.Location{}, errors.New("You reached the beginning of the locations list")
	}
}

func (provider *PokedexProvider) Catch(pokemonName string) (bool, error) {
	if _, ok := provider.state.caughtPokemons[pokemonName]; ok {
		return false, errors.New("You already caught this pokemon")
	}

	fullUrl := fmt.Sprintf("%s/pokemon/%s", baseUrl, pokemonName)
	pokemon := model.Pokemon{}
	err := provider.networkClient.Get(fullUrl, &pokemon)
	if err != nil {
		return false, err
	}

	catchChance := rand.Intn(550)

	didCatchPokemon := catchChance > pokemon.BaseExperience

	if didCatchPokemon {
		provider.state.caughtPokemons[pokemon.Name] = pokemon
	}

	return catchChance > pokemon.BaseExperience, nil
}

func (provider *PokedexProvider) Inspect(pokemonName string) (model.Pokemon, error) {
	pokemon, ok := provider.state.caughtPokemons[pokemonName]
	if !ok {
		return model.Pokemon{}, errors.New("Hold on, you need to catch the pokemon first")
	}

	return pokemon, nil
}

func (provider *PokedexProvider) CaughtPokemons(pokemonName string) ([]model.Pokemon, error) {
	return slices.Collect(maps.Values(provider.state.caughtPokemons)), nil
}

func (provider *PokedexProvider) locations(url string) ([]model.Location, error) {
	locationsResult := model.LocationsResult{}
	err := provider.networkClient.Get(url, &locationsResult)
	if err != nil {
		return []model.Location{}, err
	}
	provider.state.locationsResult = &locationsResult

	return locationsResult.Results, nil
}
