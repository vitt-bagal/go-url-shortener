package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/vitt-bagal/go-url-shortener/shortener"
	"github.com/vitt-bagal/go-url-shortener/store"
)

type UrlRequest struct {
	LongUrl string `json:"long_url"`
}

// Home handler function
func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to go-url-shortner...")
}

// ShortURL handler function
func CreateShortUrl(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var long_url UrlRequest
	json.Unmarshal(reqBody, &long_url)
	//fmt.Println("passed Longurl is", long_url.LongUrl)
	//fmt.Fprintf(w, "%+v", string(reqBody))
	host := "http://localhost:9090/"
	shortUrl := host + shortener.GenerateShortURL()
	// Copy urls to file
	surl := store.StoreToFile(shortUrl, long_url.LongUrl)
	//fmt.Println(surl)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(surl)
}

/* Redirect short-url to long-url handler function
func RedirectShortUrl(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(short - url)
	var long_url UrlRequest
	json.Unmarshal(reqBody, &long_url)
	//fmt.Println("passed Longurl is", long_url.LongUrl)
	//fmt.Fprintf(w, "%+v", string(reqBody))
	host := "http://localhost:9090/"
	shortUrl := host + shortener.GenerateShortURL()
	// Copy urls to file
	surl := store.StoreToFile(shortUrl, long_url.LongUrl)
	//fmt.Println(surl)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(surl)
}*/
