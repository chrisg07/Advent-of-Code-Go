package AoC2021

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/chrisg07/Advent-of-Code-Go/Utils"
)

//go:embed inputs/example.txt
var exampleInput string

//go:embed inputs/complete.txt
var input string

func getInput(useExample bool) []string {
	var lines []string
	var unsplitLines string
	if useExample {
		unsplitLines = strings.TrimRight(exampleInput, "\n")
	} else {
		unsplitLines = strings.TrimRight(input, "\n")
	}
	lines = strings.Split(unsplitLines, "\n")
	return lines
}

type BingoBoard struct {
	rows      [][]int
	columns   [][]int
	markedWon bool
}

func CheckBoardWinCondition(board BingoBoard, draws []int) bool {
	for _, row := range board.rows {
		if Utils.ArrayContainsArray[int](draws, row) {
			fmt.Printf("Winning row: %d\n", row)
			return true
		}
	}
	for _, column := range board.columns {
		if Utils.ArrayContainsArray[int](draws, column) {
			fmt.Printf("Winning column: %d\n", column)
			return true
		}
	}
	return false
}

func SumOfUnmarkedNumbers(board BingoBoard, draws []int) int {
	sumOfUnmarkedCells := 0
	for _, row := range board.rows {
		for _, cell := range row {
			if !slices.Contains(draws, cell) {
				sumOfUnmarkedCells += cell
			}
		}
	}
	fmt.Printf("Sum of unmarked cells: %d\n", sumOfUnmarkedCells)
	return sumOfUnmarkedCells
}

func Day4PartA2021(useExample bool) int {
	lines := getInput(useExample)

	draws := []int{}
	boards := []BingoBoard{}
	for index, line := range lines {
		if index == 0 {
			drawOrder := strings.Split(line, ",")

			for _, draw := range drawOrder {
				number, _ := strconv.Atoi(draw)
				draws = append(draws, number)
			}
		} else if index == 2 || (index%6)-2 == 0 {
			rows := [][]int{
				Utils.SplitStringToInts(lines[index], " "),
				Utils.SplitStringToInts(lines[index+1], " "),
				Utils.SplitStringToInts(lines[index+2], " "),
				Utils.SplitStringToInts(lines[index+3], " "),
				Utils.SplitStringToInts(lines[index+4], " "),
			}
			columns := [][]int{
				{rows[0][0], rows[1][0], rows[2][0], rows[3][0], rows[4][0]},
				{rows[0][1], rows[1][1], rows[2][1], rows[3][1], rows[4][1]},
				{rows[0][2], rows[1][2], rows[2][2], rows[3][2], rows[4][2]},
				{rows[0][3], rows[1][3], rows[2][3], rows[3][3], rows[4][3]},
				{rows[0][4], rows[1][4], rows[2][4], rows[3][4], rows[4][4]},
			}

			board := BingoBoard{rows, columns, false}
			boards = append(boards, board)
		}
	}

	boardHasWon := false
	winningBoard := BingoBoard{}
	winningDraws := []int{}
	for index := range draws {
		currentDraws := draws[:index]
		for _, board := range boards {
			boardHasWon = CheckBoardWinCondition(board, currentDraws)
			if boardHasWon {
				winningBoard = board
				winningDraws = currentDraws
				break
			}
		}

		if boardHasWon {
			break
		}
	}

	sumOfUnmarkedCells := SumOfUnmarkedNumbers(winningBoard, winningDraws)
	return sumOfUnmarkedCells * winningDraws[len(winningDraws)-1]
}

func Day4PartB2021(useExample bool) int {
	lines := getInput(useExample)

	draws := []int{}
	boards := []BingoBoard{}
	for index, line := range lines {
		if index == 0 {
			drawOrder := strings.Split(line, ",")

			for _, draw := range drawOrder {
				number, _ := strconv.Atoi(draw)
				draws = append(draws, number)
			}
		} else if index == 2 || (index%6)-2 == 0 {
			rows := [][]int{
				Utils.SplitStringToInts(lines[index], " "),
				Utils.SplitStringToInts(lines[index+1], " "),
				Utils.SplitStringToInts(lines[index+2], " "),
				Utils.SplitStringToInts(lines[index+3], " "),
				Utils.SplitStringToInts(lines[index+4], " "),
			}
			columns := [][]int{
				{rows[0][0], rows[1][0], rows[2][0], rows[3][0], rows[4][0]},
				{rows[0][1], rows[1][1], rows[2][1], rows[3][1], rows[4][1]},
				{rows[0][2], rows[1][2], rows[2][2], rows[3][2], rows[4][2]},
				{rows[0][3], rows[1][3], rows[2][3], rows[3][3], rows[4][3]},
				{rows[0][4], rows[1][4], rows[2][4], rows[3][4], rows[4][4]},
			}

			board := BingoBoard{rows, columns, false}
			boards = append(boards, board)
		}
	}

	boardHasWon := false
	numWinningBoards := 0
	winningBoard := BingoBoard{}
	winningDraws := []int{}
	for index := range draws {
		currentDraws := draws[:index+1]
		for boardIndex, board := range boards {
			if !board.markedWon {
				boardHasWon = CheckBoardWinCondition(board, currentDraws)
				if boardHasWon && !board.markedWon && numWinningBoards == (len(boards)-1) {
					boards[boardIndex].markedWon = true
					numWinningBoards += 1
					winningBoard = board
					winningDraws = currentDraws
					fmt.Printf("Final board won. # of winning boards: %d\nBoard: %v\nDraws: %v\n", numWinningBoards, board, currentDraws)
					break
				} else if boardHasWon && !board.markedWon {
					boards[boardIndex].markedWon = true
					numWinningBoards += 1
					fmt.Printf("Board won. # of winning boards: %d\nBoard: %v\nDraws: %v\n", numWinningBoards, board, currentDraws)
				}
			}

		}

		if len(winningDraws) > 0 {
			fmt.Println("Winning board found. Breaking loop")
			break
		}
	}

	sumOfUnmarkedCells := SumOfUnmarkedNumbers(winningBoard, winningDraws)
	return sumOfUnmarkedCells * winningDraws[len(winningDraws)-1]
}
