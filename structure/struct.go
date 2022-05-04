package structure

import (
	"encoding/json"
	"fmt"
)

// the structure of the data we wan't but this is still in the struct format, we want to marshal it into json later
type Data struct {
	Product_name string
	Price        int
	Type_product string
}

func (d *Data) Convert_to_json(thing *Data) ([]byte, error) {
	json, err := json.Marshal(*thing)
	if err != nil {
		return nil, fmt.Errorf("error because %v", err)
	}
	return json, nil
}
