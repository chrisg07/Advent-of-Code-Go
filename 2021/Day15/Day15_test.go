package AoC2021

import (
	"log"
	"os"

	"github.com/hashicorp/logutils"
)

func init() {
	filter := &logutils.LevelFilter{
		Levels:   []logutils.LogLevel{"DEBUG", "WARN", "ERROR"},
		MinLevel: logutils.LogLevel("WARN"),
		Writer:   os.Stderr,
	}
	log.SetOutput(filter)
}

// func TestDay15PartA2021Example(t *testing.T) {
// 	answer := 40
// 	solution := Day15PartA2021(true)
// 	if solution != answer {
// 		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
// 	}
// }

// func TestDay15PartA2021Complete(t *testing.T) {
// 	answer := 40
// 	solution := Day15PartA2021(false)
// 	if solution != answer {
// 		t.Fatalf(`Comeplete solution = %d, should = %d`, solution, answer)
// 	}
// }

// func TestDayXPartB2021Example(t *testing.T) {
// 	answer := 0
// 	solution := DayXPartB2021(true)
// 	if solution != answer {
// 		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
// 	}
// }

// func TestDayXPartB2021Complete(t *testing.T) {
// 	answer := 0
// 	solution := DayXPartB2021(false)
// 	if solution != answer {
// 		t.Fatalf(`Complete solution = %d, should = %d`, solution, answer)
// 	}
// }
