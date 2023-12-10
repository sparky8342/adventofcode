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

func get_data(filename string) []string {
	data, _ := ioutil.ReadFile(filename)
	if data[len(data)-1] == '\n' {
		data = data[:len(data)-1]
	}
	return strings.Split(string(data), "\n")
}

func find_start(data []string) Pos {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if data[y][x] == 'S' {
				return Pos{x: x, y: y}
			}
		}
	}
	return Pos{x: -1, y: -1}
}

func connects(data []string, pos Pos) Pos {
	north_pos := Pos{x: pos.x, y: pos.y - 1, dx: 0, dy: -1}
	north := data[pos.y-1][pos.x]
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
	east := data[pos.y][pos.x+1]
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
	south := data[pos.y+1][pos.x]
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
	west := data[pos.y][pos.x-1]
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

func move(data []string, pos Pos) Pos {
	pos.x += pos.dx
	pos.y += pos.dy

	b := data[pos.y][pos.x]
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

func is_inside(data []string, path map[Square]Empty, x int, y int) bool {
	if _, exists := path[Square{x: x, y: y}]; exists {
		return false
	}

	lx := x - 1
	crosses := 0
	for lx >= 0 {
		if _, exists := path[Square{x: lx, y: y}]; exists {
			if data[y][lx] == '|' || data[y][lx] == 'J' || data[y][lx] == 'L' {
				crosses++
			}
		}
		lx--
	}
	return crosses%2 == 1
}

func count_inside(data []string, path map[Square]Empty) int {
	inside := 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if is_inside(data, path, x, y) {
				inside++
			}
		}
	}
	return inside
}

func path_info(data []string) (int, int) {
	height = len(data)
	width = len(data[0])

	path := map[Square]Empty{}

	start := find_start(data)
	path[Square{x: start.x, y: start.y}] = Empty{}

	pos := connects(data, start)
	path[Square{x: pos.x, y: pos.y}] = Empty{}

	steps := 1
	for !(pos.x == start.x && pos.y == start.y) {
		pos = move(data, pos)
		path[Square{x: pos.x, y: pos.y}] = Empty{}
		steps++
	}

	inside := count_inside(data, path)
	return steps / 2, inside
}

func pp(data [][]byte) {
	for _, row := range data {
		fmt.Println(string(row))
	}
}

func main() {
	data := get_data("input.txt")
	distance, inside := path_info(data)
	fmt.Println(distance)
	fmt.Println(inside)
}
