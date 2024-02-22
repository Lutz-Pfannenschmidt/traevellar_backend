package api

type Station struct {
	Type     string      `json:"type"`
	ID       string      `json:"id"`
	Name     string      `json:"name"`
	Location GeoLocation `json:"location"`
}
