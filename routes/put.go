package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/MUHAMMAD-AKBR/let-s-go/regex"
	"github.com/MUHAMMAD-AKBR/let-s-go/structure"
)

func PUT(r *http.Request) {
	instance := structure.Data{}
	var path = r.URL.Path
	catch_body := json.NewDecoder(r.Body)
	catch_body.DisallowUnknownFields()
	err := catch_body.Decode(&instance)
	if err != nil {
		fmt.Println(err.Error())
	}
	// now instance is filled with the correct r.Body data

	//  use the index of r.Path eg: /datas/2 <-- index
	slice_of_string, _ := regex.Test_string(path)
	// slice_of_string is the right path because it contains a number in the end of the path
	index := regex.Find_int(strings.Join(slice_of_string, ""))
	// and then in the slice of struct replace the index reffered to the struct
	fmt.Println(structure.Struct_repo.List_of_data[index])
}
