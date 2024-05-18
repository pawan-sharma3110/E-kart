package model

type Ekart struct {
	Customer Customer `json:"customer"`
	Products Products `json:"product"`
	Order    Order    `json:"order"`
}

