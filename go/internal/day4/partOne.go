package day4

import (
	"advent-of-code-2021/internal"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Board [5][5]struct {
	number  int
	checked bool
}

type Boards []Board

func prettyPrintBoards(boards Boards) {
	for _, board := range boards {
		prettyPrintSingleBoard(board)
		fmt.Println("")
	}
}

func prettyPrintSingleBoard(board Board) {
	for _, row := range board {
		fmt.Println(row)
	}
}

func loadBingoBoards(lines []string) Boards {
	var bingoBoards Boards

	boardIndex := 0
	rowIndex := 0

	for i, line := range lines {
		// skip first two lines
		if i < 2 {
			continue
		}

		// if board at boardIndex is undefined, create a new board
		if len(bingoBoards) <= boardIndex {
			bingoBoards = append(bingoBoards, Board{})
		}

		// split strings by space, ignore duplicate spaces
		rowNumbers := strings.Fields(line)
		// fmt.Println("boardIndex", boardIndex)
		// fmt.Println("rowIndex", boardIndex)

		for columnIndex, number := range rowNumbers {
			// fmt.Println("columnIndex", boardIndex, number)

			// convert string to int
			numberToInt, _ := strconv.Atoi(number)

			bingoBoards[boardIndex][rowIndex][columnIndex] = struct {
				number  int
				checked bool
			}{
				number:  numberToInt,
				checked: false,
			}
		}
		rowIndex += 1

		if line == "" {
			// increment board index
			boardIndex += 1
			rowIndex = 0
		}
	}

	return bingoBoards
}

func loadNumbers(lines []string) []int {
	numberStrings := strings.Split(lines[0], ",")
	// convert numberStrings to integer array
	numbers := make([]int, len(numberStrings))
	for i, v := range numberStrings {
		number, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		numbers[i] = number
	}

	return numbers
}

func isRowFullyChecked(board Board, rowIndex int) bool {
	for _, column := range board[rowIndex] {
		if !column.checked {
			return false
		}
	}
	return true
}

func isColumnFullyChecked(board Board, columnIndex int) bool {
	for _, row := range board {
		if !row[columnIndex].checked {
			return false
		}
	}
	return true
}

// Returns when a board wins, with
// all boards state, the winning board index and the number that was used to win
func playBingo(numbers []int, boards Boards) (Boards, int, int) {

	for _, number := range numbers {
		// convert string to int

		for boardIndex, board := range boards {
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
							return boards, boardIndex, number
						}
					}
				}
			}
		}
	}

	// throw error
	panic("no winning board found")
}

// return the sum of all unchecked numbers in the board multiplied by the final number
func calculateScore(board Board, finalNumber int) int {
	sum := 0
	for _, row := range board {
		for _, column := range row {
			if !column.checked {
				sum += column.number
			}
		}
	}
	return sum * finalNumber
}
func SolvePartOne() {
	lines, err := internal.ReadLines("internal/day4/real_input.txt")

	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	bingoBoards := loadBingoBoards(lines)
	prettyPrintBoards(bingoBoards)

	numbers := loadNumbers(lines)
	boards, winningBoardIndex, finalNumber := playBingo(numbers, bingoBoards)
	prettyPrintSingleBoard(boards[winningBoardIndex])

	finalScore := calculateScore(boards[winningBoardIndex], finalNumber)
	fmt.Println("final score:", finalScore)
}
