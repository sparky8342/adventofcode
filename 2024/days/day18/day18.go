package day18

import (
	"fmt"
	"loader"
)

type Pos struct {
	x int
	y int
}

type State struct {
	pos      Pos
	distance int
}

func parse_data(data []string, amount int) map[Pos]struct{} {
	blocks := map[Pos]struct{}{}
	for i := 0; i < amount; i++ {
		var x, y int
		_, err := fmt.Sscanf(data[i], "%d,%d", &x, &y)
		if err != nil {
			panic(err)
		}
		blocks[Pos{x: x, y: y}] = struct{}{}
	}
	return blocks
}

func find_path(size int, blocks map[Pos]struct{}) int {
	dirs := [][]int{
		[]int{0, -1},
		[]int{1, 0},
		[]int{0, 1},
		[]int{-1, 0},
	}

	start := Pos{x: 0, y: 0}
	start_state := State{pos: start}

	visited := map[Pos]struct{}{}
	visited[start] = struct{}{}

	queue := []State{start_state}

	for len(queue) > 0 {
		state := queue[0]
		queue = queue[1:]

		if state.pos.x == size-1 && state.pos.y == size-1 {
			return state.distance
		}

		for _, dir := range dirs {
			new_x := state.pos.x + dir[0]
			new_y := state.pos.y + dir[1]
			if new_x < 0 || new_x == size || new_y < 0 || new_y == size {
				continue
			}
			pos := Pos{x: new_x, y: new_y}
			if _, ok := visited[pos]; ok {
				continue
			}
			if _, ok := blocks[pos]; ok {
				continue
			}
			visited[pos] = struct{}{}
			queue = append(queue, State{pos: pos, distance: state.distance + 1})
		}
	}

	return -1
}

func Run() {
	loader.Day = 18
	data := loader.GetStrings()

	blocks := parse_data(data, 1024)
	part1 := find_path(71, blocks)

	part2 := -1

	fmt.Printf("%d %d\n", part1, part2)
}
