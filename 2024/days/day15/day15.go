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

func expand_grid(data []string) (map[Pos]byte, Pos) {
	blocks := map[Pos]byte{}
	robot := Pos{}
	for i, line := range data {
		for j, ru := range line {
			switch ru {
			case '#':
				blocks[Pos{x: j * 2, y: i}] = 'W'
				blocks[Pos{x: j*2 + 1, y: i}] = 'W'
			case 'O':
				blocks[Pos{x: j * 2, y: i}] = 'L'
				blocks[Pos{x: j*2 + 1, y: i}] = 'R'
			case '@':
				robot.x = j * 2
				robot.y = i
			}
		}
	}
	return blocks, robot
}

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

func can_move(pos Pos, blocks map[Pos]byte, dir rune) bool {
	if dir == '<' || dir == '>' {
		var inc int
		if dir == '<' {
			inc = -1
		} else if dir == '>' {
			inc = 1
		}

		next := Pos{x: pos.x + inc, y: pos.y}
		val, ok := blocks[next]
		if !ok {
			return true
		} else if val == 'W' {
			return false
		} else if val == 'L' || val == 'R' {
			return can_move(next, blocks, dir)
		}
	} else if dir == '^' || dir == 'v' {
		var inc int
		if dir == '^' {
			inc = -1
		} else if dir == 'v' {
			inc = 1
		}

		next := Pos{x: pos.x, y: pos.y + inc}
		val, ok := blocks[next]
		if !ok {
			return true
		} else if val == 'W' {
			return false
		} else if val == 'L' || val == 'R' {
			other := Pos{x: next.x, y: next.y}
			if val == 'L' {
				other.x++
			} else if val == 'R' {
				other.x--
			}
			return can_move(next, blocks, dir) && can_move(other, blocks, dir)
		}
	}
	return false
}

func move(pos Pos, blocks map[Pos]byte, dir rune) {
	if dir == '<' || dir == '>' {
		var inc int
		if dir == '<' {
			inc = -1
		} else if dir == '>' {
			inc = 1
		}

		next := Pos{x: pos.x + inc, y: pos.y}
		_, ok := blocks[next]
		if ok {
			move(next, blocks, dir)
		}
		val := blocks[pos]
		delete(blocks, pos)
		blocks[next] = val
	} else if dir == '^' || dir == 'v' {
		var inc int
		if dir == '^' {
			inc = -1
		} else if dir == 'v' {
			inc = 1
		}

		next := Pos{x: pos.x, y: pos.y + inc}
		val, ok := blocks[next]
		if ok {
			other := Pos{x: next.x, y: next.y}
			if val == 'L' {
				other.x++
			} else if val == 'R' {
				other.x--
			}
			move(next, blocks, dir)
			move(other, blocks, dir)
		}
		val = blocks[pos]
		delete(blocks, pos)
		blocks[next] = val
	}
}

func print_blocks(blocks map[Pos]byte, robot Pos) {
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			if x == robot.x && y == robot.y {
				fmt.Print("@")
				continue
			}
			pos := Pos{x: x, y: y}
			if val, ok := blocks[pos]; ok {
				if val == 'W' {
					fmt.Print("#")
				} else if val == 'L' {
					fmt.Print("[")
				} else if val == 'R' {
					fmt.Print("]")
				}
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func move_robot_double_blocks(data [][]string) int {
	size = len(data[0]) * 2
	blocks, robot := expand_grid(data[0])
	commands := strings.Join(data[1], "")

	for _, command := range commands {
		if can_move(robot, blocks, command) {
			move(robot, blocks, command)
			robot.x, robot.y = next_space(robot.x, robot.y, command)
		}
	}

	gps := 0
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			pos := Pos{x: x, y: y}
			if val, ok := blocks[pos]; ok && val == 'L' {
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
	part2 := move_robot_double_blocks(data)

	fmt.Printf("%d %d\n", part1, part2)
}
