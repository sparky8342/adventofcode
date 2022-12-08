package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Grid struct {
	width   int
	height  int
	squares [][]uint8
}

type Pos struct {
	x int
	y int
}

type Empty struct {
}

func load_data(filename string) []string {
	data, _ := ioutil.ReadFile(filename)
	if data[len(data)-1] == '\n' {
		data = data[:len(data)-1]
	}
	return strings.Split(string(data), "\n")
}

func make_grid(data []string) Grid {
	height := len(data)
	width := len(data[0])

	squares := make([][]uint8, height)
	for i := range squares {
		squares[i] = make([]uint8, width)
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			squares[y][x] = data[y][x] - '0'
		}
	}

	return Grid{width: width, height: height, squares: squares}
}

func visible(grid Grid) int {
	trees := map[Pos]Empty{}

	// from top
	for x := 0; x < grid.width; x++ {
		var max uint8 = 0
		for y := 0; y < grid.height; y++ {
			if y == 0 || grid.squares[y][x] > max {
				trees[Pos{x: x, y: y}] = Empty{}
				max = grid.squares[y][x]
			}
		}
	}

	// from bottom
	for x := 0; x < grid.width; x++ {
		var max uint8 = 0
		for y := grid.height - 1; y >= 0; y-- {
			if y == grid.height-1 || grid.squares[y][x] > max {
				trees[Pos{x: x, y: y}] = Empty{}
				max = grid.squares[y][x]
			}
		}
	}

	// from left
	for y := 0; y < grid.height; y++ {
		var max uint8 = 0
		for x := 0; x < grid.width; x++ {
			if x == 0 || grid.squares[y][x] > max {
				trees[Pos{x: x, y: y}] = Empty{}
				max = grid.squares[y][x]
			}
		}
	}

	// from right
	for y := 0; y < grid.height; y++ {
		var max uint8 = 0
		for x := grid.width - 1; x >= 0; x-- {
			if x == grid.height-1 || grid.squares[y][x] > max {
				trees[Pos{x: x, y: y}] = Empty{}
				max = grid.squares[y][x]
			}
		}
	}

	return len(trees)
}

func scenic_score(grid Grid, pos Pos) int {
	score := 1
	tree_height := grid.squares[pos.y][pos.x]

	// up
	distance := 0
	for y := pos.y - 1; y >= 0; y-- {
		distance++
		if grid.squares[y][pos.x] >= tree_height {
			break
		}
	}
	if distance > 0 {
		score = score * distance
	}

	// down
	distance = 0
	for y := pos.y + 1; y < grid.height; y++ {
		distance++
		if grid.squares[y][pos.x] >= tree_height {
			break
		}
	}
	if distance > 0 {
		score = score * distance
	}

	// left
	distance = 0
	for x := pos.x - 1; x >= 0; x-- {
		distance++
		if grid.squares[pos.y][x] >= tree_height {
			break
		}
	}
	if distance > 0 {
		score = score * distance
	}

	// right
	distance = 0
	for x := pos.x + 1; x < grid.width; x++ {
		distance++
		if grid.squares[pos.y][x] >= tree_height {
			break
		}
	}
	if distance > 0 {
		score = score * distance
	}

	return score
}

func highest_scenic_score(grid Grid) int {
	max_score := 0
	for y := 1; y < grid.height-1; y++ {
		for x := 1; x < grid.width-1; x++ {
			score := scenic_score(grid, Pos{x: x, y: y})
			if score > max_score {
				max_score = score
			}
		}
	}
	return max_score
}

func main() {
	data := load_data("input.txt")
	grid := make_grid(data)
	trees := visible(grid)
	fmt.Println(trees)

	score := highest_scenic_score(grid)
	fmt.Println(score)
}
