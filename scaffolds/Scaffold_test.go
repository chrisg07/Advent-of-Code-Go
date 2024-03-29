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

// func TestDayXPartA2021Example(t *testing.T) {
// 	answer := 0
// 	solution := DayXPartA2021(true)
// 	if solution != answer {
// 		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
// 	}
// }

// func TestDayXPartA2021Complete(t *testing.T) {
// 	answer := 0
// 	solution := DayXPartA2021(false)
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
