package AoCScaffold

import (
	"log"
	"os"
	"reflect"
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
	log.Print("[CONSOLE] Advent of Code YEAR Day ##:\n")
	log.Print("[CONSOLE] --------------------------\n")
}

func TestPartAParseInput(t *testing.T) {
	expected := [][][]int{{{1, 2, 3}, {4, 5, 6}}, {{7, 8, 9}, {0, 1, 2}}}
	actual := ConvertImageDataToLayers(3, 2, "123456789012")
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf(`Example solution = %d, should = %d`, actual, expected)
	}
}

func TestCountZeroesInLayer(t *testing.T) {
	expected := 0
	layer := [][]int{{1, 2, 3}, {4, 5, 6}}
	actual := CountIntInLayer(layer, 0)
	if actual != expected {
		t.Fatalf(`Example solution = %d, should = %d`, actual, expected)
	}
}
func TestFindLayerWithFewestZeroes(t *testing.T) {
	expected := [][]int{{1, 2, 3}, {4, 5, 6}}
	layers := ConvertImageDataToLayers(3, 2, "123456789012")

	actual := FindLayerWithFewestZeroes(layers)
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf(`Example solution = %d, should = %d`, actual, expected)
	}
}

func TestPartAExample(t *testing.T) {
	answer := 1
	solution := PartA(true)
	if solution != answer {
		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
	}
}

func TestPartAComplete(t *testing.T) {
	answer := 1806
	solution := PartA(false)
	if solution != answer {
		t.Fatalf(`Complete solution = %d, should = %d`, solution, answer)
	} else {
		log.Printf("[CONSOLE] The number of 1 digits multiplied by the number of 2 digits from the layer with the least amount of 0's: %v", solution)
	}
}

func TestPartBExample(t *testing.T) {
	answer := 1
	solution := PartB(true)
	if solution != answer {
		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
	}
}

func TestPartBComplete(t *testing.T) {
	answer := 1
	solution := PartB(false)
	if solution != answer {
		t.Fatalf(`Complete solution = %d, should = %d`, solution, answer)
	} else {
		log.Printf("[CONSOLE] Find the Elf carrying the most Calories: %v", solution)
	}
}
