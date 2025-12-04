package day4

import (
	"fmt"
	"loader"
)

var height, width int

func in_bounds(x int, y int) bool {
	return x >= 0 && x < width && y >= 0 && y < height
}

func parse_data(data []string) [][]byte {
	height = len(data)
	width = len(data[0])

	grid := make([][]byte, height)
	for i := range grid {
		grid[i] = []byte(data[i])
	}

	return grid
}

func neighbours(grid [][]byte, x int, y int) int {
	count := 0
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dx == 0 && dy == 0 {
				continue
			}
			nx, ny := x+dx, y+dy
			if in_bounds(nx, ny) && grid[ny][nx] == '@' {
				count++
			}
		}
	}
	return count
}

func rolls_reachable(grid [][]byte) int {
	reachable := 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if grid[y][x] == '@' && neighbours(grid, x, y) < 4 {
				reachable++
			}
		}
	}
	return reachable
}

func remove_rolls(grid [][]byte) int {
	removed := 0

	for {
		to_remove := [][]int{}
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				if grid[y][x] == '@' && neighbours(grid, x, y) < 4 {
					to_remove = append(to_remove, []int{x, y})
				}
			}
		}

		if len(to_remove) == 0 {
			break
		}

		for _, r := range to_remove {
			grid[r[1]][r[0]] = '.'
		}

		removed += len(to_remove)
	}

	return removed
}

func Run() {
	loader.Day = 4
	data := loader.GetStrings()
	grid := parse_data(data)
	part1 := rolls_reachable(grid)
	part2 := remove_rolls(grid)

	fmt.Printf("%d %d\n", part1, part2)
}
