package model

import (
	"encoding/json"
)

type Place struct {
	Name      string `json:"name"`
	OtherName struct {
		English string `json:"name:en"`
	} `json:"other_names"`
	Population int       `json:"population"`
	Location   []float64 `json:"location"`
	Bbox       []float64 `json:"bbox"`
}

func (p Place) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Name       string `json:"name"`
		Population int    `json:"population"`
	}{
		Name:       p.OtherName.English,
		Population: p.Population,
	})
}
