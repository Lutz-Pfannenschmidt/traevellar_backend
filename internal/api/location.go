package api

import (
	"encoding/json"
	"strconv"
)

type Location struct {
	Type      string      `json:"type"`
	Latitude  float64     `json:"latitude"`
	Longitude float64     `json:"longitude"`
	ID        string      `json:"id"`
	Name      string      `json:"name"`
	Location  GeoLocation `json:"location"`
	Products  Products    `json:"products"`
	Station   Station     `json:"station"`
	Lines     []Line      `json:"lines"`
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
