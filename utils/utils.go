package Utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
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
	total := 0
	for _, value := range arr {
		total += value
	}
	return total
}

func ArrayFromString(str string, delimiter string) []int {
	numbers := []int{}
	str = strings.TrimSpace(str)
	numberStrs := strings.Split(str, delimiter)
	for _, str := range numberStrs {
		value, err := strconv.Atoi(str)
		if err == nil {
			numbers = append(numbers, value)
		}
	}
	return numbers
}

func MultiplyArray(arr []int) int {
	total := 1
	for _, value := range arr {
		total *= value
	}
	return total
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

func SplitStringToInts(stringToSplit string, delimeter string) []int {
	splitStrings := []string{}
	if delimeter == " " {
		splitStrings = strings.Fields(stringToSplit)
	} else {
		splitStrings = strings.Split(stringToSplit, delimeter)
	}
	numbers := []int{}
	for _, str := range splitStrings {
		number, _ := strconv.Atoi(str)
		numbers = append(numbers, number)
	}
	return numbers
}

func ArrayContainsArray[T constraints.Ordered](haystack []T, needles []T) bool {
	for _, needle := range needles {
		if !slices.Contains(haystack, needle) {
			return false
		}
	}

	return true
}

func Transpose[T constraints.Ordered](slice [][]T) [][]T {
	xl := len(slice[0])
	yl := len(slice)
	result := make([][]T, xl)
	for i := range result {
		result[i] = make([]T, yl)
	}
	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i][j] = slice[j][i]
		}
	}
	return result
}

func PrettyPrint(v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		log.Println("[WARN] " + string(b))
	}
	return
}

func CopyMap(mapToCopy map[string]int) map[string]int {
	copy := make(map[string]int)
	for k, v := range mapToCopy {
		copy[k] = v
	}
	return copy
}

func ReverseString(str string) string {
	reverseStr := ""
	for i := len(str) - 1; i >= 0; i-- {
		reverseStr += string(str[i])
	}
	return reverseStr
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(integers []int) int {
	result := integers[0]

	for i := 0; i < len(integers); i++ {
		result = ((integers[i] * result) / (GCD(integers[i], result)))
	}

	return result
}

func RotateMatrix[T constraints.Ordered](X [][]T, n int) [][]T {
	for i := 0; i < n; i++ {
		for j := i; j < n-i-1; j++ {
			temp := X[i][j]
			X[i][j] = X[j][n-1-i]
			X[j][n-1-i] = X[n-1-i][n-1-j]
			X[n-1-i][n-1-j] = X[n-1-j][i]
			X[n-1-j][i] = temp
		}
	}
	return X
}
