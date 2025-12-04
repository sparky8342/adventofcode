package day4

import (
	"fmt"
	"loader"
)

var height, width int

func in_bounds(x int, y int) bool {
	return x >= 0 && x < width && y >= 0 && y < height
}

func neighbours(grid []string, x int, y int) int {
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

func rolls_reachable(grid []string) int {
	height = len(grid)
	width = len(grid[0])

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

func Run() {
	loader.Day = 4
	grid := loader.GetStrings()
	part1 := rolls_reachable(grid)

	fmt.Printf("%d %d\n", part1, 0)
}
