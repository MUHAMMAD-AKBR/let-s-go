package routes

// the file that handles get requests

import (
	"fmt"
	"net/http"

	"github.com/MUHAMMAD-AKBR/let-s-go/structure"
)

// list of struct we ar going to convert it into a list of json

var Struct_repo = structure.Repo{
	List_of_data: []structure.Data{
		{Product_name: "mirror", Price: "30", Type_product: "housekeepings"},
		{Product_name: "blanket", Price: "10", Type_product: "housekeepings"},
		{Product_name: "pillow", Price: "20", Type_product: "housekeepings"},
		{Product_name: "bodypillow", Price: "25", Type_product: "housekeepings"},
		{Product_name: "usb", Price: "30", Type_product: "Electronics"},
	},
}

// list of converted struct s into json
var Json_repo = structure.Json_Repo{}

func GET(w http.ResponseWriter, r *http.Request) []string {
	w.WriteHeader(http.StatusAccepted)
	for _, value := range Struct_repo.List_of_data {
		jsonized, err := value.Convert_to_json(&value)
		if err != nil {
			fmt.Println("error cannot convert struct to json")
		}
		Json_repo.List_of_json = append(Json_repo.List_of_json, string(jsonized))
	}
	return Json_repo.List_of_json
}
