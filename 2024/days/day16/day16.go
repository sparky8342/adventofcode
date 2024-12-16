package day16

import (
	"fmt"
	"loader"
	"math"
)

type Pos struct {
	x   int
	y   int
	dir byte
}

type State struct {
	pos  Pos
	time int
}

var xy = map[byte][]int{
	'U': []int{0, -1},
	'D': []int{0, 1},
	'L': []int{-1, 0},
	'R': []int{1, 0},
}

var dirs = map[byte][]byte{
	'U': []byte{'U', 'L', 'R'},
	'D': []byte{'D', 'L', 'R'},
	'L': []byte{'L', 'U', 'D'},
	'R': []byte{'R', 'U', 'D'},
}

func dfs(grid []string, state State, visited map[Pos]int, best_time *int) {
	if state.time > *best_time {
		return
	}

	pos := state.pos

	if grid[pos.y][pos.x] == 'E' {
		if state.time < *best_time {
			*best_time = state.time
		}
		return
	}

	for _, dir := range dirs[pos.dir] {
		new_x := pos.x + xy[dir][0]
		new_y := pos.y + xy[dir][1]

		if grid[new_y][new_x] == '#' {
			continue
		}

		new_pos := Pos{x: new_x, y: new_y, dir: dir}
		new_state := State{pos: new_pos, time: state.time + 1}
		if dir != pos.dir {
			new_state.time += 1000
		}

		if val, ok := visited[new_pos]; ok && val < new_state.time {
			continue
		}
		visited[new_pos] = new_state.time

		dfs(grid, new_state, visited, best_time)
	}
}

func score(grid []string) int {
	size := len(grid)

	start_pos := Pos{x: 1, y: size - 2, dir: 'R'}
	start_state := State{pos: start_pos, time: 0}

	best_time := math.MaxInt32
	dfs(grid, start_state, map[Pos]int{}, &best_time)
	return best_time
}

func Run() {
	loader.Day = 16
	grid := loader.GetStrings()

	part1 := score(grid)
	part2 := -1

	fmt.Printf("%d %d\n", part1, part2)
}
