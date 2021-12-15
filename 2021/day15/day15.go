package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

type Grid struct {
	squares [][]int
	height  int
	width   int
}

type Pos struct {
	x    int
	y    int
	risk int
}

type State struct {
	pos  Pos
	risk int
}

func get_data() Grid {
	data, _ := ioutil.ReadFile("input.txt")
	data = data[:len(data)-1]
	lines := strings.Split(string(data), "\n")
	grid := Grid{}
	for _, line := range lines {
		row := []int{}
		for _, num_rune := range line {
			num := int(num_rune - '0')
			row = append(row, num)
		}
		grid.squares = append(grid.squares, row)
	}
	grid.height = len(grid.squares)
	grid.width = len(grid.squares[0])
	return grid
}

func (grid *Grid) get_neighbours(pos *Pos) []Pos {
	neighbours := []Pos{}
	if pos.x > 0 {
		neighbours = append(neighbours, Pos{x: pos.x - 1, y: pos.y, risk: grid.squares[pos.y][pos.x-1]})
	}
	if pos.x < grid.width-1 {
		neighbours = append(neighbours, Pos{x: pos.x + 1, y: pos.y, risk: grid.squares[pos.y][pos.x+1]})
	}
	if pos.y > 0 {
		neighbours = append(neighbours, Pos{x: pos.x, y: pos.y - 1, risk: grid.squares[pos.y-1][pos.x]})
	}
	if pos.y < grid.height-1 {
		neighbours = append(neighbours, Pos{x: pos.x, y: pos.y + 1, risk: grid.squares[pos.y+1][pos.x]})
	}
	return neighbours
}

func (grid *Grid) search() int {
	start_pos := Pos{x: 0, y: 0, risk: grid.squares[0][0]}
	start_state := State{pos: start_pos, risk: 0}
	queue := []State{start_state}

	visited := map[Pos]bool{}
	for len(queue) > 0 {
		state := queue[0]
		queue = queue[1:]

		fmt.Println(state)

		if state.pos.x == grid.width-1 && state.pos.y == grid.height-1 {
			return state.risk
		}

		if _, found := visited[state.pos]; found {
			continue
		}
		visited[state.pos] = true

		neighbours := grid.get_neighbours(&state.pos)
		for _, neighbour := range neighbours {
			new_state := State{pos: neighbour, risk: state.risk + neighbour.risk}
			queue = append(queue, new_state)
		}

		sort.Slice(queue, func(i, j int) bool {
			return queue[i].risk < queue[j].risk
		})
	}
	return 0
}

func main() {
	grid := get_data()
	risk := grid.search()
	fmt.Println(risk)
}
