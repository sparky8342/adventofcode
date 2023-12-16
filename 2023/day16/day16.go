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

func load_data(filename string) []string {
	data, _ := ioutil.ReadFile(filename)
	if data[len(data)-1] == '\n' {
		data = data[:len(data)-1]
	}
	return strings.Split(string(data), "\n")
}

func parse_data(data []string) [][]byte {
	grid := [][]byte{}
	for _, line := range data {
		grid = append(grid, []byte(line))
	}
	return grid
}

func energize(grid [][]byte) int {
	height := len(grid)
	width := len(grid[0])

	queue := []Beam{Beam{x: -1, y: 0, direction: EAST}}
	beams_seen := map[Beam]struct{}{}
	energized := map[Pos]struct{}{}

	for len(queue) > 0 {
		beam := queue[0]
		queue = queue[1:]

		for {
			if _, ok := beams_seen[beam]; ok {
				break
			}
			beams_seen[Beam{x: beam.x, y: beam.y, direction: beam.direction}] = struct{}{}
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
			switch grid[beam.y][beam.x] {
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

func pp(grid [][]byte) {
	for _, line := range grid {
		fmt.Println(string(line))
	}
	fmt.Println()
}

func main() {
	data := load_data("input.txt")
	grid := parse_data(data)
	count := energize(grid)
	fmt.Println(count)

}
