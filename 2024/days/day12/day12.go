package day12

import (
	"fmt"
	"loader"
)

var dirs = [][]int{
	[]int{0, -1},
	[]int{1, 0},
	[]int{0, 1},
	[]int{-1, 0},
}
var size int

type Pos struct {
	x int
	y int
}

func find_region(grid [][]byte, start_pos Pos) int {
	queue := []Pos{start_pos}
	visited := map[Pos]struct{}{}
	visited[start_pos] = struct{}{}

	plant := grid[start_pos.y][start_pos.x]
	region_size := 0
	perimeter := 0

	for len(queue) > 0 {
		pos := queue[0]
		queue = queue[1:]

		region_size++

		for _, dir := range dirs {
			new_x := pos.x + dir[0]
			new_y := pos.y + dir[1]
			if new_x < 0 || new_x == size || new_y < 0 || new_y == size {
				perimeter++
				continue
			}
			if grid[new_y][new_x] != plant {
				perimeter++
				continue
			}
			next_pos := Pos{x: new_x, y: new_y}
			if _, ok := visited[next_pos]; ok {
				continue
			}
			queue = append(queue, next_pos)
			visited[next_pos] = struct{}{}
		}
	}

	for pos := range visited {
		grid[pos.y][pos.x] = '.'
	}

	return region_size * perimeter
}

func price(data []string) int {
	grid := make([][]byte, len(data))
	for i, line := range data {
		grid[i] = []byte(line)
	}
	size = len(grid)

	total := 0
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			if grid[y][x] != '.' {
				total += find_region(grid, Pos{x: x, y: y})
			}
		}
	}
	return total
}

func Run() {
	loader.Day = 12
	data := loader.GetStrings()

	part1 := price(data)
	part2 := -1

	fmt.Printf("%d %d\n", part1, part2)
}
