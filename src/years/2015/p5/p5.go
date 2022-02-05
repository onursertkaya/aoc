package main

import (
	"fmt"
	"lib/util"
	"regexp"
	"strings"
)

func forbiddenSubstring(line *string) (skip bool) {
	skip = false
	for _, b := range []string{"ab", "cd", "pq", "xy"} {
		if strings.Contains(*line, b) {
			skip = true
			break
		}
	}
	return
}

func repetitiveChars(line *string) (repetitive bool) {
	repetitive = false
	for k, l := range *line {
		if k == 0 {
			continue
		}

		if l == rune((*line)[k-1]) {
			repetitive = true
			break
		}
	}
	return
}

func AtLeastThreeVowels(line *string) bool {
	re := regexp.MustCompile(`[aeiou]`)
	matches := re.FindAllIndex([]byte(*line), -1)
	if matches == nil {
		return false
	}

	return (len(matches) >= 3)
}

func main() {
	// https://go.dev/blog/strings

	puzzle := util.GetPuzzlePath()
	data := util.ReadPuzzleInput(puzzle)

	nice := 0

	lines := strings.Split(data, "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}

		skip := forbiddenSubstring(&line)
		if skip {
			continue
		}

		repetitive := repetitiveChars(&line)
		if !repetitive {
			continue
		}

		vowels := AtLeastThreeVowels(&line)
		if !vowels {
			continue
		}
		nice++
	}
	fmt.Println(nice)
}
