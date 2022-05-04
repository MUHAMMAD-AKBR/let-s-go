package main

import (
	"net/http"

	"github.com/MUHAMMAD-AKBR/let-s-go/routes"
)

func main() {
	// ------ STARTING THE SERVER ------//
	// settings up the http handler using the router package

	http.HandleFunc("/", routes.GET)

	http.ListenAndServe(":3000", nil)
}
