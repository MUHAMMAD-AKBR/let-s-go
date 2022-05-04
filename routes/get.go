package routes

// the file that handles get requests

import (
	"fmt"
	"net/http"

	"github.com/MUHAMMAD-AKBR/let-s-go/structure"
)

func GET(w http.ResponseWriter, r *http.Request) {
	// product is the instanceof the struct data in the struct.go file
	w.Header().Set("Content-type", "Application/json")
	Product := structure.Data{Product_name: "Webcam", Price: 90, Type_product: "electronics"}

	// ton of bytes returns a slice of bytes because the function convert_to_json does convert the struct(product) into a slice of bytes
	ton_of_bytes, err := Product.Convert_to_json(&Product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "error: %v", err)
	}
	w.WriteHeader(http.StatusAccepted)

	// to read that json we need to convert that into a string
	fmt.Fprintf(w, "data: %v", string(ton_of_bytes))
}
