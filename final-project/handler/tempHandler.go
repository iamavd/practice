package handler

import (
	"fmt"
	"net/http"
)

//Handler for test
func Sayhello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}
