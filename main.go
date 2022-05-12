package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/MUHAMMAD-AKBR/let-s-go/regex"
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

	http.HandleFunc("/datas/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			w.Header().Set("Content-type", "application/json")
			// take the url and find if its accurate
			slice, _ := regex.Test_string(r.URL.Path)
			// result is a function that returns an int from the url if there is one like /datas/3 <-- 3 is the int
			// and then convert it to string
			result := regex.Find_int(strings.Join(slice, ""))
			// each is the struct with the index of the result
			each := routes.Struct_repo.List_of_data[result]
			// marhsal it into json and get the data
			byt, _ := json.Marshal(each)
			fmt.Fprintln(w, string(byt))
		}
	})

	http.ListenAndServe(":3000", nil)
}
