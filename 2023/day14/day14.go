package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var height, width int

func load_data(filename string) []string {
	data, _ := ioutil.ReadFile(filename)
	if data[len(data)-1] == '\n' {
		data = data[:len(data)-1]
	}
	return strings.Split(string(data), "\n")
}

func parse_data(data []string) [][]byte {
	height = len(data)
	width = len(data[0])
	grid := [][]byte{}
	for _, line := range data {
		grid = append(grid, []byte(line))
	}
	return grid
}

func tilt_north(grid [][]byte) {
	for y := 1; y < height; y++ {
		for x := 0; x < width; x++ {
			if grid[y][x] == 'O' {
				dx := x
				dy := y
				for dy > 0 && grid[dy-1][dx] == '.' {
					dy--
				}
				if y != dy {
					grid[y][x] = '.'
					grid[dy][x] = 'O'
				}
			}
		}
	}
}

func get_load(grid [][]byte) int {
	load := 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if grid[y][x] == 'O' {
				load += height - y
			}
		}
	}

	return load
}

func pp(grid [][]byte) {
	for _, row := range grid {
		fmt.Println(string(row))
	}
	fmt.Println()
}

func main() {
	data := load_data("input.txt")
	grid := parse_data(data)
	tilt_north(grid)
	fmt.Println(get_load(grid))
}
