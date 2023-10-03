package main

import (
	"fmt"
	"regexp"
	"src/lib/util"
	"strconv"
	"strings"
)

type Op int

const (
	Off    Op = 0
	On     Op = 1
	Toggle Op = 2

	GridSize = 1000
)

type UpdateGrid struct {
	xmin int
	ymin int
	xmax int
	ymax int
	op   Op
}

type Grid struct {
	data [GridSize][GridSize]int
}

func (g *Grid) update(update_grid UpdateGrid) {
	for j := update_grid.ymin; j <= update_grid.ymax; j++ {
		for i := update_grid.xmin; i <= update_grid.xmax; i++ {
			if update_grid.op == Toggle {
				current := g.data[j][i]
				if current == int(On) {
					g.data[j][i] = int(Off)
				} else if current == int(Off) {
					g.data[j][i] = int(On)
				}
			} else {
				g.data[j][i] = int(update_grid.op)
			}
		}
	}
}

func (g Grid) sum() (ctr int) {
	ctr = 0
	for j := 0; j < GridSize; j++ {
		for i := 0; i < GridSize; i++ {
			ctr = ctr + g.data[j][i]
		}
	}
	return
}

func parseInt(number string) int {
	p, e := strconv.ParseInt(number, 10, 32)
	if e != nil {
		panic("not a number")
	}
	return int(p)
}

func parse(line *string) UpdateGrid {
	beginning := (*line)[:7]
	var op Op
	if beginning[len(beginning)-1] == "f"[0] {
		op = Off
	} else if beginning[len(beginning)-1] == "n"[0] {
		op = On
	} else {
		op = Toggle
	}

	pattern := regexp.MustCompile(`.*\s([0-9]+),([0-9]+)\s.*\s([0-9]+),([0-9]+)$`)
	match := pattern.FindAllStringSubmatch(*line, -1)[0]
	if match == nil {
		panic("pattern didnt match")
	}
	p := parseInt
	return UpdateGrid{p(match[1]), p(match[2]), p(match[3]), p(match[4]), op}
}

func main() {
	puzzle := util.GetPuzzlePath()
	data := util.ReadPuzzleInput(puzzle)

	lines := strings.Split(data, "\n")

	main_grid := Grid{}

	for _, line := range lines {
		if line == "" {
			continue
		}

		update_grid := parse(&line)
		main_grid.update(update_grid)
	}
	fmt.Println(main_grid.sum())
}
