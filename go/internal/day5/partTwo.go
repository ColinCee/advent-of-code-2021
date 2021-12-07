package day5

import (
	"fmt"
)

func processCoordinatesWithDiagonals(board Board, coordinates [][]int) {
	// for each coordinate x1,y1,x2,y2
	// increment the value of each point on the line between x1,y1 and x2,y2
	// ignore non 45 degree lines

	for _, coordinate := range coordinates {
		x1 := coordinate[0]
		y1 := coordinate[1]
		x2 := coordinate[2]
		y2 := coordinate[3]

		// print coordinates
		fmt.Printf("%d,%d,%d,%d\n", x1, y1, x2, y2)

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
		} else if y1 == y2 { // else if line is horizontal
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
		} else if (y2-y1)/(x2-x1) == 1 { // else if gradient is 1
			// figure out which direction to increment
			if x1 < x2 {
				// from x1 to x2
				for i := 0; i <= x2-x1; i++ {
					board[y1+i][x1+i]++
				}
			} else {
				// from x2 to x1
				for i := 0; i <= x1-x2; i++ {
					board[y1-i][x1-i]++
				}
			}
		} else if (y2-y1)/(x2-x1) == -1 { // else if gradient is -1
			// figure out which direction to increment
			if x1 < x2 {
				// we go from bottom right to top left
				for i := 0; i <= x2-x1; i++ {
					board[y1-i][x1+i]++
				}
			} else {
				// we go from top left to bottom right
				for i := 0; i <= x1-x2; i++ {
					board[y1+i][x1-i]++
				}
			}
		}

	}
}

func SolvePartTwo() {
	coordinates := loadCoordinates("internal/day5/real_input.txt")
	board, coordinates := initialiseBoard(coordinates)
	// internal.PrettyPrint2DIntArray(board)

	processCoordinatesWithDiagonals(board, coordinates)

	// internal.PrettyPrint2DIntArray(board)
	fmt.Println(calculateNumOverlaps(board))
}
