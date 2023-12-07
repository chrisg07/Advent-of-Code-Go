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

var cardsPartA = []string{"2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A"}
var cardsPartB = []string{"J", "2", "3", "4", "5", "6", "7", "8", "9", "T", "Q", "K", "A"}

func determineCardRankPartA(card string) int {
	return slices.Index(cardsPartA, card)
}

func determineCardRankPartB(card string) int {
	return slices.Index(cardsPartB, card)
}

type PokerHand struct {
	hand         string
	cardValues   []int
	fiveOfAKind  bool
	fourOfAKind  bool
	fullHouse    bool
	threeOfAKind bool
	twoPair      bool
	onePair      bool
	highCard     bool // all labels distinct
	bet          int
	numJokers    int
}

func countOccurences(haystack string, needle string) int {
	return strings.Count(haystack, needle)
}

func isNOfAKind(hand PokerHand, n int) bool {
	for _, card := range cardsPartA {
		if countOccurences(hand.hand, card) == n-hand.numJokers {
			return true
		}
	}
	return false
}

func countPairs(hand PokerHand) int {
	pairs := 0
	for _, card := range cardsPartA {
		if countOccurences(hand.hand, card) == 2 {
			pairs += 1
		}
	}
	return pairs
}

func createPokerHandPartA(hand string, bet int) PokerHand {
	pokerHand := PokerHand{
		hand: hand,
		bet:  bet,
	}
	for _, char := range hand {
		cardValue := determineCardRankPartA(string(char))
		pokerHand.cardValues = append(pokerHand.cardValues, cardValue)
	}
	pokerHand.fiveOfAKind = isNOfAKind(pokerHand, 5)
	pokerHand.fourOfAKind = isNOfAKind(pokerHand, 4)
	pokerHand.threeOfAKind = isNOfAKind(pokerHand, 3)
	pokerHand.fullHouse = isNOfAKind(pokerHand, 3) && isNOfAKind(pokerHand, 2)
	pokerHand.twoPair = countPairs(pokerHand) == 2
	pokerHand.onePair = countPairs(pokerHand) == 1
	return pokerHand
}

func createPokerHandPartB(hand string, bet int) PokerHand {
	pokerHand := PokerHand{
		hand: hand,
		bet:  bet,
	}
	for _, char := range hand {
		cardValue := determineCardRankPartA(string(char))
		pokerHand.cardValues = append(pokerHand.cardValues, cardValue)
	}
	pokerHand.numJokers = countOccurences(pokerHand.hand, "J")
	pokerHand.fiveOfAKind = isNOfAKind(pokerHand, 5)
	pokerHand.fourOfAKind = isNOfAKind(pokerHand, 4)
	pokerHand.threeOfAKind = isNOfAKind(pokerHand, 3)
	pokerHand.fullHouse = isNOfAKind(pokerHand, 3) && isNOfAKind(pokerHand, 2)
	pokerHand.twoPair = countPairs(pokerHand) == 2
	pokerHand.onePair = countPairs(pokerHand) == 1
	return pokerHand
}

func compareHandsCardsInOrder(a PokerHand, b PokerHand) int {
	for index := range a.cardValues {
		if a.cardValues[index] != b.cardValues[index] {
			return a.cardValues[index] - b.cardValues[index]
		}
	}
	return 0
}

func compareHands(a PokerHand, b PokerHand) int {
	// if five of a kind
	if a.fiveOfAKind && !b.fiveOfAKind {
		return 1
	} else if !a.fiveOfAKind && b.fiveOfAKind {
		return -1
	} else if a.fiveOfAKind && b.fiveOfAKind {
		return compareHandsCardsInOrder(a, b)
	}
	// if four of a kind
	if a.fourOfAKind && !b.fourOfAKind {
		return 1
	} else if !a.fourOfAKind && b.fourOfAKind {
		return -1
	} else if a.fourOfAKind && b.fourOfAKind {
		return compareHandsCardsInOrder(a, b)
	}
	// if full house
	if a.fullHouse && !b.fullHouse {
		return 1
	} else if !a.fullHouse && b.fullHouse {
		return -1
	} else if a.fullHouse && b.fullHouse {
		return compareHandsCardsInOrder(a, b)
	}
	// if three of a kind
	if a.threeOfAKind && !b.threeOfAKind {
		return 1
	} else if !a.threeOfAKind && b.threeOfAKind {
		return -1
	} else if a.threeOfAKind && b.threeOfAKind {
		return compareHandsCardsInOrder(a, b)
	}
	// if two pair
	if a.twoPair && !b.twoPair {
		return 1
	} else if !a.twoPair && b.twoPair {
		return -1
	} else if a.twoPair && b.twoPair {
		return compareHandsCardsInOrder(a, b)
	}
	// if one pair
	if a.onePair && !b.onePair {
		return 1
	} else if !a.onePair && b.onePair {
		return -1
	} else if a.onePair && b.onePair {
		return compareHandsCardsInOrder(a, b)
	}
	// if high card
	return compareHandsCardsInOrder(a, b)
}

func Day7PartA2023(useExample bool) int {
	lines := getInput(useExample)
	hands := []PokerHand{}
	for _, line := range lines {
		parts := strings.Split(line, " ")
		bet, _ := strconv.Atoi(parts[1])
		hand := createPokerHandPartA(parts[0], bet)
		sortedPosition := findPositionPartA(hands, hand)
		hands = slices.Insert(hands, sortedPosition, hand)
	}

	log.Printf("[WARN] hands: %v\n", hands)

	totalWinnings := 0
	for index, hand := range hands {
		totalWinnings += (index + 1) * hand.bet
	}
	return totalWinnings
}

func findPositionPartA(hands []PokerHand, hand PokerHand) int {
	for index := range hands {
		if compareHands(hand, hands[index]) < 0 {
			return index
		}
	}
	return len(hands)
}

func Day7PartB2023(useExample bool) int {
	lines := getInput(useExample)
	hands := []PokerHand{}
	for _, line := range lines {
		parts := strings.Split(line, " ")
		bet, _ := strconv.Atoi(parts[1])
		hand := createPokerHandPartB(parts[0], bet)
		sortedPosition := findPositionPartA(hands, hand)
		hands = slices.Insert(hands, sortedPosition, hand)
	}

	log.Printf("[WARN] Cards: %v\n", cardsPartA)

	totalWinnings := 0
	for index, hand := range hands {
		totalWinnings += (index + 1) * hand.bet
	}
	return totalWinnings
}
