package handler

import (
	"fmt"
	"net/http"
)

// Home handler function
func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to go-url-shortner...")
}
