package store

import (
	"log"
	"os"
)

//Initilize storage to create URLFile
func InitFileStorage() {
	urlFile, err := os.Create("URLFile.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer urlFile.Close()
	log.Println(urlFile)
}
