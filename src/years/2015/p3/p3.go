package main

import (
	"fmt"
	"lib/util"
)

type Pos struct {
	x int
	y int
}

type Set map[Pos]bool

func updatePosition(curr_pos *Pos, c rune) {
	switch c {

	case '<':
		*curr_pos = Pos{curr_pos.x - 1, curr_pos.y}
	case '>':
		*curr_pos = Pos{curr_pos.x + 1, curr_pos.y}
	case 'v':
		*curr_pos = Pos{curr_pos.x, curr_pos.y + 1}
	case '^':
		*curr_pos = Pos{curr_pos.x, curr_pos.y - 1}
	default:
		panic(nil)
	}
}

func part1(data *string) {
	curr_pos := Pos{}
	positions := Set{}

	positions[curr_pos] = true

	for _, c := range *data {
		updatePosition(&curr_pos, c)
		positions[curr_pos] = true
	}
	fmt.Println(len(positions))
}

func part2(data *string) {
	curr_pos_santa := Pos{}
	curr_pos_robot := Pos{}

	positions_santa := Set{}
	positions_robot := Set{}

	positions_santa[curr_pos_santa] = true
	positions_robot[curr_pos_santa] = true

	for i, c := range *data {
		if i%2 == 0 {
			updatePosition(&curr_pos_santa, c)
			positions_santa[curr_pos_santa] = true
		} else {
			updatePosition(&curr_pos_robot, c)
			positions_robot[curr_pos_robot] = true
		}

	}

	final := Set{}

	for k, v := range positions_santa {
		final[k] = v
	}

	for k, v := range positions_robot {
		final[k] = v
	}
	fmt.Println(len(final))

}

func main() {
	puzzle := util.GetPuzzlePath()
	data := util.ReadPuzzleInput(puzzle)

	part1(&data)

	part2(&data)
}
