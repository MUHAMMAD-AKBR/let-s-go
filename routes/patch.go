package routes

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/MUHAMMAD-AKBR/let-s-go/regex"
	"github.com/MUHAMMAD-AKBR/let-s-go/structure"
)

func PATCH(r *http.Request) {
	// pointer to the real repo
	var ptr *structure.Repo = &structure.Struct_repo

	// read the request body
	text, _ := ioutil.ReadAll(r.Body)

	// find index
	slice_of_url_path, err := regex.Test_string(r.URL.Path)
	if err != nil {
		fmt.Println("path is not right it should be like localhost:3000/datas/<number>")
	}
	index := regex.Find_int(strings.Join(slice_of_url_path, ""))

	// get key_from_text
	remove_after_colon := regex.Paramater_wants(string(text), `.+:`)
	remove_exept_key := regex.Paramater_wants(strings.Join(remove_after_colon, ""), `[a-zA-Z_]+`)
	stringified_value_from_text1 := strings.Join(remove_exept_key, "")

	// get value_from_text
	remove_before_colon := regex.Paramater_wants(string(text), `:.+`)
	remove_exept_value := regex.Paramater_wants(strings.Join(remove_before_colon, ""), `[a-zA-Z_0-9]+`)
	stringified_value_from_txt2 := strings.Join(remove_exept_value, "")

	if stringified_value_from_text1 == "Product_name" {
		fmt.Printf("that is a correct key=Product_name! here is the value %v", stringified_value_from_txt2)
		ptr.List_of_data[index].Product_name = stringified_value_from_txt2
	} else if stringified_value_from_text1 == "Price" {
		fmt.Printf("that is a correct key=price! here is the value %v", stringified_value_from_txt2)
		ptr.List_of_data[index].Price = stringified_value_from_txt2
	} else if stringified_value_from_text1 == "Type_product" {
		fmt.Printf("that is a correct key=Type_product! here is the value %v", stringified_value_from_txt2)
		ptr.List_of_data[index].Type_product = stringified_value_from_txt2
	}
}
