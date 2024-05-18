package main

import (
	"ekart/db"
	"ekart/router"
	"log"
	"net/http"
)

func main() {
	db.DbInIt()

	r := router.MyRouter()

	log.Fatal(http.ListenAndServe(":8080", r))
}
