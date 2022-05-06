package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MUHAMMAD-AKBR/let-s-go/structure"
)

func POST(w http.ResponseWriter, r *http.Request) ([]byte, error) {
	//  the struct that we want
	newproduct := structure.Data{}
	//  parse incoming request body
	decoded := json.NewDecoder(r.Body)
	// disallowUnknownFields job is disallow any field that doesnt match to the struct fields
	decoded.DisallowUnknownFields()
	// convert the request body into a json
	// somehow the decode manages to add the values of the correct request body to the newproduct
	err := decoded.Decode(&newproduct)
	if err != nil {
		fmt.Println(err.Error())
	}
	bytes, ok := newproduct.Convert_to_json(&newproduct)
	if ok != nil {
		return nil, fmt.Errorf("errro %v", ok)
	}
	return bytes, nil

}
