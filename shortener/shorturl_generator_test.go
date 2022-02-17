package shortener

import (
	"testing"

	"github.com/vitt-bagal/go-url-shortener/store"
)

func TestGenerateShortURL(t *testing.T) {
	url := GenerateShortURL()
	var length = len([]byte(url))
	if length != 8 {
		t.Errorf("URL Got Nil But expected Non Nil")
	}
}

func TestSameLongURL(t *testing.T) {
	long_url := "https://www.guru3d.com/news-story/philips-monitors-launches-three-new-pc-gaming-monitors-from-its-m3000-series.html"
	actualString := store.StoreToFile("http://localhost:9090/GFbHs4PY", long_url)
	expectedString := "http://localhost:9090/GFbHs4PY"
	if actualString != expectedString {
		t.Errorf("Expected String(%s) is not same as"+
			" actual string (%s)", expectedString, actualString)
	}
}
