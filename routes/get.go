package routes

// the file that handles get requests

import (
	"fmt"
	"net/http"

	"github.com/MUHAMMAD-AKBR/let-s-go/structure"
)

//repository := structure.Repo{
//	List_of_data: []structure.Data{
//		{Product_name: "mirror", Price: "30", Type_product: "housekeepings"},
//		{Product_name: "blanket", Price: "10", Type_product: "housekeepings"},
//			{Product_name: "pillow", Price: "20", Type_product: "housekeepings"},
//	{Product_name: "bodypillow", Price: "25", Type_product: "housekeepings"},
//	{Product_name: "usb", Price: "30", Type_product: "Electronics"},
//},
//}

var repo = structure.Repo{
	List_of_data: []structure.Data{
		{Product_name: "mirror", Price: "30", Type_product: "housekeepings"},
		{Product_name: "blanket", Price: "10", Type_product: "housekeepings"},
		{Product_name: "pillow", Price: "20", Type_product: "housekeepings"},
		{Product_name: "bodypillow", Price: "25", Type_product: "housekeepings"},
		{Product_name: "usb", Price: "30", Type_product: "Electronics"},
	},
}

func GET(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusAccepted)
	for _, v := range repo.List_of_data {
		data, err := v.Convert_to_json(&v)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			panic(err)
		}
		fmt.Fprintf(w, ",%v", string(data))
	}
}
