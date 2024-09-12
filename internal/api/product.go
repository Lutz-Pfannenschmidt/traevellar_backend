package api

type ProductLine struct {
	ProductLine string `json:"productLine"`
	Segment     string `json:"segment"`
}

type Products struct {
	NationalExpress bool `json:"nationalExpress"`
	National        bool `json:"national"`
	RegionalExpress bool `json:"regionalExpress"`
	Regional        bool `json:"regional"`
	Suburban        bool `json:"suburban"`
	Bus             bool `json:"bus"`
	Ferry           bool `json:"ferry"`
	Subway          bool `json:"subway"`
	Tram            bool `json:"tram"`
	Taxi            bool `json:"taxi"`
}

var ALL_PRODUCTS = &Products{
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
}

var TRAIN_PRODUCTS = &Products{
	NationalExpress: true,
	National:        true,
	RegionalExpress: true,
	Regional:        true,
	Suburban:        true,
}
