package day10

import (
	"advent-of-code-2021/internal"
	"fmt"
	"sort"
)

func getIncompleteStacks(lines []string) [][]rune {
	incompleteStacks := make([][]rune, 0)

	for _, line := range lines {
		isLegal, _, stack := isLineLegal(line)
		if isLegal {
			incompleteStacks = append(incompleteStacks, stack)

			internal.PrintRuneArray(stack)
		}
	}

	return incompleteStacks
}

func calculatePointsForStack(stack []rune) int {
	points := 0
	charPointMap := map[rune]int{
		'(': 1,
		'[': 2,
		'{': 3,
		'<': 4,
	}

	// for each char in reverse order
	for i := len(stack) - 1; i >= 0; i-- {
		char := stack[i]
		points *= 5
		points += charPointMap[char]
	}

	return points
}

func cacluatePointsPartTwo(stacks [][]rune) int {
	points := []int{}

	for _, stack := range stacks {
		points = append(points, calculatePointsForStack(stack))
	}
	// find middle point
	middle := len(points) / 2
	// sort points
	sort.Ints(points)
	internal.PrettyPrintIntArray(points)

	// return middle point
	return points[middle]
}

func SolvePartTwo() {
	lines := readInput("internal/day10/real_input.txt")
	stacks := getIncompleteStacks(lines)
	points := cacluatePointsPartTwo(stacks)
	fmt.Println("Points: ", points)
}
