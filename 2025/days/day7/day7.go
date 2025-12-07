package day7

import (
	"fmt"
	"loader"
)

var width, height int

type Pos struct {
	x int
	y int
}

func beam(grid []string, pos Pos, splits map[Pos]int) int {
	for grid[pos.y][pos.x] == '.' {
		pos.y++
		if pos.y == height-1 {
			return 1
		}
	}
	if val, ok := splits[pos]; ok {
		return val
	}
	n := beam(grid, Pos{x: pos.x - 1, y: pos.y}, splits)
	n += beam(grid, Pos{x: pos.x + 1, y: pos.y}, splits)
	splits[pos] = n
	return n
}

func follow_beam(grid []string) (int, int) {
	height = len(grid)
	width = len(grid[0])

	start := Pos{x: width / 2, y: 1}

	splits := map[Pos]int{}
	timelines := beam(grid, start, splits)
	return len(splits), timelines
}

func Run() {
	loader.Day = 7
	grid := loader.GetStrings()
	part1, part2 := follow_beam(grid)

	fmt.Printf("%d %d\n", part1, part2)
}
