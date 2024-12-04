package day4

import (
	"fmt"
	"loader"
)

var dirs [][]int

func init() {
	dirs = [][]int{
		[]int{0, -1},
		[]int{1, -1},
		[]int{1, 0},
		[]int{1, 1},
		[]int{0, 1},
		[]int{-1, 1},
		[]int{-1, 0},
		[]int{-1, -1},
	}
}

func count_xmas(grid []string) int {
	height := len(grid)
	width := len(grid[0])

	count := 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if grid[y][x] == 'X' {
				for _, dir := range dirs {
					word := []byte{}
					next_x := x
					next_y := y
					for i := 0; i < 3; i++ {
						next_x += dir[0]
						next_y += dir[1]
						if next_x < 0 || next_x == width || next_y < 0 || next_y == height {
							break
						}
						word = append(word, grid[next_y][next_x])
					}
					if string(word) == "MAS" {
						count++
					}
				}
			}
		}
	}

	return count
}

func Run() {
	loader.Day = 4
	grid := loader.GetStrings()

	part1 := count_xmas(grid)
	part2 := -1

	fmt.Printf("%d %d\n", part1, part2)
}
