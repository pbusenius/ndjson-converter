package model

type FeatureCollection struct {
	Type     string    `json:"type"`
	Features []Feature `json:"features"`
}

type Feature struct {
	Type       string     `json:"type"`
	Geometry   Geometry   `json:"geometry"`
	Properties Properties `json:"properties"`
}

func NewFeature(geometryType string, place Place) Feature {
	return Feature{
		Type: "Feature",
		Geometry: Geometry{
			Type:        geometryType,
			Coordinates: place.Location,
		},
		Properties: Properties{
			Name:       place.Name,
			Population: place.Population,
		},
	}
}

type Geometry struct {
	Coordinates []float64 `json:"coordinates"`
	Type        string    `json:"type"`
}

type Properties struct {
	Name       string `json:"name"`
	Population int    `json:"population"`
}
