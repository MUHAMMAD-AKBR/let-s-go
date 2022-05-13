package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MUHAMMAD-AKBR/let-s-go/structure"
)

func POST(w http.ResponseWriter, r *http.Request) {
	//  the struct that we want
	newproduct := structure.Data{}
	//  parse incoming request body
	decoded := json.NewDecoder(r.Body)
	// disallowUnknownFields job is disallow any field that doesnt match to the struct fields
	decoded.DisallowUnknownFields()
	// convert the request body into a struct by using the decode method
	// then we push it into the slice of struct
	err := decoded.Decode(&newproduct)
	if err != nil {
		fmt.Println(err.Error())
	}
	// it works the slice of struct updates before then it doesnt updates maybe becuase i use := instead of = because := means i declare a new variable rather than overwriting the variable
	Struct_repo.List_of_data = append(Struct_repo.List_of_data, newproduct)
}
