package main

import (
	"strconv"
	"time"

	"github.com/Lutz-Pfannenschmidt/traevellar_backend/internal/api"
	"github.com/Lutz-Pfannenschmidt/yagll"
)

const (
	API_URL  = "https://v6.db.transport.rest"
	ALPHABET = "abcdefghijklmnopqrstuvwxyzäüöß"
)

var db api.Api

func main() {
	db = *api.NewApi(API_URL)

	locations, err := db.GetLocations("Holtwick", 99, true, true, true, "en")
	if err != nil {
		panic(err)
	}

	for _, location := range locations {
		yagll.Infof("Location: %s %s", location.ID, location.Name)
	}

	byId, err := db.GetLocationById("906075", true, "en")
	if err != nil {
		panic(err)
	}

	yagll.Infof("Location by id: %s %s", byId.ID, byId.Name)

	unixTime := time.Now().Unix()
	_, err = db.GetDepartures("906075", strconv.FormatInt(unixTime, 10), 100, 99, true, true, api.Products{
		NationalExpress: true,
		National:        true,
		RegionalExpress: true,
		Regional:        true,
		Suburban:        true,
		Bus:             true,
		Ferry:           true,
		Subway:          true,
		Tram:            true,
		Taxi:            true,
	})

	if err != nil {
		panic(err)
	}

}
