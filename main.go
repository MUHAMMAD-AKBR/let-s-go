package main

import (
	"fmt"
	"net/http"
	_ "strings"

	"github.com/MUHAMMAD-AKBR/let-s-go/routes"
)

func main() {

	// ------ STARTING THE SERVER ------//
	// settings up the http handler using the router package

	http.HandleFunc("/datas", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			w.Header().Set("Content-type", "application/json")
			data := routes.GET(w, r)
			fmt.Fprintln(w, data)
		case "POST":
			routes.POST(w, r)
		case "PUT", "PATCH":
			fmt.Fprintf(w, "still working on it your request is PUT/PATCH")
		case "DELETE":
			fmt.Fprintln(w, "still working on it your request is DELETE")
		}

	})

	http.ListenAndServe(":3000", nil)
}
