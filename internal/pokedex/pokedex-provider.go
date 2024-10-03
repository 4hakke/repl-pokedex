package pokedex

import (
	"fmt"

	"github.com/4hakke/repl-pokedex/internal/pokedex/model"
)

type NetworkClientInterface interface {
	Get(url string, resultedObject any) error
}

type PokedexProvider struct {
	networkClient NetworkClientInterface
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

func (provider *PokedexProvider) Locations(offset, limit int) ([]model.Location, error) {
	fullUrl := fmt.Sprintf("%s/location/?offset=%d&limit=%d", baseUrl, offset, limit)
	locationsResult := model.LocationsResult{}
	err := provider.networkClient.Get(fullUrl, &locationsResult)
	if err != nil {
		return []model.Location{}, err
	}

	return locationsResult.Results, nil
}
