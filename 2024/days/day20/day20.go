package day20

import (
	"fmt"
	"loader"
)

type Pos struct {
	x int
	y int
}

type State struct {
	pos   Pos
	time  int
	cheat Pos
}

type Entry struct {
	x     int
	y     int
	cheat Pos
}

var dirs = [][]int{
	[]int{0, -1},
	[]int{1, 0},
	[]int{0, 1},
	[]int{-1, 0},
}

var size int

func find_cheats(data []string, min_save int) int {
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

	times := bfs(grid, start, end, -1)
	var normal_time int
	for k := range times {
		normal_time = k
	}

	savings := 0
	for i := 0; i <= normal_time; i++ {
		times := bfs(grid, start, end, i)
		for k, v := range times {
			saved := normal_time - k
			if saved >= min_save {
				savings += v
			}
		}
	}

	return savings
}

func bfs(grid [][]byte, start Pos, end Pos, cheat_time int) map[int]int {
	start_state := State{pos: start}
	visited := map[Entry]struct{}{}
	visited[Entry{x: start.x, y: start.y, cheat: Pos{x: -1, y: -1}}] = struct{}{}

	queue := []State{start_state}

	times := map[int]int{}

	for len(queue) > 0 {
		state := queue[0]
		queue = queue[1:]

		if state.pos == end {
			times[state.time]++
		}

		for _, dir := range dirs {
			cheat := state.cheat

			pos := Pos{x: state.pos.x + dir[0], y: state.pos.y + dir[1]}
			if pos == start {
				continue
			}

			if pos.x < 0 || pos.x == size || pos.y < 0 || pos.y == size {
				continue
			}
			if grid[pos.y][pos.x] == '#' {
				if state.time == cheat_time {
					cheat = pos
				} else {
					continue
				}
			}

			entry := Entry{x: pos.x, y: pos.y, cheat: cheat}
			if _, ok := visited[entry]; ok {
				continue
			}
			visited[entry] = struct{}{}

			queue = append(queue, State{pos: pos, time: state.time + 1, cheat: cheat})
		}
	}

	return times
}

func Run() {
	loader.Day = 20
	data := loader.GetStrings()

	part1 := find_cheats(data, 100)
	part2 := -1

	fmt.Printf("%d %d\n", part1, part2)
}
