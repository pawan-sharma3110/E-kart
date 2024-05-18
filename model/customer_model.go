package model

import "time"

type Customer struct {
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Mobile     string    `json:"mobile"`
	Password   string    `json:"password"`
	Gender     string    `json:"gender"`
	Adult      bool      `json:"adult"`
	Address    *Address  `json:"address"`
	Created_at time.Time `json:"created_at"`
}
type Address struct {
	HouseNo  string `json:"house_no"`
	Town     string `json:"town"`
	District string `json:"district"`
	State    string `json:"state"`
	Pincode  int    `json:"pincode"`
}
