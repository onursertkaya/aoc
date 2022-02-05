package util

import (
	"errors"
	"os"
	"path"
	"runtime"
)

func ReadPuzzleInput(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func GetPuzzlePath() (puzzle_path string) {
	_, caller_file, _, _ := runtime.Caller(1)
	puzzle_path = path.Join(path.Dir(caller_file), "input.txt")

	_, err := os.Stat(puzzle_path)
	if errors.Is(err, os.ErrNotExist) {
		panic("Could not find puzzle.")
	}

	return
}
