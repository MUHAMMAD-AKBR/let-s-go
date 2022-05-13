package routes

import (
	"fmt"

	"github.com/MUHAMMAD-AKBR/let-s-go/structure"
)

func DELETE(the_slice []structure.Data, delete_index int) {
	the_slice = append(the_slice[:delete_index], the_slice[delete_index+1:]...)
	fmt.Println(the_slice)
}
