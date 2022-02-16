package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vitt-bagal/go-url-shortener/handler"
)

func main() {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	// calls to handlers functions
	myRouter.HandleFunc("/", handler.HomePage)
	// start server on 9090 port
	fmt.Println("Starting server....")
	log.Fatal(http.ListenAndServe(":9090", myRouter))
}
