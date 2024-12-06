package day6

import (
	"fmt"
	"loader"
)

var height, width int

func parse_data(data []string) ([][]byte, int, int) {
	height = len(data)
	width = len(data[0])
	grid := [][]byte{}
	for _, line := range data {
		grid = append(grid, []byte(line))
	}
	x, y := find_start(grid)
	return grid, x, y
}

func find_start(grid [][]byte) (int, int) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if grid[y][x] == '^' {
				return x, y
			}
		}
	}
	return -1, -1
}

func walk(grid [][]byte, x int, y int) (bool, int) {
	dir := 'U'

	visited := map[[2]int]int{}
	visited[[2]int{x, y}] = 1

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
			return false, len(visited)
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
		visited[[2]int{x, y}]++
		if visited[[2]int{x, y}] == 5 {
			return true, 0
		}
	}

	return false, 0
}

func obstructions(grid [][]byte, start_x int, start_y int) int {
	count := 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if grid[y][x] == '.' {
				grid[y][x] = '#'
				loop, _ := walk(grid, start_x, start_y)
				if loop {
					count++
				}
				grid[y][x] = '.'
			}
		}
	}
	return count
}

func Run() {
	loader.Day = 6
	data := loader.GetStrings()
	grid, start_x, start_y := parse_data(data)

	_, part1 := walk(grid, start_x, start_y)
	part2 := obstructions(grid, start_x, start_y)

	fmt.Printf("%d %d\n", part1, part2)
}
