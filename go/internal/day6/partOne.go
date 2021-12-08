package day6

import (
	"advent-of-code-2021/internal"
	"fmt"
	"strings"
)

func loadInitialState(filepath string) []int {
	lines, err := internal.ReadLines(filepath)

	if err != nil {
		panic(err)
	}

	stringState := strings.Split(lines[0], ",")
	return internal.StringArrayToIntArray(stringState)
}

func simulateDays(days int, initialState []int) []int {
	// initialise intialState with 3,4,3,1,2
	var fishes = initialState

	// for each day
	// follow the following rules
	// decrease the value of each fish by 1
	// if fish reaches 0, reset it to 6 and append a new fish with value of 8

	for i := 0; i < days; i++ {
		// declare array to store new fishes for each day
		var newFishes []int

		// for each fish
		for j := 0; j < len(fishes); j++ {
			// decrease the value of each fish by 1
			fishes[j] = fishes[j] - 1
			// if fish is 0, reset it to 6 and append a new fish with value of 8
			if fishes[j] == -1 {
				fishes[j] = 6
				newFishes = append(newFishes, 8)
			}
		}

		fishes = append(fishes, newFishes...)
		fmt.Println(i+1, "| num", len(fishes), fishes)
		// fmt.Println("day", i+1, "fishes", len(fishes))
	}

	return fishes
}

func SolvePartOne() {
	days := 20

	initialState := []int{1}
	// initialState := loadInitialState("internal/day6/real_input.txt")
	fishes := simulateDays(days, initialState)
	// print number of fishes
	fmt.Println("number of fishes:", len(fishes))
}
