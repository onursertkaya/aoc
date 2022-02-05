package main

import (
	"fmt"
	"lib/util"
)

// file -> package
// dir with go.mod -> module
// only uppercase symbols are exported
// filenames in a module has no effect
// file locations in a module is important

func countParentheses(arr string) (ctr int, first_char_pos_to_basement int) {

	for i := 0; i < len(arr); i++ {
		ch := string(arr[i])

		if ch == "(" {
			ctr++
		} else if ch == ")" {
			ctr--
		}

		if (first_char_pos_to_basement == 0) && (ctr == -1) {
			first_char_pos_to_basement = i + 1
		}

		if ch == "\n" {
			break
		}
	}

	return
}

func main() {
	puzzle := util.GetPuzzlePath()
	data := util.ReadPuzzleInput(puzzle)
	ctr, first_char_pos_to_basement := countParentheses(data)

	fmt.Println(ctr)
	fmt.Println(first_char_pos_to_basement)
}
