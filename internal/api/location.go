package api

import (
	"encoding/json"
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

// GetLocations returns a list of locations based on the query string
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

// GetNearby returns a list of locations near the given coordinates
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

// GetLocationById returns a location based on the given id
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
