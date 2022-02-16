package shortener

import (
	"math/rand"
	"time"
)

// Generate random short url
func GenerateShortURL() string {
	var base62 = []byte("012345689abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	rand.Seed(time.Now().UnixNano())
	shortId := make([]byte, 8)
	for i := range shortId {
		shortId[i] = base62[rand.Intn(len(base62))]
	}
	return string(shortId)
}
