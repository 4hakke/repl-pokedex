package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationsResult struct {
	Results []Location `json:"results"`
	Count   int        `json:"count"`
}

type Location struct {
	Name string `json:"name"`
}

func locations(offset, limit int) (LocationsResult, error) {
	fullUrl := fmt.Sprintf("https://pokeapi.co/api/v2/location/?offset=%d&limit=%d", offset, limit)
	response, err := http.Get(fullUrl)
	if err != nil {
		return LocationsResult{}, err
	}

	defer response.Body.Close()
	locationsResult := LocationsResult{}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return LocationsResult{}, err
	}

	err = json.Unmarshal(body, &locationsResult)
	if err != nil {
		return LocationsResult{}, err
	}

	return locationsResult, nil
}
