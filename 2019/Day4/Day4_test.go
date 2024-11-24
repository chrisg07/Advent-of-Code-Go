package AoCScaffold

import (
	"log"
	"os"
	"testing"

	"github.com/hashicorp/logutils"
)

func init() {
	filter := &logutils.LevelFilter{
		Levels:   []logutils.LogLevel{"DEBUG", "WARN", "ERROR", "CONSOLE"},
		MinLevel: logutils.LogLevel("WARN"),
		Writer:   os.Stderr,
	}
	log.SetFlags(0)
	log.SetOutput(filter)
	log.Print("[CONSOLE] --------------------------\n")
	log.Print("[CONSOLE] Advent of Code 2019 Day 4:\n")
	log.Print("[CONSOLE] --------------------------\n")
}

// func TestPartAExample(t *testing.T) {
// 	answer := 278384
// 	solution := PartA(true)
// 	if solution != answer {
// 		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
// 	}
// }

func TestPasswordWithoutAdjacentDuplicates(t *testing.T) {
	solution := PasswordHasAdjacentDuplicates(123456)
	if solution != false {
		t.Fatal(`Expected false`)
	}
}

func TestPasswordWithAdjacentDuplicates(t *testing.T) {
	solution := PasswordHasAdjacentDuplicates(113456)
	if solution != true {
		t.Fatal(`Expected true`)
	}
}

func TestPasswordWithDecreasingDigits(t *testing.T) {
	solution := PasswordHasNoDecreasingDigits(223450)
	if solution != true {
		t.Fatal(`Expected true`)
	}
}

func TestPasswordWithAllSameDigits(t *testing.T) {
	solution := PasswordHasNoDecreasingDigits(111111)
	if solution != true {
		t.Fatal(`Expected true`)
	}
}

func TestPasswordWithNoDecreasingDigits(t *testing.T) {
	solution := PasswordHasNoDecreasingDigits(12345)
	if solution != true {
		t.Fatal(`Expected true`)
	}
}

func TestPotentialPasswords(t *testing.T) {
	password := 122345
	actual := IsPotentialPassword(password)
	if actual != true {
		t.Fatal(`Expected true`)
	}
	password = 223450
	actual = IsPotentialPassword(password)
	if actual != false {
		t.Fatal(`Expected false`)
	}
	password = 123789
	actual = IsPotentialPassword(password)
	if actual != false {
		t.Fatal(`Expected false`)
	}
}

func TestPartAComplete(t *testing.T) {
	answer := 921
	solution := PartA(false)
	if solution != answer {
		t.Fatalf(`Complete solution = %d, should = %d`, solution, answer)
	} else {
		log.Printf("[CONSOLE] There are %v different passwords within the provided range.", solution)
	}
}

func TestPotentialPasswordsPartB(t *testing.T) {
	password := 278889
	actual := IsPotentialPasswordPartB(password)
	if actual != false {
		t.Fatal(`Expected false`)
	}
	password = 278888
	actual = IsPotentialPasswordPartB(password)
	if actual != false {
		t.Fatal(`Expected false`)
	}
}

func TestPartBComplete(t *testing.T) {
	answer := 603
	solution := PartB(false)
	if solution != answer {
		t.Fatalf(`Complete solution = %d, should = %d`, solution, answer)
	} else {
		log.Printf("[CONSOLE] There are %v different passwords within the provided range that meet all the criteria.", solution)
	}
}
