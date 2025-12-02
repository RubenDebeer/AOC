package Solutions

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

// [Range 1] [Range 2] [Range 3]  ....

// [11-22]   [95-115]  [998-1012] ....

// Process Range

func Day2(data []byte) {

	// // 1)
	// // Read Data into a DS
	// // Get all of the RangesConv
	// // Split Line into a slice of Ranges

	input := string(data)
	line := strings.Split(strings.TrimSpace(input), "\n")
	ranges := strings.Split(line[0], ",")
	var invalidIds []int

	for _, r := range ranges {

		problemIds := buildRanges(r)
		invalidIds = appendProblemIds(problemIds)
		fmt.Println(invalidIds)
	}

	// 2)
	// Iterate through the Range
	// Identify the Invalid ID's and add it to a list
}

func appendProblemIds(invalidIds []int) []int {
	var ids []int
	for _, value := range invalidIds {
		ids = append(ids, value)
	}

	return ids
}

func buildRanges(idRanges string) []int {

	var rangeVal []int
	split := strings.Split(idRanges, "-")

	start, err := strconv.Atoi(split[0])
	if err != nil {
		log.Fatal("there was a error converting from string to int")
	}

	end, err := strconv.Atoi(split[1])
	if err != nil {
		log.Fatal("there was a error converting from string to int")
	}

	for i := start; i <= end; i++ {
		if isValid(strconv.Itoa(i)) {
			continue
		} else {
			rangeVal = append(rangeVal, i)
		}
	}

	return rangeVal
}

func isValid(id string) bool {

	length := len(id)

	split := length / 2

	first := id[:split]
	second := id[split:]

	return first != second
}
