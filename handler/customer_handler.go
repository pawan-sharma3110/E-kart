package handler

import (
	"ekart/db"
	"ekart/model"
	"encoding/json"
	"net/http"
)

func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Aplication/json")
	w.WriteHeader(http.StatusCreated)
	var RequestData struct {
		NewCustomer model.Customer
		Address     model.Address
	}
	json.NewDecoder(r.Body).Decode(&RequestData)
	db.SaveCustomer(RequestData.NewCustomer, RequestData.Address)
	json.NewEncoder(w).Encode(RequestData)

}
