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
	myRouter.HandleFunc("/", handler.HomePageHandler)
	myRouter.HandleFunc("/create-short-url", handler.CreateShortUrlHandler).Methods("POST")
	//	myRouter.HandleFunc("/{short-url}", handler.RedirectShortUrl)
	// start server on 9090 port
	fmt.Println("Starting server on port 9090....")
	log.Fatal(http.ListenAndServe(":9090", myRouter))
}
