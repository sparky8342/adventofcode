package day10

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

func bfs(grid []string, pos Pos) int {
	queue := []Pos{pos}

	found := map[Pos]struct{}{}

	for len(queue) > 0 {
		pos := queue[0]
		queue = queue[1:]

		if grid[pos.y][pos.x] == '9' {
			found[pos] = struct{}{}
			continue
		}

		for _, dir := range dirs {
			new_x := pos.x + dir[0]
			new_y := pos.y + dir[1]
			if new_x < 0 || new_x == size || new_y < 0 || new_y == size {
				continue
			}
			if grid[new_y][new_x] == grid[pos.y][pos.x]+1 {
				queue = append(queue, Pos{x: new_x, y: new_y})
			}
		}
	}

	return len(found)
}

func score(grid []string) int {
	size = len(grid)

	score := 0
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			if grid[y][x] == '0' {
				score += bfs(grid, Pos{x: x, y: y})
			}
		}
	}

	return score
}

func Run() {
	loader.Day = 10
	grid := loader.GetStrings()

	part1 := score(grid)
	part2 := -1

	fmt.Printf("%d %d\n", part1, part2)
}
