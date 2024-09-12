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

	locations, err := db.GetLocations("Holtwick", 10, true, true, true, "en")
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
	departures, err := db.GetDepartures("906075", strconv.FormatInt(unixTime, 10), 100, 99, true, true, *api.TRAIN_PRODUCTS)

	if err != nil {
		panic(err)
	}

	for _, departure := range departures {
		yagll.Infof("Departure: %s %s %s %s", departure.TripId, departure.When, departure.Platform, departure.Direction)
		yagll.Infof("Stop: %s %s", departure.Stop.ID, departure.Stop.Name)
		yagll.Infof("Line: %s %s %s %s", departure.Line.ID, departure.Line.Name, departure.Line.Product, departure.Line.ProductName)
		for _, remark := range departure.Remarks {
			yagll.Infof("Remark: %v", remark)
		}
	}

}
