package Utils

import (
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/exp/constraints"
)

func ReadAoCInput(file string) []string {
	content, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")

	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-2]
	}

	return lines
}

func FormatMessage(message string) {
	fmt.Println(message)
}

func SumArray(arr []int) int {
	sum := 0
	for _, value := range arr {
		sum += value
	}
	return sum
}

func ReverseArray[T constraints.Ordered](arr []T) []T {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}

	return arr
}

func ToPower(number int, exponent int) int {
	if exponent == 0 {
		return 1
	}
	answer := number
	for i := 1; i < exponent; i += 1 {
		answer *= number
	}
	return answer
}

func BitArrayToDecimal(arr []int) int {
	decimalRepresentation := 0

	reversedBitArray := ReverseArray(arr)

	for index, value := range reversedBitArray {
		if value == 1 {
			amountToAdd := ToPower(2, index)
			decimalRepresentation += amountToAdd
		}
	}

	ReverseArray(arr)

	return decimalRepresentation
}

func FlipBitArray(arr []int) []int {
	flippedBitArray := []int{}

	for index, value := range arr {
		flippedBitArray = append(flippedBitArray, 0)
		if value == 1 {
			flippedBitArray[index] = 0
		} else {
			flippedBitArray[index] = 1
		}
	}

	return flippedBitArray
}

func RemoveIndex[T constraints.Ordered](arr []T, index int) []T {
	return append(arr[:index], arr[index+1:]...)
}

func DuplicateMatrix[T constraints.Ordered](matrix [][]T) [][]T {
	duplicate := make([][]T, len(matrix))
	for i := range matrix {
		duplicate[i] = make([]T, len(matrix[i]))
		copy(duplicate[i], matrix[i])
	}
	return duplicate
}
