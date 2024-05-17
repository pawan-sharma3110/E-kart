package router

import (
	"ekart/handler"

	"github.com/gorilla/mux"
)

func MyRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/ekart/customer/create", handler.CreateCustomer).Methods("POST")
	return router
}
