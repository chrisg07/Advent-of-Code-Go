package AoCScaffold

import (
	_ "embed"
	"strconv"
	"strings"
)

//go:embed inputs/example.txt
var exampleInput string

//go:embed inputs/complete.txt
var input string

func ConvertImageDataToLayers(width int, height int, imageData string) [][][]int {
	layers := [][][]int{}

	rows := [][]int{}
	for i := 0; i < len(imageData); i += width {
		row := []int{}

		substring := imageData[i:(i + width)]
		for _, char := range substring {
			pixelValue, _ := strconv.Atoi(string(char))
			row = append(row, pixelValue)
		}

		rows = append(rows, row)
	}

	for y := 0; y < len(rows); y += height {
		layer := rows[y : y+height]
		layers = append(layers, layer)
	}

	return layers
}

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

func CountIntInLayer(layer [][]int, needle int) int {
	occurrences := 0
	for _, row := range layer {
		for _, pixel := range row {
			if pixel == needle {
				occurrences++
			}
		}
	}
	return occurrences
}

func FindLayerWithFewestZeroes(layers [][][]int) [][]int {
	minZeroes := 10000000
	layerWithFewestZeroes := layers[0]

	for _, layer := range layers {
		zeroes := CountIntInLayer(layer, 0)

		if zeroes < minZeroes {
			minZeroes = zeroes
			layerWithFewestZeroes = layer
		}
	}

	return layerWithFewestZeroes
}

func parseInput(lines []string) []string {
	input := []string{}
	for _, line := range lines {
		input = append(input, line)
	}
	return input
}

func PartA(useExample bool) int {
	lines := getInput(useExample)
	input := parseInput(lines)
	layers := ConvertImageDataToLayers(25, 6, input[0])
	layerWithFewestZeros := FindLayerWithFewestZeroes(layers)

	numOnes := CountIntInLayer(layerWithFewestZeros, 1)
	numTwos := CountIntInLayer(layerWithFewestZeros, 2)

	return numOnes * numTwos
}

func PartB(useExample bool) int {
	lines := getInput(useExample)
	input := parseInput(lines)

	return len(input)
}
