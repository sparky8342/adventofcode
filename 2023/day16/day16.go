package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Pos struct {
	x int
	y int
}

type Beam struct {
	x         int
	y         int
	direction int
}

const (
	NORTH = 1
	EAST  = 2
	SOUTH = 3
	WEST  = 4
)

var height, width int

func load_data(filename string) []string {
	data, _ := ioutil.ReadFile(filename)
	if data[len(data)-1] == '\n' {
		data = data[:len(data)-1]
	}
	return strings.Split(string(data), "\n")
}

func parse_data(data []string) {
	height = len(data)
	width = len(data[0])
}

func energize_top_left(data []string) int {
	return energize(data, Beam{x: -1, y: 0, direction: EAST})
}

func max_energy(data []string) int {
	max := 0

	for y := 0; y < height; y++ {
		energy := energize(data, Beam{x: -1, y: y, direction: EAST})
		if energy > max {
			max = energy
		}

		energy = energize(data, Beam{x: width, y: y, direction: WEST})
		if energy > max {
			max = energy
		}
	}

	for x := 0; x < width; x++ {
		energy := energize(data, Beam{x: x, y: -1, direction: SOUTH})
		if energy > max {
			max = energy
		}

		energy = energize(data, Beam{x: x, y: height, direction: NORTH})
		if energy > max {
			max = energy
		}
	}

	return max
}

func energize(data []string, start Beam) int {
	queue := []Beam{start}
	beams_seen := map[Beam]struct{}{}
	energized := map[Pos]struct{}{}

	for len(queue) > 0 {
		beam := queue[0]
		queue = queue[1:]

		for {
			if _, ok := beams_seen[beam]; ok {
				break
			}
			beams_seen[beam] = struct{}{}

			// move
			switch beam.direction {
			case NORTH:
				beam.y--
			case EAST:
				beam.x++
			case SOUTH:
				beam.y++
			case WEST:
				beam.x--
			}

			if beam.x < 0 || beam.x == width || beam.y < 0 || beam.y == height {
				break
			}

			energized[Pos{x: beam.x, y: beam.y}] = struct{}{}

			// turn/split
			switch data[beam.y][beam.x] {
			case '/':
				switch beam.direction {
				case NORTH:
					beam.direction = EAST
				case EAST:
					beam.direction = NORTH
				case SOUTH:
					beam.direction = WEST
				case WEST:
					beam.direction = SOUTH
				}
			case '\\':
				switch beam.direction {
				case NORTH:
					beam.direction = WEST
				case EAST:
					beam.direction = SOUTH
				case SOUTH:
					beam.direction = EAST
				case WEST:
					beam.direction = NORTH
				}
			case '-':
				if beam.direction == NORTH || beam.direction == SOUTH {
					split := Beam{x: beam.x, y: beam.y, direction: EAST}
					queue = append(queue, split)
					beam.direction = WEST
				}
			case '|':
				if beam.direction == EAST || beam.direction == WEST {
					split := Beam{x: beam.x, y: beam.y, direction: SOUTH}
					queue = append(queue, split)
					beam.direction = NORTH
				}
			}
		}
	}

	return len(energized)
}

func pp(data [][]byte) {
	for _, line := range data {
		fmt.Println(string(line))
	}
	fmt.Println()
}

func main() {
	data := load_data("input.txt")
	parse_data(data)
	count := energize_top_left(data)
	fmt.Println(count)
	fmt.Println(max_energy(data))
}
