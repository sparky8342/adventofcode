package day15

import (
	"fmt"
	"loader"
	"strings"
)

type Pos struct {
	x int
	y int
}

var size int

func find_robot(grid [][]byte) Pos {
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			if grid[y][x] == '@' {
				return Pos{x: x, y: y}
			}
		}
	}
	return Pos{}
}

func next_space(x int, y int, command rune) (int, int) {
	switch command {
	case '^':
		y--
	case '>':
		x++
	case 'v':
		y++
	case '<':
		x--
	}
	return x, y
}

func move_robot(data [][]string) int {
	size = len(data[0])
	grid := make([][]byte, size)
	for i, line := range data[0] {
		grid[i] = []byte(line)
	}
	commands := strings.Join(data[1], "")

	robot := find_robot(grid)
	grid[robot.y][robot.x] = '.'
	for _, command := range commands {
		next_x, next_y := next_space(robot.x, robot.y, command)
		if grid[next_y][next_x] == '.' {
			robot.x = next_x
			robot.y = next_y
		} else if grid[next_y][next_x] == '#' {
			continue
		} else if grid[next_y][next_x] == 'O' {
			after_x, after_y := next_x, next_y
			for {
				after_x, after_y = next_space(after_x, after_y, command)
				if grid[after_y][after_x] == '#' {
					break
				} else if grid[after_y][after_x] == '.' {
					grid[after_y][after_x] = 'O'
					grid[next_y][next_x] = '.'
					robot.x = next_x
					robot.y = next_y
					break
				}
			}
		}
	}

	gps := 0
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			if grid[y][x] == 'O' {
				gps += y*100 + x
			}
		}
	}
	return gps
}

func Run() {
	loader.Day = 15
	data := loader.GetStringGroups()

	part1 := move_robot(data)
	part2 := -1

	fmt.Printf("%d %d\n", part1, part2)
}
