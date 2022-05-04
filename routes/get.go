package routes

// the file that handles get requests

import (
	"fmt"
	"net/http"
)

func GET(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "testing")
}
