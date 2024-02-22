package api

type Line struct {
	Type    string `json:"type"`
	ID      string `json:"id"`
	FahrtNr string `json:"fahrtNr"`
	Name    string `json:"name"`
	Public  bool   `json:"public"`
	Mode    string `json:"mode"`
	Product string `json:"product"`
}
