package Solutions

import (
	"log"
	"strconv"
	"strings"
)

func Day1(data []byte) int {
	input := string(data)
	lines := strings.Split(strings.TrimSpace(input), "\n")

	placeholderValue := 50
	zeroCount := 0

	for _, value := range lines {

		direction := value[:1]
		distance, err := strconv.Atoi(value[1:])
		if err != nil {
			log.Fatal("These was a error converting the string")
		}

		switch direction {
		case "L":
			placeholderValue = (placeholderValue - distance) % 100

		case "R":
			placeholderValue = (placeholderValue + distance) % 100

		default:
			log.Fatal("invalid direction")
		}

		if placeholderValue < 0 {
			placeholderValue += 100
		}

		if placeholderValue == 0 {
			zeroCount++
		}

	}

	return zeroCount
}

func Day1Part2(data []byte) int {
	input := strings.TrimSpace(string(data))
	if input == "" {
		return 0
	}
	lines := strings.Split(input, "\n")

	pos := 50
	zeroCount := 0

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		dir := strings.ToUpper(line[:1])
		dist, err := strconv.Atoi(strings.TrimSpace(line[1:]))
		if err != nil {
			log.Fatalf("error parsing %q: %v", line, err)
		}

		zeroCount += hitsZero(pos, dir, dist)

		switch dir {
		case "R":
			pos = (pos + dist) % 100
		case "L":
			pos = (pos - dist) % 100
		default:
			log.Fatalf("invalid direction %q", dir)
		}

		if pos < 0 {
			pos += 100
		}
	}

	return zeroCount
}

func hitsZero(pos int, dir string, dist int) int {
	if dist <= 0 {
		return 0
	}

	pos %= 100
	if pos < 0 {
		pos += 100
	}

	var t0 int
	switch dir {
	case "R":
		t0 = (100 - pos) % 100
	case "L":
		t0 = pos
	default:
		log.Fatalf("invalid direction %q", dir)
	}

	if t0 == 0 {
		t0 = 100
	}

	if dist < t0 {
		return 0
	}
	return 1 + (dist-t0)/100
}
