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

func beam(grid []string, pos Pos, splits map[Pos]struct{}) {
	for grid[pos.y][pos.x] == '.' {
		pos.y++
		if pos.y == height-1 {
			return
		}
	}
	if _, ok := splits[pos]; ok {
		return
	}
	splits[pos] = struct{}{}
	beam(grid, Pos{x: pos.x - 1, y: pos.y}, splits)
	beam(grid, Pos{x: pos.x + 1, y: pos.y}, splits)
}

func follow_beam(grid []string) int {
	height = len(grid)
	width = len(grid[0])

	start := Pos{x: width / 2, y: 1}

	splits := map[Pos]struct{}{}
	beam(grid, start, splits)
	return len(splits)
}

func Run() {
	loader.Day = 7
	grid := loader.GetStrings()
	part1 := follow_beam(grid)

	fmt.Printf("%d %d\n", part1, 0)
}
