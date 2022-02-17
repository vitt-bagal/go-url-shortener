package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vitt-bagal/go-url-shortener/shortener"
	"github.com/vitt-bagal/go-url-shortener/store"
)

type UrlRequest struct {
	LongUrl string `json:"long_url"`
}

var host = "http://localhost:9090/"

// Home handler function
func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to go-url-shortner...")
}

// ShortURL handler function
func CreateShortUrlHandler(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var long_url UrlRequest
	json.Unmarshal(reqBody, &long_url)
	//fmt.Println("passed Longurl is", long_url.LongUrl)
	//fmt.Fprintf(w, "%+v", string(reqBody))
	shortUrl := host + shortener.GenerateShortURL()
	// Copy urls to file
	surl := store.StoreToFile(shortUrl, long_url.LongUrl)
	//fmt.Println(surl)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(surl)
}

//Redirect short-url to long-url handler function
func RedirectShortUrlHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	surl := params["short-url"]
	shortUrl := host + surl
	// Get longurl to file
	lurl := store.GetLongURLFromFile(shortUrl)
	// fmt.Println(lurl)
	http.Redirect(w, r, lurl, http.StatusFound)

}
