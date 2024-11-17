package AoC2021

import (
	_ "embed"
	"log"
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

func determineDifferences(array []int) []int {
	differences := []int{}
	for i := 0; i < len(array)-1; i++ {
		difference := array[i+1] - array[i]
		differences = append(differences, difference)
	}
	return differences
}

func containsAllZeroes(array []int) bool {
	isZero := true
	for _, value := range array {
		if value != 0 {
			return false
		}
	}
	return isZero
}

func Day92023(useExample bool, reverse bool) int {
	lines := getInput(useExample)
	histories := [][]int{}
	for _, line := range lines {
		array := Utils.ArrayFromString(line, " ")
		if reverse {
			array = Utils.ReverseArray[int](array)
		}
		histories = append(histories, array)
	}

	nextValuesInHistory := []int{}
	for _, history := range histories {
		differences := [][]int{history}
		isAllZeroes := containsAllZeroes(differences[0])
		for !isAllZeroes {
			historyDifferences := determineDifferences(differences[len(differences)-1])
			isAllZeroes = containsAllZeroes(historyDifferences)
			differences = append(differences, historyDifferences)
		}

		for i := len(differences) - 1; i > 0; i-- {
			extrapolatedValue := differences[i-1][len(differences[i-1])-1] + differences[i][len(differences[i])-1]
			differences[i-1] = append(differences[i-1], extrapolatedValue)
		}

		nextValuesInHistory = append(nextValuesInHistory, differences[0][len(differences[0])-1])
	}
	return Utils.SumArray(nextValuesInHistory)
}

func Day9PartB2023(useExample bool) int {
	lines := getInput(useExample)
	for _, line := range lines {
		for _, char := range line {
			log.Print(string(char))
		}
		log.Println("")
	}

	return 0
}
