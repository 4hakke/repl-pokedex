package pokedex

import "fmt"

type NetworkClientInterface interface {
	Get(url string, resultedObject any) error
}

type PokedexProvider struct {
	networkClient NetworkClientInterface
}

func (provider *PokedexProvider) GetLocationArea(name string) (LocationArea, error) {
	fullUrl := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", name)
	locationArea := LocationArea{}
	err := provider.networkClient.Get(fullUrl, &locationArea)
	if err != nil {
		return LocationArea{}, err
	}

	return locationArea, nil
}
