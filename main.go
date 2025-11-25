package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Build the Request
	request, err := newRequestBuilder().AddURL("1", "2024").Build()
	if err != nil {
		log.Fatalf("there was an error %v", err)
	}

	//_ = godotenv.Load()
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading .env file: %v", err)
	}

	request.Header.Add("Cookie", os.Getenv("AOC_COOKIE_SESSION"))

	// Initialize the Client
	res, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Fatalf("there was an error making the call %v", err)
	}
	defer res.Body.Close()

	responseBody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("there was an error reading the response %v", err)
	}

	log.Printf("response: %s", responseBody)
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
	}

	rb.request = request
	return request, nil
}
