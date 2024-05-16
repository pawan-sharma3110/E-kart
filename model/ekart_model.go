package model

type Ekart struct {
	Customer Customer `json:"customer"`
	Products Products `json:"product"`
	Oder     Oder     `json:"oder"`
}
