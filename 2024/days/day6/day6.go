package day6

import (
	"fmt"
	"loader"
)

func find_start(grid []string) (int, int) {
	height := len(grid)
	width := len(grid[0])
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if grid[y][x] == '^' {
				return x, y
			}
		}
	}
	return -1, -1
}

func walk(grid []string) int {
	height := len(grid)
	width := len(grid[0])

	x, y := find_start(grid)
	dir := 'U'

	visited := map[[2]int]struct{}{}
	visited[[2]int{x, y}] = struct{}{}

	for {
		next_x := x
		next_y := y

		switch dir {
		case 'U':
			next_y--
		case 'D':
			next_y++
		case 'L':
			next_x--
		case 'R':
			next_x++
		}

		if next_x < 0 || next_x == width || next_y < 0 || next_y == height {
			break
		}

		if grid[next_y][next_x] == '#' {
			switch dir {
			case 'U':
				dir = 'R'
			case 'D':
				dir = 'L'
			case 'L':
				dir = 'U'
			case 'R':
				dir = 'D'
			}
			continue
		}

		x, y = next_x, next_y
		visited[[2]int{x, y}] = struct{}{}
	}

	return len(visited)
}

func Run() {
	loader.Day = 6
	grid := loader.GetStrings()

	part1 := walk(grid)
	part2 := -1

	fmt.Printf("%d %d\n", part1, part2)
}
