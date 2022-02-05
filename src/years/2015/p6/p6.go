package main

import (
	"lib/util"
	"strings"
)

func main() {
	puzzle := util.GetPuzzlePath()
	data := util.ReadPuzzleInput(puzzle)

	lines := strings.Split(data, "\n")

	for _, line := range lines {
		if line == "" {
			continue
		}

		line
	}
}
