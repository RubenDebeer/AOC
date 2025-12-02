package main

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"

	"github.com/RubenDebeer/AOC/src"
)

// Refactor this code
func main() {
	responseData, err := src.Run()

	if err != nil {
		panic(err)
	}

	Day1(responseData)
}

func Day1(inputs []byte) int {
	//fmt.Println(string(inputs))
	//fmt.Println("test this " + string(inputs[5]))

	// Parse Input
	//left := []int{}
	//right := []int{}

	scanner := bufio.NewScanner(bytes.NewReader(inputs))

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		fields := strings.Fields(line)
		if len(fields) != 2 {
			panic("Invalid input")
		}

		fmt.Println("Test The line :" + fields[0])
	}

	return 0
}
