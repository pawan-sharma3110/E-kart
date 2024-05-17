package main

import (
	"ekart/db"
	"ekart/router"
	"net/http"
)

func main() {
	db.DbInIt()

	r := router.MyRouter()

	http.ListenAndServe(":8030", r)
}
