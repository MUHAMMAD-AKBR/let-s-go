package routes

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"

	"github.com/MUHAMMAD-AKBR/let-s-go/regex"
	"github.com/MUHAMMAD-AKBR/let-s-go/structure"
)

func PATCH(r *http.Request) {
	// instance of struct
	var instace_of_struct = structure.Data{}
	var path = r.URL.Path
	// this is the function to make the struct iterable
	typeof := reflect.TypeOf(instace_of_struct)
	// this will return a slice of string than contains the struct names values and so on
	slice_of_fields := reflect.VisibleFields(typeof)

	// we read the request body
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("there are no keys provided")
	}
	// strinigfy the request body and then validate the string using regex to find all the leter and _
	stringified := string(bytes)
	// return a slice of string than contains the correct keys of struct on the first index the second index is the value of struct
	filtered_stringified := regex.Only_letters(stringified)

	for _, value := range slice_of_fields {
		// validate if the value name of slice_of_fields if correct

		if value.Name == filtered_stringified[0] {
			fmt.Println("this is valid")
		}

		// get the index from url path
		// test the string first if its the right form like /datas/2 <-- index
		result, _ := regex.Test_string(path)
		// get the int finally that is the index
		final_result := regex.Find_int(strings.Join(result, ""))
		fmt.Println(final_result)
	}

}
