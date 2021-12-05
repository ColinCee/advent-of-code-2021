package day4

import (
	"advent-of-code-2021/internal"
	"fmt"
	"log"
)

func playLosingBingo(numbers []int, boards Boards) (Boards, int, int) {
	// create set of numbers 0 to length of boards
	remainingBoards := make(map[int]int)
	for i := 0; i < len(boards); i++ {
		remainingBoards[i] = i
	}

	for _, number := range numbers {
		// convert string to int

		for boardIndex, board := range boards {
			// if boardIndex is not in remainingBoards, skip
			if _, ok := remainingBoards[boardIndex]; !ok {
				continue
			}

			// This in N^2
			for rowIndex, row := range board {
				for columnIndex, column := range row {
					if column.number == number {
						// We can't use board created by range as it seems to be a copy of the original board

						boards[boardIndex][rowIndex][columnIndex].checked = true
						// if row or column is fully checked then return
						if isRowFullyChecked(boards[boardIndex], rowIndex) ||
							isColumnFullyChecked(boards[boardIndex], columnIndex) {
							fmt.Println("bingo! board index:", boardIndex)

							prettyPrintBoards(boards)
							// if remainingBoards has only one element return
							if len(remainingBoards) == 1 {
								// get first value of remainingBoards
								for _, index := range remainingBoards {
									return boards, index, number
								}
							}
							delete(remainingBoards, boardIndex)
						}
					}
				}
			}
		}
	}

	panic("This shouldn't happen")
}

func SolvePartTwo() {
	lines, err := internal.ReadLines("internal/day4/real_input.txt")

	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	bingoBoards := loadBingoBoards(lines)
	prettyPrintBoards(bingoBoards)

	numbers := loadNumbers(lines)
	boards, losingBoardIndex, finalNumber := playLosingBingo(numbers, bingoBoards)
	// prettyPrintSingleBoard( boards[losingBoardIndex])
	fmt.Println("losing board index:", losingBoardIndex, "final number", finalNumber)
	finalScore := calculateScore(boards[losingBoardIndex], finalNumber)
	fmt.Println("final score:", finalScore)
}
