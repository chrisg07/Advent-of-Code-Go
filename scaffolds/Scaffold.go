package AoC2021

import (
	"fmt"

	_ "embed"
	"strings"

	"github.com/chrisg07/Advent-of-Code-Go/utils"
)

//go:embed inputs/example.txt
var exampleInput string

//go:embed inputs/complete.txt
var input string

func getInput(useExample bool) []string {
	var lines []string
	if useExample {
		exampleInput = strings.TrimRight(exampleInput, "\n")
		lines = strings.Split(exampleInput, "\n")
	} else {
		input = strings.TrimRight(input, "\n")
		lines = strings.Split(input, "\n")
	}
	return lines
}
func DayNumberPartLetter() {
	utils.FormatMessage("Insert message here")

	lines := utils.ReadAoCInput("./2021/inputs/Day1/PartAExample.txt")

	for _, line := range lines {
		fmt.Println(line)
	}
}
