package AoC2021

import (
	_ "embed"
	"log"
	"slices"
	"strconv"
	"strings"
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

type Game struct {
	cardNumbers     []int
	winningNumbers  []int
	copies          int
	matchingNumbers int
}

func Day4PartA2023(useExample bool) int {
	lines := getInput(useExample)

	cards := buildState(lines)

	sum := 0
	// for each game determine win
	for index, card := range cards {
		points := 0
		for _, cardNumber := range card.cardNumbers {
			if slices.Contains(card.winningNumbers, cardNumber) && cardNumber != 0 {
				log.Printf("[WARN] Card %d has winning number %d", index, cardNumber)
				if points == 0 {
					points = 1
				} else {
					points *= 2
				}
			}
		}
		log.Printf("[WARN] Card %d won with %d points", index, points)
		sum += points
	}
	return sum
}

func buildState(lines []string) []Game {
	cards := []Game{}
	for _, line := range lines {
		parts := strings.Split(line, ": ")
		numbersStrs := strings.Split(parts[1], " | ")
		winningNumbersStr := strings.Split(numbersStrs[0], " ")
		winningNumbers := []int{}
		for _, str := range winningNumbersStr {
			str = strings.TrimSpace(str)
			value, _ := strconv.Atoi(str)
			winningNumbers = append(winningNumbers, value)
		}
		cardNumbersStr := strings.Split(numbersStrs[1], " ")
		cardNumbers := []int{}
		for _, str := range cardNumbersStr {
			str = strings.TrimSpace(str)
			value, _ := strconv.Atoi(str)
			cardNumbers = append(cardNumbers, value)
		}
		card := Game{
			winningNumbers:  winningNumbers,
			cardNumbers:     cardNumbers,
			copies:          1,
			matchingNumbers: 0,
		}
		cards = append(cards, card)
	}
	return cards
}

func Day4PartB2023(useExample bool) int {
	lines := getInput(useExample)

	cards := buildState(lines)

	for index, card := range cards {
		copiesToAdd := 0
		for _, cardNumber := range card.cardNumbers {
			if slices.Contains(card.winningNumbers, cardNumber) && cardNumber != 0 {
				cards[index].matchingNumbers += 1
			}
			copiesToAdd += cards[index].matchingNumbers
		}
	}

	for i := 0; i < len(cards); i++ {
		for numWins := 0; numWins < cards[i].matchingNumbers; numWins++ {
			if i+1+numWins < len(cards) {
				cards[i+1+numWins].copies += cards[i].copies
			}
		}
	}
	sum := 0

	for index, card := range cards {
		log.Printf("[WARN] Card %d has %d copies and won %d times\n", index+1, card.copies, card.matchingNumbers)
		sum += card.copies
	}
	return sum
}
