package day6

import (
	"fmt"
	"loader"
)

type Pos struct {
	x int
	y int
}

type Guard struct {
	x   int
	y   int
	dir byte
}

var height, width int

func parse_data(data []string) ([][]byte, Guard) {
	height = len(data)
	width = len(data[0])
	grid := [][]byte{}
	for _, line := range data {
		grid = append(grid, []byte(line))
	}
	guard := find_start(grid)
	return grid, guard
}

func find_start(grid [][]byte) Guard {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if grid[y][x] == '^' {
				return Guard{x: x, y: y, dir: 'U'}
			}
		}
	}
	return Guard{}
}

func walk(grid [][]byte, guard Guard) (bool, map[Pos]struct{}) {
	visited := map[Pos]struct{}{}
	visited[Pos{x: guard.x, y: guard.y}] = struct{}{}
	visited_with_dir := map[Guard]struct{}{}
	visited_with_dir[guard] = struct{}{}

	for {
		next_x := guard.x
		next_y := guard.y

		switch guard.dir {
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
			return false, visited
		}

		if grid[next_y][next_x] == '#' {
			switch guard.dir {
			case 'U':
				guard.dir = 'R'
			case 'D':
				guard.dir = 'L'
			case 'L':
				guard.dir = 'U'
			case 'R':
				guard.dir = 'D'
			}
			continue
		}

		guard.x, guard.y = next_x, next_y
		visited[Pos{x: guard.x, y: guard.y}] = struct{}{}
		if _, ok := visited_with_dir[guard]; ok {
			return true, map[Pos]struct{}{}
		}
		visited_with_dir[guard] = struct{}{}
	}

	return false, map[Pos]struct{}{}
}

func obstructions(grid [][]byte, guard Guard, visited map[Pos]struct{}) int {
	count := 0
	for pos := range visited {
		if pos.x == guard.x && pos.y == guard.y {
			continue
		}
		grid[pos.y][pos.x] = '#'
		loop, _ := walk(grid, guard)
		if loop {
			count++
		}
		grid[pos.y][pos.x] = '.'
	}
	return count
}

func Run() {
	loader.Day = 6
	data := loader.GetStrings()
	grid, guard := parse_data(data)

	_, visited := walk(grid, guard)
	part1 := len(visited)
	part2 := obstructions(grid, guard, visited)

	fmt.Printf("%d %d\n", part1, part2)
}
