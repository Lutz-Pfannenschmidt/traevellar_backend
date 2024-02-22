package main

import (
	"github.com/Lutz-Pfannenschmidt/traevellar_backend/internal/api"
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
		println(location.Name)
	}

}
