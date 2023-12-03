package AoC2021

import (
	_ "embed"
	"log"
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

type Point struct {
	x int
	y int
}

type EngineNumber struct {
	start Point
	end   Point
	value int
}

func buildState(lines []string) ([]EngineNumber, []Point) {
	engineNumbers := []EngineNumber{}
	symbols := []Point{}
	for x, line := range lines {
		number := 0
		start := Point{0, 0}
		end := Point{0, 0}
		for y, char := range line {
			parsedInt, parseErr := strconv.Atoi(string(char))
			if parseErr == nil && number == 0 {
				log.Printf("[WARN] Parse error is nil and value is 0\n")
				if number == 0 {
					start = Point{x: x, y: y}
				}
				number *= 10
				number += parsedInt
			} else if parseErr == nil && number > 0 {
				log.Printf("[WARN] Parse error is nil and value is: %d\n", parsedInt)
				number *= 10
				number += parsedInt
			} else if parseErr != nil && number > 0 {
				log.Printf("[WARN] Parse error is not nil and value is: %d\n", parsedInt)
				end = Point{x: x, y: y - 1}
				engineNumber := EngineNumber{
					start: start,
					end:   end,
					value: number,
				}
				engineNumbers = append(engineNumbers, engineNumber)
				number = 0
			}
			if parseErr != nil && string(char) != "." {
				symbol := Point{x: x, y: y}
				symbols = append(symbols, symbol)
			}
		}
		if number > 0 {
			end = Point{x: x, y: len(line) - 1}
			engineNumber := EngineNumber{
				start: start,
				end:   end,
				value: number,
			}
			engineNumbers = append(engineNumbers, engineNumber)
			number = 0
		}
	}
	return engineNumbers, symbols
}

func symbolIsAdjacentToNumber(symbol Point, engineNumber EngineNumber) bool {
	withinRangeX := symbol.x >= engineNumber.start.x-1 && symbol.x <= engineNumber.end.x+1
	withingRangeY := symbol.y >= engineNumber.start.y-1 && symbol.y <= engineNumber.end.y+1
	return withinRangeX && withingRangeY
}

func Day3PartA2023(useExample bool) int {
	lines := getInput(useExample)

	engineNumbers, symbols := buildState(lines)

	partNumbers := []EngineNumber{}
	for _, engineNumber := range engineNumbers {
		for _, symbol := range symbols {
			if symbolIsAdjacentToNumber(symbol, engineNumber) {
				partNumbers = append(partNumbers, engineNumber)
				break
			}
		}
	}

	sum := 0
	for _, partNumber := range partNumbers {
		sum += partNumber.value
	}
	return sum
}

func Day3PartB2023(useExample bool) int {
	lines := getInput(useExample)

	engineNumbers, symbols := buildState(lines)

	gearRatios := []int{}
	for _, symbol := range symbols {
		numNumbersAdjacent := 0
		gearRatio := 0
		for _, engineNumber := range engineNumbers {
			if symbolIsAdjacentToNumber(symbol, engineNumber) {
				numNumbersAdjacent += 1
				if gearRatio == 0 {
					gearRatio += engineNumber.value
				} else {
					gearRatio *= engineNumber.value
				}
			}
		}
		if numNumbersAdjacent == 2 {
			gearRatios = append(gearRatios, gearRatio)
		}
	}

	return Utils.SumArray(gearRatios)
}
