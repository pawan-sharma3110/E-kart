package model

type Customer struct {
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	Mobile   string  `json:"mobile"`
	Password string  `json:"password"`
	Gander   string  `json:"gander"`
	Adult    string  `json:"adult"`
	Address  Address `json:"string"`
}
type Address struct {
	HouseNo  string `json:"house_no"`
	Town     string `json:"town"`
	District string `json:"district"`
	State    string `json:"state"`
	Pincode  int    `json:"pincode"`
}
