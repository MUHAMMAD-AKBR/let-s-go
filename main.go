package main

import (
	"fmt"
	"net/http"

	"github.com/MUHAMMAD-AKBR/let-s-go/routes"
	"github.com/MUHAMMAD-AKBR/let-s-go/structure"
)

func main() {
	// product is the instanceof the struct data in the struct.go file
	Product := structure.Data{Product_name: "Webcam", Price: 90, Type_product: "electronics"}

	// ton of bytes contains a slice of bytes because the function convert_to_json does convert the struct(product) into a slice of bytes
	ton_of_bytes := Product.Convert_to_json(&Product)

	// to read that json we need to convert that into a string
	fmt.Println(string(ton_of_bytes))

	// ------ STARTING THE SERVER ------//
	// settings up the http handler using the router package

	http.HandleFunc("/", routes.GET)
	http.ListenAndServe(":3000", nil)
}
