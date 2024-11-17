package AoC2021

import (
	"log"
	"os"
	"testing"

	"github.com/hashicorp/logutils"
)

func init() {
	filter := &logutils.LevelFilter{
		Levels:   []logutils.LogLevel{"DEBUG", "WARN", "ERROR"},
		MinLevel: logutils.LogLevel("ERROR"),
		Writer:   os.Stderr,
	}
	log.SetOutput(filter)
}

func TestDay7PartA2023Example(t *testing.T) {
	answer := 6440
	solution := Day7PartA2023(true)
	if solution != answer {
		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
	}
}

func TestDay7PartA2023Complete(t *testing.T) {
	answer := 252295678
	solution := Day7PartA2023(false)
	if solution != answer {
		t.Fatalf(`Comeplete solution = %d, should = %d`, solution, answer)
	}
}

func TestDay7PartB2023Example(t *testing.T) {
	answer := 5905
	solution := Day7PartB2023(true)
	if solution != answer {
		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
	}
}

func TestCompareHands(t *testing.T) {
	answer := true

	// five of a kind against five of a kind with a joker
	a := CreatePokerHandPartB("AAAAA", 0)
	b := CreatePokerHandPartB("AAAAJ", 0)
	solution := CompareHands(a, b) > 0
	if solution != answer {
		t.Fatalf(`Complete solution = %t, should = %t`, solution, answer)
	}

	// five of a kind against five of a kind with a joker
	a = CreatePokerHandPartB("JAAAA", 0)
	b = CreatePokerHandPartB("KKKKK", 0)
	solution = CompareHands(a, b) < 0
	if solution != answer {
		t.Fatalf(`Complete solution = %t, should = %t`, solution, answer)
	}

	// true four of a kind against five of a kind with a joker
	a = CreatePokerHandPartB("2AAAA", 0)
	b = CreatePokerHandPartB("KKKKJ", 0)
	solution = CompareHands(a, b) < 0
	if solution != answer {
		t.Fatalf(`Complete solution = %t, should = %t`, solution, answer)
	}

	// full house with a joker against a true four of a kind
	a = CreatePokerHandPartB("AAKKJ", 0)
	b = CreatePokerHandPartB("KKKK2", 0)
	solution = CompareHands(a, b) < 0
	if solution != answer {
		t.Fatalf(`Complete solution = %t, should = %t`, solution, answer)
	}

	// full house with a joker against a true four of a kind
	a = CreatePokerHandPartB("K3333", 0)
	b = CreatePokerHandPartB("A2222", 0)
	solution = CompareHands(a, b) < 0
	if solution != answer {
		t.Fatalf(`Complete solution = %t, should = %t`, solution, answer)
	}

	a = CreatePokerHandPartB("JKKK2", 0)
	b = CreatePokerHandPartB("QQQQ2", 0)
	solution = CompareHands(a, b) < 0
	if solution != answer {
		t.Fatalf(`Complete solution = %t, should = %t`, solution, answer)
	}
}

func TestDay7PartB2023Complete(t *testing.T) {
	answer := 250577259
	solution := Day7PartB2023(false)
	if solution <= 248366140 {
		t.Fatalf(`The answer is too low`)
	}
	if solution == 248274675 || solution == 248469178 || solution == 248366724 || solution == 248914224 {
		t.Fatalf(`This is the same as a previous answer`)
	}
	if solution != answer {
		t.Fatalf(`Complete solution = %d, should = %d`, solution, answer)
	}
}
