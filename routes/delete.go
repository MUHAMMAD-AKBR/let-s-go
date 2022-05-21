package routes

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/MUHAMMAD-AKBR/let-s-go/regex"
	"github.com/MUHAMMAD-AKBR/let-s-go/structure"
)

func DELETE(w http.ResponseWriter, r *http.Request) {
	slice_of_correct_string, err := regex.Test_string(r.URL.Path)
	if err != nil {
		fmt.Println("the path is not correct")
	}
	the_index := regex.Find_int(strings.Join(slice_of_correct_string, ""))
	// append(the_slice[:delete_index], the_slice[delete_index+1:]...)
	structure.Struct_repo.List_of_data = append(structure.Struct_repo.List_of_data[:the_index], structure.Struct_repo.List_of_data[the_index+1:]...)
	fmt.Fprintln(w, structure.Struct_repo.List_of_data)
}
