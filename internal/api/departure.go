package api

type departureResponse struct {
	Departures            []Departure `json:"departures"`
	RealtimeDataUpdatedAt any         `json:"realtimeDataUpdatedAt"`
}

type Departure struct {
	TripId              string      `json:"tripId"`
	Stop                Location    `json:"stop"`
	When                string      `json:"when"`
	PlannedWhen         string      `json:"plannedWhen"`
	Delay               int         `json:"delay"`
	Platform            string      `json:"platform"`
	PlannedPlatform     string      `json:"plannedPlatform"`
	PrognosisType       string      `json:"prognosisType"`
	Direction           string      `json:"direction"`
	Provenance          any         `json:"provenance"`
	Line                Line        `json:"line"`
	Remarks             []any       `json:"remarks"`
	Origin              any         `json:"origin"`
	Destination         Location    `json:"destination"`
	CurrentTripPosition GeoLocation `json:"currentTripPosition"`
}
