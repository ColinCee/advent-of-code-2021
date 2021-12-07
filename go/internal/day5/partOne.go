package day5

import (
	"advent-of-code-2021/internal"
	"fmt"
	"log"
	"strings"
)

type Board [][]int

func loadCoordinates(inputFile string) [][]int {
	lines, err := internal.ReadLines(inputFile)
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	// line example:
	// x1,y1 -> x2,y2

	coordinates := Board{}
	for _, line := range lines {
		// replace -> with nothing
		parsedLine := strings.Replace(line, "->", ",", 1)
		// remove whitespace
		parsedLine = strings.Replace(parsedLine, " ", "", -1)
		// split by comma
		stringCoords := strings.Split(parsedLine, ",")
		// convert array to []int
		coordinateLine := internal.StringArrayToIntArray(stringCoords)

		coordinates = append(coordinates, coordinateLine)
	}

	return coordinates
}

func initialiseBoard(coordinates [][]int) (Board, [][]int) {
	// calculate width and height of board
	width := 0
	height := 0
	for _, coordinates := range coordinates {
		if coordinates[0] > width {
			width = coordinates[0]
		}

		if coordinates[2] > width {
			width = coordinates[2]
		}

		if coordinates[1] > height {
			height = coordinates[1]
		}

		if coordinates[3] > height {
			height = coordinates[3]
		}
	}

	// print dimensions
	log.Printf("width: %d, height: %d", width, height)
	// initialise Board with zeros for width and height
	board := Board{}
	for i := 0; i < height+1; i++ {
		board = append(board, make([]int, width+1))
	}
	return board, coordinates
}

func processCoordinates(board Board, coordinates [][]int) {
	// for each coordinate x1,y1,x2,y2
	// increment the value of each point on the line between x1,y1 and x2,y2
	// ignore diagonal lines

	for _, coordinate := range coordinates {
		x1 := coordinate[0]
		y1 := coordinate[1]
		x2 := coordinate[2]
		y2 := coordinate[3]

		// if line is vertical
		if x1 == x2 {
			// figure out which direction to increment
			if y1 < y2 {
				// from y1 to y2
				for y := y1; y <= y2; y++ {
					board[y][x1]++
				}
			} else {
				// from y2 to y1
				for y := y1; y >= y2; y-- {
					board[y][x1]++
				}
			}
		}

		// if line is horizontal
		if y1 == y2 {
			// figure out which direction to increment
			if x1 < x2 {
				// from x1 to x2
				for x := x1; x <= x2; x++ {
					board[y1][x]++
				}
			} else {
				// from x2 to x1
				for x := x1; x >= x2; x-- {
					board[y1][x]++
				}
			}
		}

	}
}

// return number of points that have a value > 1
func calculateNumOverlaps(board Board) int {
	numOverlaps := 0
	for _, row := range board {
		for _, value := range row {
			if value > 1 {
				numOverlaps++
			}
		}
	}
	return numOverlaps
}

func SolvePartOne() {
	coordinates := loadCoordinates("internal/day5/real_input.txt")
	board, coordinates := initialiseBoard(coordinates)
	// internal.PrettyPrint2DIntArray(board)

	processCoordinates(board, coordinates)

	// internal.PrettyPrint2DIntArray(board)
	fmt.Println(calculateNumOverlaps(board))
}
