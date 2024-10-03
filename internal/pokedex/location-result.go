package pokedex

type LocationsResult struct {
	Results []Location `json:"results"`
	Count   int        `json:"count"`
}
