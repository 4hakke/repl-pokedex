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

func Locations(offset, limit int) ([]Location, error) {
	response, err := http.Get("https://pokeapi.co/api/v2/location/")
	if err != nil {
		return []Location{}, err
	}

	defer response.Body.Close()
	locationsResult := LocationsResult{}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return []Location{}, err
	}

	err = json.Unmarshal(body, &locationsResult)
	if err != nil {
		return []Location{}, err
	}

	fmt.Printf("%v", locationsResult.Results)

	return locationsResult.Results, nil
}
