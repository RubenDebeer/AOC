package Solutions

import (
	"fmt"
	"strings"
)

func Day1(data []byte) int {
	input := string(data)
	lines := strings.Split(strings.TrimSpace(input), "\n")

	fmt.Println(lines)
	return 2
}
