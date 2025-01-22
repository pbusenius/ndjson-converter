package model

import (
	"encoding/json"
	"fmt"
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
		Type       string `json:"type"`
		Geometry   string `json:"geometry"`
		Properties string `json:"properties"`
	}{
		Type:       "feature",
		Geometry:   fmt.Sprintf("{\"type\": %s, \"coordinates\": %v}", "point", p.Location),
		Properties: fmt.Sprintf("{\"name\": %s, \"population\": %d}", p.OtherName.English, p.Population),
	})
}
