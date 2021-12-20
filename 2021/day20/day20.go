package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Pos struct {
	x, y int
}

type Grid struct {
	algo    string
	squares map[Pos]struct{}
}

type Range struct {
	x1, x2, y1, y2 int
}

func get_data() Grid {
	data, _ := ioutil.ReadFile("input.txt")
	data = data[:len(data)-1]
	sections := strings.Split(string(data), "\n\n")

	grid := Grid{algo: sections[0], squares: map[Pos]struct{}{}}
	for y, line := range strings.Split(sections[1], "\n") {
		for x, ru := range line {
			if ru == '#' {
				pos := Pos{x: x, y: y}
				grid.squares[pos] = struct{}{}
			}
		}
	}

	return grid
}

func (grid *Grid) get_range() Range {
	var x_min, x_max, y_min, y_max int

	first := true
	for pos, _ := range grid.squares {
		if first {
			x_min = pos.x
			x_max = pos.x
			y_min = pos.y
			y_max = pos.y
			first = false
			continue
		}
		if pos.x < x_min {
			x_min = pos.x
		}
		if pos.x > x_max {
			x_max = pos.x
		}
		if pos.y < y_min {
			y_min = pos.y
		}
		if pos.y > y_max {
			y_max = pos.y
		}
	}

	return Range{x1: x_min, x2: x_max, y1: y_min, y2: y_max}
}

func (grid *Grid) step(even bool) {
	new_squares := map[Pos]struct{}{}

	r := grid.get_range()

	adjust := 0
	if even {
		adjust = 3
	}

	for y := r.y1 - adjust; y <= r.y2+adjust; y++ {
		for x := r.x1 - adjust; x <= r.x2+adjust; x++ {
			bin := ""
			p := Pos{x: x, y: y}
			_, found := grid.squares[p]

			if !even && (x == r.x1 || x == r.x2 || y == r.y1 || y == r.y2) && found {
				// if even step, switched on and on the outside edge, assume it's surrounded by infinite 1s
				bin = "111111111"
			} else {
				for dy := -1; dy <= 1; dy++ {
					for dx := -1; dx <= 1; dx++ {
						p := Pos{x + dx, y + dy}

						if _, found := grid.squares[p]; found {
							bin += "1"
						} else {
							bin += "0"
						}
					}
				}
			}

			num, _ := strconv.ParseInt(bin, 2, 64)
			if grid.algo[num] == '#' {
				new_squares[p] = struct{}{}
			}
		}
	}

	grid.squares = new_squares
}

func (grid *Grid) print() {
	r := grid.get_range()

	for y := r.y1; y <= r.y2; y++ {
		for x := r.x1; x <= r.x2; x++ {
			p := Pos{x, y}
			if _, found := grid.squares[p]; found {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func main() {
	grid := get_data()

	for i := 0; i < 50; i++ {
		grid.step(i%2 == 0)
		if i == 1 {
			fmt.Println(len(grid.squares))
		}
	}
	fmt.Println(len(grid.squares))
}
