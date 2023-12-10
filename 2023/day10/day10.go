package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Pos struct {
	x  int
	y  int
	dx int
	dy int
}

type Square struct {
	x int
	y int
}

type Empty struct {
}

var height, width int

func get_data(filename string) [][]byte {
	data, _ := ioutil.ReadFile(filename)
	if data[len(data)-1] == '\n' {
		data = data[:len(data)-1]
	}
	lines := strings.Split(string(data), "\n")
	grid := [][]byte{}
	for _, line := range lines {
		grid = append(grid, []byte(line))
	}
	return grid
}

func find_start(grid [][]byte) Pos {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if grid[y][x] == 'S' {
				return Pos{x: x, y: y}
			}
		}
	}
	return Pos{x: -1, y: -1}
}

func connects(grid [][]byte, pos Pos) Pos {
	north_pos := Pos{x: pos.x, y: pos.y - 1, dx: 0, dy: -1}
	north := grid[pos.y-1][pos.x]
	if north == '|' || north == '7' || north == 'F' {
		if north == '7' {
			north_pos.dx = -1
			north_pos.dy = 0
		} else if north == 'F' {
			north_pos.dx = 1
			north_pos.dy = 0
		}
		return north_pos
	}

	east_pos := Pos{x: pos.x + 1, y: pos.y, dx: 1, dy: 0}
	east := grid[pos.y][pos.x+1]
	if east == '-' || east == 'J' || east == '7' {
		if east == 'J' {
			east_pos.dx = 0
			east_pos.dy = -1
		} else if east == '7' {
			east_pos.dx = 0
			east_pos.dy = 1
		}
		return east_pos
	}

	south_pos := Pos{x: pos.x, y: pos.y + 1, dx: 0, dy: 1}
	south := grid[pos.y+1][pos.x]
	if south == '|' || south == 'L' || south == 'J' {
		if south == 'L' {
			south_pos.dx = 1
			south_pos.dy = 0
		} else if south == 'J' {
			south_pos.dx = -1
			south_pos.dy = 0
		}
		return south_pos
	}

	west_pos := Pos{x: pos.x - 1, y: pos.y, dx: -1, dy: 0}
	west := grid[pos.y][pos.x-1]
	if west == '-' || west == 'L' || west == 'F' {
		if west == 'L' {
			west_pos.dx = 0
			west_pos.dy = -1
		} else if west == 'F' {
			west_pos.dx = 0
			west_pos.dy = 1
		}
		return west_pos
	}

	return Pos{}
}

func move(grid [][]byte, pos Pos) Pos {
	pos.x += pos.dx
	pos.y += pos.dy

	b := grid[pos.y][pos.x]
	if pos.dx == 1 {
		// going east
		if b == 'J' {
			pos.dx = 0
			pos.dy = -1
		} else if b == '7' {
			pos.dx = 0
			pos.dy = 1
		}
	} else if pos.dx == -1 {
		// going west
		if b == 'L' {
			pos.dx = 0
			pos.dy = -1
		} else if b == 'F' {
			pos.dx = 0
			pos.dy = 1
		}
	} else if pos.dy == 1 {
		// going south
		if b == 'L' {
			pos.dx = 1
			pos.dy = 0
		} else if b == 'J' {
			pos.dx = -1
			pos.dy = 0
		}
	} else if pos.dy == -1 {
		// going north
		if b == 'F' {
			pos.dx = 1
			pos.dy = 0
		} else if b == '7' {
			pos.dx = -1
			pos.dy = 0
		}
	}
	return pos
}

func is_inside(grid [][]byte, path map[Square]Empty, x int, y int) bool {
	if _, exists := path[Square{x: x, y: y}]; exists {
		return false
	}

	lx := x - 1
	crosses := 0
	for lx >= 0 {
		if _, exists := path[Square{x: lx, y: y}]; exists {
			if grid[y][lx] == '|' || grid[y][lx] == 'J' || grid[y][lx] == 'L' {
				crosses++
			}
		}
		lx--
	}
	return crosses%2 == 1
}

func count_inside(grid [][]byte, path map[Square]Empty) int {
	inside := 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if is_inside(grid, path, x, y) {
				inside++
			}
		}
	}
	return inside
}

func path_info(grid [][]byte) (int, int) {
	height = len(grid)
	width = len(grid[0])

	path := map[Square]Empty{}

	start := find_start(grid)
	corners := []Pos{}

	path[Square{x: start.x, y: start.y}] = Empty{}

	first_pos := connects(grid, start)
	if grid[first_pos.y][first_pos.x] != '|' && grid[first_pos.y][first_pos.x] != '-' {
		corners = append(corners, first_pos)
	}
	path[Square{x: first_pos.x, y: first_pos.y}] = Empty{}

	pos := first_pos
	steps := 1
	for !(pos.x == start.x && pos.y == start.y) {
		pos = move(grid, pos)
		if grid[pos.y][pos.x] != '|' && grid[pos.y][pos.x] != '-' {
			corners = append(corners, pos)
		}

		path[Square{x: pos.x, y: pos.y}] = Empty{}
		steps++
	}

	inside := count_inside(grid, path)
	return steps / 2, inside
}

func pp(grid [][]byte) {
	for _, row := range grid {
		fmt.Println(string(row))
	}
}

func main() {
	grid := get_data("input.txt")
	distance, inside := path_info(grid)
	fmt.Println(distance)
	fmt.Println(inside)
}
