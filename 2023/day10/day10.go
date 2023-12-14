package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Pos struct {
	x         int
	y         int
	direction int
}

type Square struct {
	x int
	y int
}

type Empty struct {
}

const (
	NORTH = 1
	EAST  = 2
	SOUTH = 3
	WEST  = 4
)

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
	north := data[pos.y-1][pos.x]
	if north == '|' || north == '7' || north == 'F' {
		north_pos := Pos{x: pos.x, y: pos.y - 1}
		if north == '7' {
			north_pos.direction = WEST
		} else if north == 'F' {
			north_pos.direction = EAST
		} else {
			north_pos.direction = NORTH
		}
		return north_pos
	}

	east := data[pos.y][pos.x+1]
	if east == '-' || east == 'J' || east == '7' {
		east_pos := Pos{x: pos.x + 1, y: pos.y}
		if east == 'J' {
			east_pos.direction = NORTH
		} else if east == '7' {
			east_pos.direction = SOUTH
		} else {
			east_pos.direction = EAST
		}
		return east_pos
	}

	south := data[pos.y+1][pos.x]
	if south == '|' || south == 'L' || south == 'J' {
		south_pos := Pos{x: pos.x, y: pos.y + 1}
		if south == 'L' {
			south_pos.direction = EAST
		} else if south == 'J' {
			south_pos.direction = WEST
		} else {
			south_pos.direction = SOUTH
		}
		return south_pos
	}

	west := data[pos.y][pos.x-1]
	if west == '-' || west == 'L' || west == 'F' {
		west_pos := Pos{x: pos.x - 1, y: pos.y}
		if west == 'L' {
			west_pos.direction = NORTH
		} else if west == 'F' {
			west_pos.direction = SOUTH
		} else {
			west_pos.direction = WEST
		}
		return west_pos
	}

	return Pos{}
}

func move(data []string, pos Pos) Pos {
	switch pos.direction {
	case NORTH:
		pos.y--
		b := data[pos.y][pos.x]
		if b == 'F' {
			pos.direction = EAST
		} else if b == '7' {
			pos.direction = WEST
		}
	case EAST:
		pos.x++
		b := data[pos.y][pos.x]
		if b == 'J' {
			pos.direction = NORTH
		} else if b == '7' {
			pos.direction = SOUTH
		}
	case SOUTH:
		pos.y++
		b := data[pos.y][pos.x]
		if b == 'L' {
			pos.direction = EAST
		} else if b == 'J' {
			pos.direction = WEST
		}
	case WEST:
		pos.x--
		b := data[pos.y][pos.x]
		if b == 'L' {
			pos.direction = NORTH
		} else if b == 'F' {
			pos.direction = SOUTH
		}
	}

	return pos
}

func count_inside(data []string, path map[Square]Empty) int {
	squares_inside := 0
	inside := false
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if _, exists := path[Square{x: x, y: y}]; exists {
				if data[y][x] == '|' || data[y][x] == 'J' || data[y][x] == 'L' {
					inside = !inside
				}
			} else if inside {
				squares_inside++
			}
		}
	}
	return squares_inside
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
