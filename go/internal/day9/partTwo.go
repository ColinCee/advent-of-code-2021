package day9

import (
	"advent-of-code-2021/internal"
	"fmt"
	"sort"
	"strconv"
)

type Node struct {
	row int
	col int
	val int
}

func getNeighbors(input [][]int, row int, col int) []Node {
	neighbors := []Node{}

	// North neighbor
	if row > 0 {
		neighbors = append(neighbors, Node{row: row - 1, col: col, val: input[row-1][col]})
	}

	// South neighbor
	if row < len(input)-1 {
		neighbors = append(neighbors, Node{row: row + 1, col: col, val: input[row+1][col]})
	}

	// East neighbor
	if col > 0 {
		neighbors = append(neighbors, Node{row: row, col: col - 1, val: input[row][col-1]})
	}

	// West neighbor
	if col < len(input[0])-1 {
		neighbors = append(neighbors, Node{row: row, col: col + 1, val: input[row][col+1]})
	}
	return neighbors
}

// starting from row and col
// perform depth first search
// add number to the stack if it is greater than the current number but not 9
// keep track of the number of elements we've considered
func findBasinSize(input [][]int, row int, col int) (int, map[Node]bool) {
	fmt.Printf("\nStarting at (%d, %d)\n", row, col)
	// create stack of nodes
	node := Node{row, col, input[row][col]}
	stack := []Node{node}
	// create visited set
	visited := make(map[Node]bool)
	counter := 0
	for len(stack) > 0 {
		// pop the top element
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// check if it is in the visited set
		if _, ok := visited[node]; ok {
			continue
		} else {
			visited[node] = true
			counter += 1
		}

		// for each neighbor
		for _, neighbor := range getNeighbors(input, node.row, node.col) {
			if neighbor.val > node.val && neighbor.val != 9 {
				stack = append(stack, neighbor)
			}
		}
	}
	// debugBasin(input, visited)
	return counter, visited
}

func debugBasin(input [][]int, visitedNodes map[Node]bool) {
	debugBasin := copyInputShapeReplaceWithX(input)

	for node, _ := range visitedNodes {
		debugBasin[node.row][node.col] = strconv.Itoa(node.val)
	}

	internal.PrettyPrint2DStringArray(debugBasin)
}

func calculateSizeOfThreeLargestBasins(input [][]int, lowPoints [][]int) int {

	sizes := []int{}

	for _, lowPoint := range lowPoints {
		row := lowPoint[0]
		col := lowPoint[1]
		size, _ := findBasinSize(input, row, col)
		sizes = append(sizes, size)
		fmt.Printf("(%d, %d) has size %d\n", row, col, size)
	}

	// sort sizes
	sort.Ints(sizes)
	// multiply top three sizes and return it
	sum := 1
	for _, num := range sizes[len(sizes)-3:] {
		sum *= num
	}

	return sum
}

func SolvePartTwo() {
	input := readInput("internal/day9/real_input.txt")
	internal.PrettyPrint2DIntArray(input)
	lowPoints := findLowPoints(input)
	sum := calculateSizeOfThreeLargestBasins(input, lowPoints)

	fmt.Printf("The sum of the sizes of the three largest basins is %d\n", sum)
}
