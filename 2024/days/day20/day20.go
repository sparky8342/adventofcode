package day20

import (
	"fmt"
	"loader"
	"utils"
)

type Pos struct {
	x int
	y int
}

type State struct {
	pos  Pos
	time int
}

var dirs = [][]int{
	[]int{0, -1},
	[]int{1, 0},
	[]int{0, 1},
	[]int{-1, 0},
}

var size int

func find_cheats(data []string, min_save int, cheat_dist int) int {
	size = len(data)

	grid := make([][]byte, size)
	for i, line := range data {
		grid[i] = []byte(line)
	}

	var start, end Pos
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			if grid[y][x] == 'S' {
				start = Pos{x: x, y: y}
			} else if grid[y][x] == 'E' {
				end = Pos{x: x, y: y}
			}
		}
	}

	distances := bfs(grid, start, end)

	savings := 0
	for space, distance := range distances {
		for dy := -cheat_dist; dy <= cheat_dist; dy++ {
			for dx := -cheat_dist; dx <= cheat_dist; dx++ {
				if dx == 0 && dy == 0 {
					continue
				}
				dist := utils.Abs(dx) + utils.Abs(dy)
				if dist > cheat_dist {
					continue
				}
				next_pos := Pos{x: space.x + dx, y: space.y + dy}
				if val, ok := distances[next_pos]; ok {
					if val-(distance+dist) >= min_save {
						savings++
					}
				}
			}
		}
	}

	return savings
}

func bfs(grid [][]byte, start Pos, end Pos) map[Pos]int {
	start_state := State{pos: start}
	visited := map[Pos]struct{}{}
	visited[start] = struct{}{}

	queue := []State{start_state}

	distances := map[Pos]int{}

	for len(queue) > 0 {
		state := queue[0]
		queue = queue[1:]

		distances[state.pos] = state.time

		if state.pos == end {
			break
		}

		for _, dir := range dirs {
			pos := Pos{x: state.pos.x + dir[0], y: state.pos.y + dir[1]}
			if grid[pos.y][pos.x] == '#' {
				continue
			}
			if _, ok := visited[pos]; ok {
				continue
			}
			visited[pos] = struct{}{}

			queue = append(queue, State{pos: pos, time: state.time + 1})
		}
	}

	return distances
}

func Run() {
	loader.Day = 20
	data := loader.GetStrings()

	part1 := find_cheats(data, 100, 2)
	part2 := find_cheats(data, 100, 20)

	fmt.Printf("%d %d\n", part1, part2)
}
