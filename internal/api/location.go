package api

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type Location struct {
	Type      string        `json:"type"`
	Latitude  float64       `json:"latitude"`
	Longitude float64       `json:"longitude"`
	ID        string        `json:"id"`
	Name      string        `json:"name"`
	Location  GeoLocation   `json:"location"`
	Products  Products      `json:"products"`
	Station   Station       `json:"station"`
	Lines     []Line        `json:"lines"`
	Entrances []GeoLocation `json:"entrances"`
}

func (a *Api) GetLocations(query string, results int, fuzzy, stops, linesOfStops bool, lang string) ([]Location, error) {
	q := map[string]string{
		"query":        query,
		"results":      strconv.FormatInt(int64(results), 10),
		"fuzzy":        strconv.FormatBool(fuzzy),
		"stops":        strconv.FormatBool(stops),
		"addresses":    "false",
		"poi":          "false",
		"linesOfStops": strconv.FormatBool(linesOfStops),
		"language":     lang,
		"pretty":       "false",
	}
	res, err := a.get(a.buildQuery("/locations", q))
	if err != nil {
		return []Location{}, err
	}
	defer res.Body.Close()
	var locations []Location
	err = json.NewDecoder(res.Body).Decode(&locations)
	if err != nil {
		return []Location{}, err
	}
	return locations, nil
}

func (a *Api) GetNearby(lat, long float64, results, distance int, stops, linesOfStops bool, language string) ([]Location, error) {
	q := map[string]string{
		"latitude":     strconv.FormatFloat(lat, 'f', -1, 64),
		"longitude":    strconv.FormatFloat(long, 'f', -1, 64),
		"results":      strconv.FormatInt(int64(results), 10),
		"distance":     strconv.FormatInt(int64(distance), 10),
		"stops":        strconv.FormatBool(stops),
		"addresses":    "false",
		"poi":          "false",
		"linesOfStops": strconv.FormatBool(linesOfStops),
		"language":     language,
		"pretty":       "false",
	}
	res, err := a.get(a.buildQuery("/locations/nearby", q))
	if err != nil {
		return []Location{}, err
	}
	defer res.Body.Close()
	var locations []Location
	err = json.NewDecoder(res.Body).Decode(&locations)
	if err != nil {
		return []Location{}, err
	}
	return locations, nil
}

func (a *Api) GetLocationById(id string, linesOfStops bool, language string) (Location, error) {
	q := map[string]string{
		"linesOfStops": strconv.FormatBool(linesOfStops),
		"language":     language,
		"pretty":       "false",
	}
	res, err := a.get(a.buildQuery("/stops/"+id, q))
	if err != nil {
		return Location{}, err
	}
	defer res.Body.Close()
	var location Location
	err = json.NewDecoder(res.Body).Decode(&location)
	if err != nil {
		return Location{}, err
	}
	return location, nil
}

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
	fmt.Println(a.buildQuery("/stops/"+id+"/departures", q))
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
