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

var height, width int

func load_data(filename string) []string {
	data, _ := ioutil.ReadFile(filename)
	if data[len(data)-1] == '\n' {
		data = data[:len(data)-1]
	}
	return strings.Split(string(data), "\n")
}

func get_at(data []string, pos Pos) byte {
	if pos.x < 0 || pos.x == width || pos.y < 0 || pos.y > height {
		return '.'
	} else {
		return data[pos.y][pos.x]
	}
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

func connects(data []string, pos Pos) []Pos {
	positions := []Pos{}

	north_pos := Pos{x: pos.x, y: pos.y - 1, dx: 0, dy: -1}
	north := get_at(data, north_pos)
	if north == '|' || north == '7' || north == 'F' {
		if north == '7' {
			north_pos.dx = -1
			north_pos.dy = 0
		} else if north == 'F' {
			north_pos.dx = 1
			north_pos.dy = 0
		}
		positions = append(positions, north_pos)
	}

	east_pos := Pos{x: pos.x + 1, y: pos.y, dx: 1, dy: 0}
	east := get_at(data, east_pos)
	if east == '-' || east == 'J' || east == '7' {
		if east == 'J' {
			east_pos.dx = 0
			east_pos.dy = -1
		} else if east == '7' {
			east_pos.dx = 0
			east_pos.dy = 1
		}
		positions = append(positions, east_pos)
	}

	south_pos := Pos{x: pos.x, y: pos.y + 1, dx: 0, dy: 1}
	south := get_at(data, south_pos)
	if south == '|' || south == 'L' || south == 'J' {
		if south == 'L' {
			south_pos.dx = 1
			south_pos.dy = 0
		} else if south == 'J' {
			south_pos.dx = -1
			south_pos.dy = 0
		}
		positions = append(positions, south_pos)
	}

	west_pos := Pos{x: pos.x - 1, y: pos.y, dx: -1, dy: 0}
	west := get_at(data, west_pos)
	if west == '-' || west == 'L' || west == 'F' {
		if west == 'L' {
			west_pos.dx = 0
			west_pos.dy = -1
		} else if west == 'F' {
			west_pos.dx = 0
			west_pos.dy = 1
		}
		positions = append(positions, west_pos)
	}

	return positions
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

func distance(data []string) int {
	height = len(data)
	width = len(data[0])

	start := find_start(data)

	positions := connects(data, start)
	pos1 := positions[1]

	steps := 1
	for !(pos1.x == start.x && pos1.y == start.y) {
		pos1 = move(data, pos1)
		steps++
	}

	return steps / 2
}

func main() {
	data := load_data("input.txt")
	fmt.Println(distance(data))
}
