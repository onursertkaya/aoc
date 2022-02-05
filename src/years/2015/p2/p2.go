package main

import (
	"fmt"
	"lib/util"
	"strconv"
	"strings"
)

type Dim struct {
	x int
	y int
	z int
}

func minOfThree(a, b, c int) int {
	m := a
	if b < m {
		m = b
	}

	if c < m {
		m = c
	}

	return m
}

func (d Dim) area() int {
	return 2 * ((d.y+d.z)*d.x + d.y*d.z)
}

func (d Dim) extra() int {
	a := d.x * d.y
	b := d.x * d.z
	c := d.y * d.z

	return minOfThree(a, b, c)

}

func (d Dim) ribbon() int {
	a := d.x + d.y
	b := d.x + d.z
	c := d.y + d.z

	return 2*minOfThree(a, b, c) + d.x*d.y*d.z
}

func strToInt(s string) (out int) {
	out, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return
}

func main() {
	puzzle := util.GetPuzzlePath()
	data := util.ReadPuzzleInput(puzzle)

	lines := strings.Split(data, "\n")
	total_wrap := 0
	total_ribbon := 0
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if line == "" {
			break
		}
		fields := strings.Split(line, "x")
		d := Dim{strToInt(fields[0]), strToInt(fields[1]), strToInt(fields[2])}
		this_wrap := d.area() + d.extra()
		total_wrap += this_wrap

		total_ribbon += d.ribbon()
	}
	fmt.Println(total_wrap)
	fmt.Println(total_ribbon)
}
