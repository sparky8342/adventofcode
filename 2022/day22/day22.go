package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Command struct {
	steps int
	turn  rune
}

type Pos struct {
	x int
	y int
}

type Grid struct {
	squares  map[Pos]rune
	width    int
	height   int
	commands []Command
}

func load_data(filename string) []string {
	data, _ := ioutil.ReadFile(filename)
	if data[len(data)-1] == '\n' {
		data = data[:len(data)-1]
	}
	return strings.Split(string(data), "\n")
}

func parse_data(data []string) *Grid {
	command_str := data[len(data)-1]
	data = data[:len(data)-1]

	grid := Grid{squares: map[Pos]rune{}, height: len(data)}

	for y, row := range data {
		for x, ru := range row {
			grid.squares[Pos{x: x, y: y}] = ru
			if x > grid.width {
				grid.width = x
			}
		}
	}
	grid.width++

	num_str := ""
	for _, ru := range command_str {
		if ru == 'L' || ru == 'R' {
			num, _ := strconv.Atoi(num_str)
			grid.commands = append(grid.commands, Command{steps: num, turn: ru})
			num_str = ""
		} else {
			num_str += string(ru)
		}
	}
	num, _ := strconv.Atoi(num_str)
	grid.commands = append(grid.commands, Command{steps: num})

	return &grid
}

func (grid *Grid) walk() int {
	var x, y int
	for i := 0; i < grid.width; i++ {
		val, exists := grid.squares[Pos{x: i, y: 0}]
		if exists && val == '.' {
			x = i
			y = 0
			break
		}
	}

	dx := 1
	dy := 0

	for _, command := range grid.commands {
		for i := 0; i < command.steps; i++ {
			next_x := x
			next_y := y
			for {
				next_x += dx
				next_y += dy
				if next_x == -1 {
					next_x = grid.width - 1
				} else if next_x == grid.width {
					next_x = 0
				}
				if next_y == -1 {
					next_y = grid.height - 1
				} else if next_y == grid.height {
					next_y = 0
				}

				val, exists := grid.squares[Pos{x: next_x, y: next_y}]
				if !exists {
					continue
				}

				if val == '.' {
					x = next_x
					y = next_y
					break
				} else if val == '#' {
					break
				}
			}

		}

		if command.turn == 'L' {
			if dx == 0 && dy == -1 {
				dx = -1
				dy = 0
			} else if dx == -1 && dy == 0 {
				dx = 0
				dy = 1
			} else if dx == 0 && dy == 1 {
				dx = 1
				dy = 0
			} else if dx == 1 && dy == 0 {
				dx = 0
				dy = -1
			}
		} else if command.turn == 'R' {
			if dx == 0 && dy == -1 {
				dx = 1
				dy = 0
			} else if dx == 1 && dy == 0 {
				dx = 0
				dy = 1
			} else if dx == 0 && dy == 1 {
				dx = -1
				dy = 0
			} else if dx == -1 && dy == 0 {
				dx = 0
				dy = -1
			}
		}

	}

	var facing int
	if dx == 0 && dy == -1 {
		facing = 3
	} else if dx == 0 && dy == 1 {
		facing = 1
	} else if dx == -1 && dy == 0 {
		facing = 2
	} else if dx == 1 && dy == 0 {
		facing = 0
	}

	return (y+1)*1000 + (x+1)*4 + facing
}

func (grid *Grid) print_grid(you_x int, you_y int) {
	for y := 0; y < grid.height; y++ {
		for x := 0; x < grid.width; x++ {
			if x == you_x && y == you_y {
				fmt.Print("o")
				continue
			}
			pos := Pos{x: x, y: y}
			val, exists := grid.squares[pos]
			if exists {
				fmt.Print(string(val))
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func main() {
	data := load_data("input.txt")
	grid := parse_data(data)
	fmt.Println(grid.walk())
}
