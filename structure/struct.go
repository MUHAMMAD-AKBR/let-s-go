package structure

import (
	"encoding/json"
	"fmt"
)

// the structure of the data we wan't but this is still in the struct format, we want to marshal it into json later
type Data struct {
	Product_name string
	Price        string
	Type_product string
}

// list of structs
type Repo struct {
	List_of_data []Data
}

//  list of jsonized structs
type Json_Repo struct {
	List_of_json []string
}

var Struct_repo = Repo{
	List_of_data: []Data{
		{Product_name: "mirror", Price: "30", Type_product: "housekeepings"},
		{Product_name: "blanket", Price: "10", Type_product: "housekeepings"},
		{Product_name: "pillow", Price: "20", Type_product: "housekeepings"},
		{Product_name: "bodypillow", Price: "25", Type_product: "housekeepings"},
		{Product_name: "usb", Price: "30", Type_product: "Electronics"},
		{Product_name: "glass-cup", Price: "0.5", Type_product: "housekeepings"},
	},
}

func (d *Data) Convert_to_json(thing *Data) ([]byte, error) {
	json, err := json.Marshal(*thing)
	if err != nil {
		return nil, fmt.Errorf("error because %v", err)
	}
	return json, nil
}
