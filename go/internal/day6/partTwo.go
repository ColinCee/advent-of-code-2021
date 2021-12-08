package day6

import "fmt"

func SolvePartTwo() {
	numDays := 256
	initialState := loadInitialState("internal/day6/real_input.txt")

	// initialise array with 9 zeros
	var state = make([]int, 9)
	for i := 0; i < 9; i++ {
		state[i] = 0
	}
	fmt.Println(state)

	// set initial state
	for i := 0; i < len(initialState); i++ {
		state[initialState[i]] += 1
	}
	fmt.Println(state)

	for i := 0; i < numDays; i++ {
		// rotate state one element to the left
		firstElement := state[0]
		state = append(state[1:], firstElement)
		state[6] += state[8]
	}

	// sum state
	sum := 0
	for i := 0; i < len(state); i++ {
		sum += state[i]
	}
	fmt.Println(state)
	fmt.Println(sum)
}
