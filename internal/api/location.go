package api

import (
	"encoding/json"
	"fmt"
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

func (a *Api) GetLocations(query string, results int, fuzzy bool, stops, linesOfStops bool, lang string, filters ...map[string]string) ([]Location, error) {
	q := map[string]string{
		"query":        query,
		"results":      strconv.FormatInt(int64(results), 10),
		"fuzzy":        strconv.FormatBool(fuzzy),
		"stops":        strconv.FormatBool(stops),
		"addresses":    "false",
		"poi":          "false",
		"linesOfStops": strconv.FormatBool(linesOfStops),
		"lang":         lang,
	}
	filters = append(filters, q)
	fmt.Println(a.buildQuery("/locations", filters...))
	res, err := a.get(a.buildQuery("/locations", filters...))
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
