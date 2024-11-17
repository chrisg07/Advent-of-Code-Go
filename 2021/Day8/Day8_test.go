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

// func TestDay8PartA2021Example(t *testing.T) {
// 	answer := 26
// 	solution := Day8PartA2021(true)
// 	if solution != answer {
// 		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
// 	}
// }

func TestDay8PartA2021Complete(t *testing.T) {
	answer := 321
	solution := Day8PartA2021(false)
	if solution != answer {
		t.Fatalf(`Comeplete solution = %d, should = %d`, solution, answer)
	}
}

func TestStringContainsOtherStringsChars(t *testing.T) {
	answer := true
	solution := StringContainsOtherStringsChars("cefabd", "eafb")
	if solution != answer {
		t.Fatalf(`Example solution = %t, should = %t`, solution, answer)
	}

	answer = false
	solution = StringContainsOtherStringsChars("true", "false")
	if solution != answer {
		t.Fatalf(`Example solution = %t, should = %t`, solution, answer)
	}
}

func TestDecodeLine(t *testing.T) {
	answer := 5353
	solution := DecodeLine("acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf")
	if solution != answer {
		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
	}
}

// func TestDay8PartB2021Example(t *testing.T) {
// 	answer := 61229
// 	solution := Day8PartB2021(true)
// 	if solution != answer {
// 		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
// 	}
// }

func TestDay8PartB2021Complete(t *testing.T) {
	answer := 1028926
	solution := Day8PartB2021(false)
	if solution != answer {
		t.Fatalf(`Complete solution = %d, should = %d`, solution, answer)
	}
}
