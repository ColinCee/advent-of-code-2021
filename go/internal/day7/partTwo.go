package day7

import (
	"advent-of-code-2021/internal"
	"fmt"
	"math"
	"strings"
)

func sumOfOneToN(n int) int {
	return (n * (n + 1)) / 2
}

// for a list of numbers find the minimum difference when each number is compared
func solvePartTwo(numbers []int) int {

	min := math.MaxInt64

	// count number of times each number occurs
	counts := make(map[int]int)
	for _, num := range numbers {
		counts[num]++
	}

	// for each number in count, find the minimum difference between all other numbers
	for num, _ := range counts {
		totalDiff := 0

		for otherNum, otherCount := range counts {
			if num == otherNum {
				continue
			}
			diff := internal.AbsInt(num - otherNum)

			totalDiff += sumOfOneToN(diff) * otherCount
		}

		if totalDiff < min {
			min = totalDiff
		}
	}

	return min

}

func SolvePartTwo() {
	lines, err := internal.ReadLines("internal/day7/real_input.txt")
	if err != nil {
		panic(err)
	}

	// split lines[0] by comma
	stringNums := strings.Split(lines[0], ",")
	// convert string to int
	numbers := internal.StringArrayToIntArray(stringNums)
	result := solvePartTwo(numbers)

	fmt.Println("Part Two: ", result)
}
