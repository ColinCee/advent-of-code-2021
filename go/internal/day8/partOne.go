package day8

import (
	"advent-of-code-2021/internal"
	"fmt"
	"strings"
)

func countUniqueSegmentsInOutput(lines []string) int {
	total := 0
	for _, line := range lines {
		inputOutput := strings.Split(line, "|")
		outputSegments := inputOutput[1]
		outputSegments = strings.TrimSpace(outputSegments)

		for _, segment := range strings.Split(outputSegments, " ") {
			// segement values 1,4,7 and 8 are the only ones that are unique
			// therefore if segment length is 2,3,4 or 7 incremenet total by 1
			if len(segment) == 2 || len(segment) == 3 || len(segment) == 4 || len(segment) == 7 {
				total++
			}
		}

	}

	return total
}
func SolvePartOne() {
	lines, err := internal.ReadLines("internal/day8/real_input.txt")
	if err != nil {
		panic(err)
	}

	total := countUniqueSegmentsInOutput(lines)
	fmt.Println("total in output: ", total)
}
