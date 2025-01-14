package day4

import (
	"fmt"
	"loader"
)

func count_xmas(grid []string) int {
	height := len(grid)
	width := len(grid[0])

	dirs := [][]int{
		[]int{0, -1},
		[]int{1, -1},
		[]int{1, 0},
		[]int{1, 1},
		[]int{0, 1},
		[]int{-1, 1},
		[]int{-1, 0},
		[]int{-1, -1},
	}

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

func count_x_mas(grid []string) int {
	height := len(grid)
	width := len(grid[0])

	count := 0

	for y := 1; y < height-1; y++ {
		for x := 1; x < width-1; x++ {
			if grid[y][x] == 'A' {
				letters := string([]byte{grid[y-1][x-1], grid[y-1][x+1], grid[y+1][x+1], grid[y+1][x-1]})
				if letters == "MMSS" || letters == "SMMS" || letters == "SSMM" || letters == "MSSM" {
					count++
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
	part2 := count_x_mas(grid)

	fmt.Printf("%d %d\n", part1, part2)
}
