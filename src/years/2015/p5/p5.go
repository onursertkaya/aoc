package main

import (
	"fmt"
	"src/lib/util"
	"regexp"
	"strings"
)

func anyPartIsInLine(line *string, parts *[]string) (skip bool) {
	skip = false
	for _, b := range *parts {
		if strings.Contains(*line, b) {
			skip = true
			break
		}
	}
	return
}

func repetitivePairOfLetters(line *string) (repetitive bool) {
	repetitive = false
	for k, l := range *line {
		if k < 1 {
			continue
		}
		rest_of_the_line := string((*line)[k+1:])
		part := string((*line)[k-1]) + string(l)
		if anyPartIsInLine(&rest_of_the_line,  &[]string{part}) {
			repetitive = true
			break
		}
	}
	return
}

func repetitiveChars(line *string, spacing int) (bool) {
	for k, l := range *line {
		if k < spacing {
			continue
		}

		if l == rune((*line)[k-spacing]) {
			return true
		}
	}
	return false
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

	part1_nice := 0
	part2_nice := 0

	lines := strings.Split(data, "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}

		contains_forbidden := anyPartIsInLine(&line, &[]string{"ab", "cd", "pq", "xy"})
		repetitive := repetitiveChars(&line, 1)
		vowels := AtLeastThreeVowels(&line)
		if !contains_forbidden && repetitive && vowels {
			part1_nice++
		}
		repeats_in_between := repetitiveChars(&line, 2)
		repeats_a_pair := repetitivePairOfLetters(&line)
		if repeats_in_between && repeats_a_pair {
			part2_nice++
		}

	}
	fmt.Println(part1_nice)
	fmt.Println(part2_nice)
}
