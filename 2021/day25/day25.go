package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Grid struct {
	squares [][]rune
	height  int
	width   int
}

func get_data() Grid {
	data, _ := ioutil.ReadFile("input.txt")
	data = data[:len(data)-1]
	lines := strings.Split(string(data), "\n")
	grid := Grid{}
	for _, line := range lines {
		row := []rune{}
		for _, ru := range line {
			row = append(row, ru)
		}
		grid.squares = append(grid.squares, row)
	}
	grid.height = len(grid.squares)
	grid.width = len(grid.squares[0])
	return grid
}

func (grid *Grid) step() bool {
	changed := false
	for y := 0; y < grid.height; y++ {
		move := []int{}
		for x := 0; x < grid.width; x++ {
			if grid.squares[y][x] == '>' && grid.squares[y][(x+1)%grid.width] == '.' {
				move = append(move, x)
			}
		}
		if len(move) > 0 {
			changed = true
			for _, x := range move {
				grid.squares[y][x] = '.'
				grid.squares[y][(x+1)%grid.width] = '>'
			}
		}
	}

	for x := 0; x < grid.width; x++ {
		move := []int{}
		for y := 0; y < grid.height; y++ {
			if grid.squares[y][x] == 'v' && grid.squares[(y+1)%grid.height][x] == '.' {
				move = append(move, y)
			}
		}
		if len(move) > 0 {
			changed = true
			for _, y := range move {
				grid.squares[y][x] = '.'
				grid.squares[(y+1)%grid.height][x] = 'v'
			}
		}
	}

	return changed
}

func (grid *Grid) print() {
	for y := 0; y < grid.height; y++ {
		for x := 0; x < grid.width; x++ {
			fmt.Print(string(grid.squares[y][x]))
		}
		fmt.Println()
	}
	fmt.Println()
}

func main() {
	grid := get_data()
	steps := 1
	for grid.step() == true {
		steps++
	}
	fmt.Println(steps)
}
