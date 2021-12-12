package day10

import (
	"advent-of-code-2021/internal"
	"fmt"
)

func readInput(path string) []string {
	lines, err := internal.ReadLines(path)
	if err != nil {
		panic(err)
	}

	return lines
}

func isOpeningParenthesis(char rune) bool {
	return char == '(' || char == '[' || char == '{' || char == '<'
}

func getComplementOpenParenthesis(char rune) rune {
	switch char {
	case ')':
		return '('
	case ']':
		return '['
	case '}':
		return '{'
	case '>':
		return '<'
	}
	return ' '
}

func isLineLegal(line string) (bool, rune, []rune) {
	stack := make([]rune, 0)

	for _, c := range line {
		if isOpeningParenthesis(c) {
			stack = append(stack, c)
			continue
		}
		// pop char
		if len(stack) == 0 {
			return false, c, stack
		}

		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if top != getComplementOpenParenthesis(c) {
			return false, c, stack
		}
	}

	return true, ' ', stack
}

func calculatePoints(lines []string) int {
	points := 0

	charPointMap := map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}

	for _, line := range lines {
		isLegal, illegalChar, _ := isLineLegal(line)
		if !isLegal {
			points += charPointMap[illegalChar]
		}
	}
	return points
}

func SolvePartOne() {
	lines := readInput("internal/day10/real_input.txt")
	points := calculatePoints(lines)
	fmt.Println("Part One:", points)
}
