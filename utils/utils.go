package utils

import (
	"fmt"
	"log"
	"os"
	"strings"
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
