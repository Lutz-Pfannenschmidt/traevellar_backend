package api

type GeoLocation struct {
	Type      string  `json:"type"`
	ID        string  `json:"id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
