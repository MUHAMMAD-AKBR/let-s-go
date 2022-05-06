package main

import (
	"net/http"

	"github.com/MUHAMMAD-AKBR/let-s-go/routes"
)

func main() {

	// ------ STARTING THE SERVER ------//
	// settings up the http handler using the router package

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			routes.GET(w, r)
		case "POST":
			routes.POST(w, r)
		}

	})

	http.ListenAndServe(":3000", nil)
}
