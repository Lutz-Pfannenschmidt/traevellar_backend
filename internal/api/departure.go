package api

import (
	"encoding/json"
	"strconv"
)

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
	Remarks             []Remark    `json:"remarks"`
	Origin              any         `json:"origin"`
	Destination         Location    `json:"destination"`
	CurrentTripPosition GeoLocation `json:"currentTripPosition"`
}

// GetDepartures returns a list of departures for the given location
func (a *Api) GetDepartures(id, when string, duration, results int, linesOfStops, remarks bool, config Products) ([]Departure, error) {
	q := map[string]string{
		"when":            when,
		"duration":        strconv.FormatInt(int64(duration), 10),
		"results":         strconv.FormatInt(int64(results), 10),
		"linesOfStops":    strconv.FormatBool(linesOfStops),
		"remarks":         strconv.FormatBool(remarks),
		"nationalExpress": strconv.FormatBool(config.NationalExpress),
		"national":        strconv.FormatBool(config.National),
		"regionalExpress": strconv.FormatBool(config.RegionalExpress),
		"regional":        strconv.FormatBool(config.Regional),
		"suburban":        strconv.FormatBool(config.Suburban),
		"bus":             strconv.FormatBool(config.Bus),
		"ferry":           strconv.FormatBool(config.Ferry),
		"subway":          strconv.FormatBool(config.Subway),
		"tram":            strconv.FormatBool(config.Tram),
		"taxi":            strconv.FormatBool(config.Taxi),
		"pretty":          "false",
	}
	res, err := a.get(a.buildQuery("/stops/"+id+"/departures", q))
	if err != nil {
		return []Departure{}, err
	}
	defer res.Body.Close()
	var departures departureResponse
	err = json.NewDecoder(res.Body).Decode(&departures)
	if err != nil {
		return []Departure{}, err
	}
	return departures.Departures, nil
}
