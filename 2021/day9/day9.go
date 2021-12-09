package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

type Grid struct {
	squares [][]int
	width   int
	height  int
}

type Pos struct {
	x   int
	y   int
	val int
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

func get_neighbours(grid *Grid, pos *Pos) []*Pos {
	neighbours := []*Pos{}
	if pos.x > 0 {
		neighbours = append(neighbours, &Pos{x: pos.x - 1, y: pos.y, val: grid.squares[pos.y][pos.x-1]})
	}
	if pos.x < grid.width-1 {
		neighbours = append(neighbours, &Pos{x: pos.x + 1, y: pos.y, val: grid.squares[pos.y][pos.x+1]})
	}
	if pos.y > 0 {
		neighbours = append(neighbours, &Pos{x: pos.x, y: pos.y - 1, val: grid.squares[pos.y-1][pos.x]})
	}
	if pos.y < grid.height-1 {
		neighbours = append(neighbours, &Pos{x: pos.x, y: pos.y + 1, val: grid.squares[pos.y+1][pos.x]})
	}
	return neighbours
}

func lower_than_surrounding(grid *Grid, pos *Pos) bool {
	neighbours := get_neighbours(grid, pos)
	higher_neighbour := false
	for _, neighbour := range neighbours {
		if grid.squares[neighbour.y][neighbour.x] < pos.val {
			return false
		} else if grid.squares[neighbour.y][neighbour.x] > pos.val {
			higher_neighbour = true
		}
	}
	return higher_neighbour
}

func floodfill(grid *Grid, pos *Pos) int {
	queue := []*Pos{pos}
	visited := map[Pos]bool{}

	size := 0

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]

		if p.val == 9 {
			continue
		}

		if _, seen := visited[*p]; seen {
			continue
		}
		visited[*p] = true

		size++

		neighbours := get_neighbours(grid, p)
		queue = append(queue, neighbours...)
	}

	return size
}

func main() {
	grid := get_data()

	// part 1
	low_points := []Pos{}
	risk_sum := 0
	for y := 0; y < grid.height; y++ {
		for x := 0; x < grid.width; x++ {
			pos := Pos{x: x, y: y, val: grid.squares[y][x]}
			if lower_than_surrounding(&grid, &pos) {
				low_points = append(low_points, pos)
				risk_sum += pos.val + 1
			}
		}
	}
	fmt.Println(risk_sum)

	// part 2
	sizes := []int{}
	for _, point := range low_points {
		size := floodfill(&grid, &point)
		sizes = append(sizes, size)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(sizes)))
	fmt.Println(sizes[0] * sizes[1] * sizes[2])
}
