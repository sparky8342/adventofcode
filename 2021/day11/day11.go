package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Grid struct {
	squares     [][]int
	height      int
	width       int
	flash_count int
}

type Pos struct {
	x int
	y int
}

func get_data() Grid {
	data, _ := ioutil.ReadFile("input.txt")
	data = data[:len(data)-1]
	lines := strings.Split(string(data), "\n")
	grid := Grid{flash_count: 0}
	for _, line := range lines {
		row := []int{}
		for _, num_rune := range line {
			num := int(num_rune - '0')
			row = append(row, num)
		}
		grid.squares = append(grid.squares, row)
	}
	grid.height = len(grid.squares)
	grid.width = len(grid.squares[0])
	return grid
}

func (grid *Grid) step() bool {
	for y := 0; y < grid.height; y++ {
		for x := 0; x < grid.width; x++ {
			grid.squares[y][x]++
		}
	}

	flashed_this_step := 0
	queue := []Pos{}
	for y := 0; y < grid.height; y++ {
		for x := 0; x < grid.width; x++ {
			if grid.squares[y][x] > 9 {
				queue = append(queue, Pos{x: x, y: y})
			}
		}
	}

	flashed := map[Pos]bool{}
	for len(queue) > 0 {
		pos := queue[0]
		queue = queue[1:]

		if _, found := flashed[pos]; found {
			continue
		}
		flashed[pos] = true

		grid.flash_count++
		flashed_this_step++
		for dy := -1; dy <= 1; dy++ {
			for dx := -1; dx <= 1; dx++ {
				if dx == 0 && dy == 0 {
					continue
				}
				new_x := pos.x + dx
				new_y := pos.y + dy
				if new_x < 0 || new_x == grid.width || new_y < 0 || new_y == grid.height {
					continue
				}
				grid.squares[new_y][new_x]++
				if grid.squares[new_y][new_x] > 9 {
					queue = append(queue, Pos{x: new_x, y: new_y})
				}
			}
		}
	}

	for y := 0; y < grid.height; y++ {
		for x := 0; x < grid.width; x++ {
			if grid.squares[y][x] > 9 {
				grid.squares[y][x] = 0
			}
		}
	}

	if flashed_this_step == grid.width*grid.height {
		return true
	} else {
		return false
	}
}

func main() {
	grid := get_data()

	i := 0
	for {
		i++
		all_flashed := grid.step()
		if all_flashed {
			fmt.Println(i)
			break
		}

		if i == 100 {
			fmt.Println(grid.flash_count)
		}
	}
}
