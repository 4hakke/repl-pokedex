package model

type LocationsResult struct {
	Results  []Location `json:"results"`
	Count    int        `json:"count"`
	Previous *string    `json:"previous"`
	Next     *string    `json:"next"`
}
