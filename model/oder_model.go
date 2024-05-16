package model

import "time"

type Oder struct {
	Customer        Customer `json:"customer"`
	PaymentStatus   string   `json:"payment_status"`
	Products        Products `json:"products"`
	DeliveryDetails Delivery `json:"delivery_details"`
}

type Delivery struct {
	DeliveryStatus string    `json:"delivery_status"`
	Address        Address   `json:"address"`
	DeliveryType   string    `json:"delivery_type"`
	Delivery_date  time.Time `json:"delivery_date"`
}
