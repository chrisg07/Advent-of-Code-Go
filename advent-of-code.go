package main

import (
	"log"
	"os"

	AoC2021 "github.com/chrisg07/Advent-of-Code-Go/2021"
	"github.com/hashicorp/logutils"
)

func main() {
	filter := &logutils.LevelFilter{
		Levels:   []logutils.LogLevel{"DEBUG", "WARN", "ERROR"},
		MinLevel: logutils.LogLevel("WARN"),
		Writer:   os.Stderr,
	}
	log.SetOutput(filter)

	AoC2021.Solutions()
}
