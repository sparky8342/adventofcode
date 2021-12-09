package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Grid struct {
	squares [][]int
	width   int
	height  int
}

func get_data() Grid {
	data, _ := ioutil.ReadFile("input.txt")
	data = data[:len(data)-1]
	lines := strings.Split(string(data), "\n")

	grid := Grid{}
	for _, line := range lines {
		row := []int{}
		for _, ch := range line {
			row = append(row, int(ch-'0'))
		}
		grid.squares = append(grid.squares, row)
	}
	grid.height = len(grid.squares)
	grid.width = len(grid.squares[0])
	return grid
}

func lower_than_surrounding(grid Grid, x int, y int) bool {
	val := grid.squares[y][x]
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			new_y := y + dy
			new_x := x + dx
			if new_y < 0 || new_y == grid.height || new_x < 0 || new_x == grid.width {
				continue
			}
			if grid.squares[new_y][new_x] < val {
				return false
			}
		}
	}
	return true
}

func main() {
	grid := get_data()

	risk_sum := 0
	for y := 0; y < grid.height; y++ {
		for x := 0; x < grid.width; x++ {
			if lower_than_surrounding(grid, x, y) {
				risk_sum += grid.squares[y][x] + 1
			}
		}
	}
	fmt.Println(risk_sum)
}
