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

	input := "67562556-67743658,62064792-62301480,4394592-4512674,3308-4582,69552998-69828126,9123-12332,1095-1358,23-48,294-400,3511416-3689352,1007333-1150296,2929221721-2929361280,309711-443410,2131524-2335082,81867-97148,9574291560-9574498524,648635477-648670391,1-18,5735-8423,58-72,538-812,698652479-698760276,727833-843820,15609927-15646018,1491-1766,53435-76187,196475-300384,852101-903928,73-97,1894-2622,58406664-58466933,6767640219-6767697605,523453-569572,7979723815-7979848548,149-216"

	b := []byte(input) // <-- array/slice of bytes of the string
	_ = b

	// data, err := getData("2")
	// FatalIfError(err, "There was a error in the get Data method ")

	result := Solutions.Day2(b)
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
