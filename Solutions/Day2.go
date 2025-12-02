package Solutions

import (
	"log"
	"strconv"
	"strings"
)

func Day2(data []byte) int {

	input := string(data)
	line := strings.Split(input, "\n")
	ranges := strings.Split(line[0], ",")

	var invalidList []int
	for _, r := range ranges {
		invalidList = append(invalidList, getInvalidListFromRange(r)...)
	}

	var sum int
	for _, id := range invalidList {
		sum += id
	}

	return sum
}

func getInvalidListFromRange(idRanges string) []int {

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

// 3859 3859
func isValid(id string) bool {

	length := len(id)

	split := length / 2
	first := id[:split]
	second := id[split:]

	return first != second

}
