package day16

import (
	"fmt"
	"loader"
)

type Square struct {
	x int
	y int
}

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

func dfs(grid []string, state State, visited map[Pos]int, best_time *int, path []Square, best_path_squares map[Square]struct{}) {
	if state.time > *best_time {
		return
	}

	pos := state.pos

	if grid[pos.y][pos.x] == 'E' {
		if state.time == *best_time {
			for _, square := range path {
				best_path_squares[square] = struct{}{}
			}
		}
		if state.time < *best_time {
			*best_time = state.time
			// empty best_path_squares
			// for some reason using
			// 'best_path_squares = make(map[Square]struct{})'
			// doesn't work correctly
			for square := range best_path_squares {
				delete(best_path_squares, square)
			}
			for _, square := range path {
				best_path_squares[square] = struct{}{}
			}
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

		square := Square{x: new_x, y: new_y}
		new_path := make([]Square, len(path))
		copy(new_path, path)
		new_path = append(new_path, square)

		dfs(grid, new_state, visited, best_time, new_path, best_path_squares)
	}
}

func print_squares(grid []string, squares map[Square]struct{}) {
	size := len(grid)
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			if grid[y][x] == '#' {
				fmt.Print("#")
			} else {
				sq := Square{x: x, y: y}
				if _, ok := squares[sq]; ok {
					fmt.Print("O")
				} else {
					fmt.Print(" ")
				}
			}
		}
		fmt.Println()
	}
}

func score(grid []string) (int, int) {
	size := len(grid)

	start_pos := Pos{x: 1, y: size - 2, dir: 'R'}
	start_state := State{pos: start_pos, time: 0}

	best_time := 100000
	visited := map[Pos]int{}
	path := []Square{Square{x: start_pos.x, y: start_pos.y}}
	best_path_squares := map[Square]struct{}{}
	dfs(grid, start_state, visited, &best_time, path, best_path_squares)

	print_squares(grid, best_path_squares)
	return best_time, len(best_path_squares)
}

func Run() {
	loader.Day = 16
	grid := loader.GetStrings()

	part1, part2 := score(grid)

	fmt.Printf("%d %d\n", part1, part2)
}
