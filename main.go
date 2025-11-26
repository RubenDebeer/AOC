package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// Refactor this code

func main() {

	err := run()
	fatalIfError(err, "There was a problem running the program:")

}

func run() error {
	const cookieName = "AOC_COOKIE_SESSION"
	godotenv.Load()
	cookie := os.Getenv(cookieName)

	request, rqerr := buildRequest("2024", "1", "https://adventofcode.com/", cookie)
	fatalIfError(rqerr, "There was a problem building the request:")

	// Initialize the client and use the request
	response, rserr := http.DefaultClient.Do(request)
	fatalIfError(rserr, "There was a problem making the request:")
	defer response.Body.Close()

	responseData, rerr := io.ReadAll(response.Body)
	fatalIfError(rerr, "There was a problem reading the response:")

	fmt.Println(string(responseData))
	return nil
}

// Function that builds the request
func buildRequest(year, day string, baseURL, cookie string) (*http.Request, error) {

	fullUrl := baseURL + year + "/day/" + day + "/input"

	request, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Add("Cookie", cookie)

	return request, nil
}

// Small helper function to log fatal errors
func fatalIfError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %v", msg, err)
	}
}
