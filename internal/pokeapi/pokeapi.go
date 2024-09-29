package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/4hakke/repl-pokedex/internal/pokecache"
)

type LocationsResult struct {
	Results []Location `json:"results"`
	Count   int        `json:"count"`
}

type Location struct {
	Name string `json:"name"`
}

type LocationArea struct {
	Name              string `json:"name"`
	PokemonEncounters []struct {
		Pokemon Pokemon `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

type Pokemon struct {
	Name string `json:"name"`
}

var cache = pokecache.NewCache(20 * time.Second)

// TODO: Refactor
func locations(offset, limit int) (LocationsResult, error) {
	fullUrl := fmt.Sprintf("https://pokeapi.co/api/v2/location/?offset=%d&limit=%d", offset, limit)
	cachedResult, ok := cache.Get(fullUrl)
	if ok {
		return parseLocations(cachedResult)
	}

	response, err := http.Get(fullUrl)
	if err != nil {
		return LocationsResult{}, err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return LocationsResult{}, err
	}
	result, err := parseLocations(body)
	if err == nil {
		cache.Add(fullUrl, body)
	}
	return result, err
}

func GetLocationArea(name string) (LocationArea, error) {
	fullUrl := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", name)
	cachedResult, ok := cache.Get(fullUrl)
	if ok {
		return parseLocationArea(cachedResult)
	}

	response, err := http.Get(fullUrl)
	if err != nil {
		return LocationArea{}, err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return LocationArea{}, err
	}
	result, err := parseLocationArea(body)
	if err == nil {
		cache.Add(fullUrl, body)
	}
	return result, err
}

func parseLocations(payload []byte) (LocationsResult, error) {
	locationsResult := LocationsResult{}
	err := json.Unmarshal(payload, &locationsResult)
	if err != nil {
		return LocationsResult{}, err
	}

	return locationsResult, nil
}

func parseLocationArea(payload []byte) (LocationArea, error) {
	locationArea := LocationArea{}
	err := json.Unmarshal(payload, &locationArea)
	if err != nil {
		return LocationArea{}, err
	}

	return locationArea, nil
}
