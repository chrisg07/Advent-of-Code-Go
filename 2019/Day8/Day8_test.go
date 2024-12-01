package AoCScaffold

import (
	"bytes"
	"log"
	"os"
	"reflect"
	"strings"
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
	log.Print("[CONSOLE] Advent of Code 2019 Day 8:\n")
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

func TestPartAComplete(t *testing.T) {
	answer := 1806
	solution := PartA(false)
	if solution != answer {
		t.Fatalf(`Complete solution = %d, should = %d`, solution, answer)
	} else {
		log.Printf("[CONSOLE] The number of 1 digits multiplied by the number of 2 digits from the layer with the least amount of 0's: %v", solution)
	}
}

func TestPartBComplete(t *testing.T) {
	var str bytes.Buffer
	log.SetOutput(&str)
	expectedOutput := []string{
		"[CONSOLE] ▓▓░░▓▓░░▓▓░░░░▓░░░▓▓▓░░▓▓",
		"[CONSOLE] ▓▓▓░▓░▓▓░▓░▓▓▓▓░▓▓░▓░▓▓░▓",
		"[CONSOLE] ▓▓▓░▓░▓▓░▓░░░▓▓░▓▓░▓░▓▓░▓",
		"[CONSOLE] ▓▓▓░▓░░░░▓░▓▓▓▓░░░▓▓░░░░▓",
		"[CONSOLE] ░▓▓░▓░▓▓░▓░▓▓▓▓░▓░▓▓░▓▓░▓",
		"[CONSOLE] ▓░░▓▓░▓▓░▓░▓▓▓▓░▓▓░▓░▓▓░▓",
	}

	PartB(false)

	output := str.String()

	for _, line := range expectedOutput {
		if strings.Contains(output, line) == false {
			t.Fatal(`Captured output didn't include expected output\n`)
		}
	}
}
