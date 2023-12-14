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
				dy := y
				for dy > 0 && grid[dy-1][x] == '.' {
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

func tilt_west(grid [][]byte) {
	for x := 1; x < width; x++ {
		for y := 0; y < height; y++ {
			if grid[y][x] == 'O' {
				dx := x
				for dx > 0 && grid[y][dx-1] == '.' {
					dx--
				}
				if x != dx {
					grid[y][x] = '.'
					grid[y][dx] = 'O'
				}
			}
		}
	}
}

func tilt_south(grid [][]byte) {
	for y := height - 2; y >= 0; y-- {
		for x := 0; x < width; x++ {
			if grid[y][x] == 'O' {
				dy := y
				for dy < height-1 && grid[dy+1][x] == '.' {
					dy++
				}
				if y != dy {
					grid[y][x] = '.'
					grid[dy][x] = 'O'
				}
			}
		}
	}
}

func tilt_east(grid [][]byte) {
	for x := width - 2; x >= 0; x-- {
		for y := 0; y < height; y++ {
			if grid[y][x] == 'O' {
				dx := x
				for dx < width-1 && grid[y][dx+1] == '.' {
					dx++
				}
				if x != dx {
					grid[y][x] = '.'
					grid[y][dx] = 'O'
				}
			}
		}
	}
}

func spin(grid [][]byte) {
	tilt_north(grid)
	tilt_west(grid)
	tilt_south(grid)
	tilt_east(grid)
}

func flatten(grid [][]byte) string {
	str := ""
	for _, line := range grid {
		str += string(line)
	}
	return str
}

func spins(grid [][]byte, amount int) {
	done := 0

	seen := map[string]int{}
	seen[flatten(grid)] = 0

	var first int
	for {
		spin(grid)
		done++
		flat := flatten(grid)
		if value, exists := seen[flat]; exists {
			first = value
			break
		} else {
			seen[flat] = done
		}
	}

	length := done - first

	for i := 0; i < (amount-done)%length; i++ {
		spin(grid)
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

	spins(grid, 1000000000)
	fmt.Println(get_load(grid))
}
