package handler

import (
	"ekart/db"
	"ekart/model"
	"encoding/json"
	"net/http"
)

func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	var customer model.Customer

	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "Aplication/json")
	w.WriteHeader(http.StatusCreated)
	id := db.NewCustomer(customer)
	err = json.NewEncoder(w).Encode(map[string]int{"id": id})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
