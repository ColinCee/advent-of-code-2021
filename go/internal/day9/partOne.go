package day9

import (
	"advent-of-code-2021/internal"
	"fmt"
	"strconv"
)

func readInput(path string) [][]int {
	lines, err := internal.ReadLines(path)
	if err != nil {
		panic(err)
	}

	input := make([][]int, len(lines))

	for rowIndex, line := range lines {
		input[rowIndex] = make([]int, len(line))
		for colIndex, num := range line {
			// convert rune to int
			input[rowIndex][colIndex], err = strconv.Atoi(string(num))
			if err != nil {
				panic(err)
			}
		}
	}

	return input
}

func doAnyAdjacentCellsHaveSmallerInt(input [][]int, rowIndex int, colIndex int) bool {
	// check north
	if rowIndex > 0 {
		if input[rowIndex-1][colIndex] < input[rowIndex][colIndex] {
			return true
		}
	}

	// check south
	if rowIndex < len(input)-1 {
		if input[rowIndex+1][colIndex] < input[rowIndex][colIndex] {
			return true
		}
	}

	// check west
	if colIndex > 0 {
		if input[rowIndex][colIndex-1] < input[rowIndex][colIndex] {
			return true
		}
	}

	// check east
	if colIndex < len(input[rowIndex])-1 {
		if input[rowIndex][colIndex+1] < input[rowIndex][colIndex] {
			return true
		}
	}

	return false
}

func copyInputShapeReplaceWithX(input [][]int) [][]string {
	// copy input size and initialise with X
	copyInput := make([][]string, len(input))
	for rowIndex, _ := range input {
		copyInput[rowIndex] = make([]string, len(input[rowIndex]))
		for colIndex, _ := range input[rowIndex] {
			copyInput[rowIndex][colIndex] = "X"
		}
	}
	return copyInput
}

func debugLowPoints(input [][]int, lowPoints [][]int) {

	debugLowPoints := copyInputShapeReplaceWithX(input)
	for _, lowPoint := range lowPoints {
		row := lowPoint[0]
		col := lowPoint[1]
		debugLowPoints[row][col] = strconv.Itoa(input[row][col])
	}

	internal.PrettyPrint2DStringArray(debugLowPoints)
}

func findLowPoints(input [][]int) [][]int {
	// copy input
	lowPoints := [][]int{}

	for rowIndex, row := range input {
		for colIndex, _ := range row {
			if !doAnyAdjacentCellsHaveSmallerInt(input, rowIndex, colIndex) {
				lowPoints = append(lowPoints, []int{rowIndex, colIndex})
			}
		}
	}
	return lowPoints
}

func calculateRisk(input [][]int, lowPoints [][]int) int {
	risk := 0
	for _, lowPoint := range lowPoints {
		row := lowPoint[0]
		col := lowPoint[1]
		risk += input[row][col] + 1
	}

	return risk
}

func SolvePartOne() {
	input := readInput("internal/day9/test_input.txt")
	internal.PrettyPrint2DIntArray(input)
	lowPoints := findLowPoints(input)
	debugLowPoints(input, lowPoints)
	risk := calculateRisk(input, lowPoints)

	fmt.Println("Risk:", risk)
}
