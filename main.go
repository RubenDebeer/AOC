package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Get ENV Variables

	request, err := newRequestBuilder().AddURL("1", "2024").Build()

	if err != nil {
		log.Fatalf("there was an error %v", err)
	}

	_ = godotenv.Load()
	request.Header.Add("Cookie", os.Getenv("AOC_COOKIE_SESSION"))
	fmt.Print(request)
}

type RequestBuilder struct {
	year    string
	day     string
	url     string
	request *http.Request
	cookie  string
}

func newRequestBuilder() *RequestBuilder {
	return &RequestBuilder{}
}

func (rb *RequestBuilder) AddURL(day, year string) *RequestBuilder {
	rb.day = day
	rb.year = year
	rb.url = "https://adventofcode.com/" + year + "/day/" + day + "/input"
	return rb
}

func (rb *RequestBuilder) Build() (*http.Request, error) {

	request, err := http.NewRequest("GET", rb.url, nil)

	if err != nil {
		log.Fatalf("there was an error %v", err)
		return nil, err
	}

	rb.request = request
	return request, nil
}
