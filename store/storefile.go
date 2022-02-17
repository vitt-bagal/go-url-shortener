package store

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

// Function to store data to URLFile in short-url=long-url and return shorturl
func StoreToFile(surl, lurl string) string {
	var shorturl string
	urlFile, err := os.OpenFile("URLFile.txt", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	// check for error
	if err != nil {
		log.Fatalln("Could not open URLFile.txt")
	}
	defer urlFile.Close()
	// Check whether LongURl is already present in urlfile if not then add
	if isLongURLExistInFile(lurl, "URLFile.txt") {
		scanner := bufio.NewScanner(urlFile)
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			isPresent, err := regexp.Match(lurl, []byte(scanner.Text()))
			if err != nil {
				panic(err)
			}
			if isPresent {
				matchedLine := scanner.Text()
				shorturl = strings.Split(matchedLine, "=")[0]
			} else {
				continue
			}
		}

	} else {
		_, err2 := urlFile.WriteString(surl + "=" + lurl + "\n")
		if err2 != nil {
			log.Println("Could not write text to URLFile.txt")
		} else {
			log.Println("Operation successful!!! Stored data to file successfully....")
		}
		shorturl = surl
	}
	return shorturl
}

// Function to check string existance in file
func isLongURLExistInFile(str, filepath string) bool {
	b, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}

	isExist, err := regexp.Match(str, b)
	if err != nil {
		panic(err)
	}
	return isExist
}

// Function to get correspoding long url of short-url from URLFile
func GetLongURLFromFile(surl string) string {
	var lurl string
	urlFile, err := os.OpenFile("URLFile.txt", os.O_RDONLY, 0644)
	// check for error
	if err != nil {
		log.Fatalln("Could not open URLFile.txt")
	}
	defer urlFile.Close()
	// Read urlfile line by line and get LongURL corresponding to short-url
	scanner := bufio.NewScanner(urlFile)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		isPresent, err := regexp.Match(surl, []byte(scanner.Text()))
		if err != nil {
			panic(err)
		}
		if isPresent {
			matchedLine := scanner.Text()
			lurl = strings.Split(matchedLine, "=")[1]
		} else {
			continue
		}
	}
	return lurl
}
