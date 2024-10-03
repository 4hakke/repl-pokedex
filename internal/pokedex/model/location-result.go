package model

type LocationsResultDto struct {
	Results []Location `json:"results"`
	Count   int        `json:"count"`
}
