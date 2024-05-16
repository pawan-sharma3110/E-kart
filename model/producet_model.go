package model

type Products struct {
	CompanyName string  `json:"company_name"`
	ItemName    string  `json:"item_name"`
	Quantity    string  `json:"quantity"`
	Cost        float32 `json:"cost"`
}
