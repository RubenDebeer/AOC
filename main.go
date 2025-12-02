package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/RubenDebeer/AOC/Solutions"
	"github.com/joho/godotenv"
)

func main() {
	// Day 1
	// data, err := getData("1")
	// FatalIfError(err, "There was a error in the get Data method ")
	//data := []byte("L68\nL30\nR48\nL5\nR60\nL55\nL1\nL99\nR14\nL82\n")

	// result := Solutions.Day1Part2(data)
	// log.Printf("Day 1 result: %d", result)

	data, err := getData("2")
	FatalIfError(err, "There was a error in the get Data method ")

	result := Solutions.Day2(data)
	log.Printf("Day 2 result: %d", result)
}

func getData(day string) ([]byte, error) {
	const cookieName = "AOC_COOKIE_SESSION"
	enverror := godotenv.Load()
	FatalIfError(enverror, "There was a problem loading the environment variables:")

	cookie := os.Getenv(cookieName)

	request, rqerr := BuildRequest("2025", day, "https://adventofcode.com/", cookie)
	FatalIfError(rqerr, "There was a problem building the request:")

	response, rserr := http.DefaultClient.Do(request)
	FatalIfError(rserr, "There was a problem making the request:")
	defer response.Body.Close()

	responseData, rerr := io.ReadAll(response.Body)
	FatalIfError(rerr, "There was a problem reading the response:")

	return responseData, nil
}

// BuildRequest Function that builds the request
func BuildRequest(year, day string, baseURL, cookie string) (*http.Request, error) {

	fullUrl := baseURL + year + "/day/" + day + "/input"

	request, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Add("Cookie", cookie)

	return request, nil
}

// FatalIfError Small helper function to log fatal errors
func FatalIfError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %v", msg, err)
	}
}
