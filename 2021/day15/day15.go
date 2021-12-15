package main

import (
	"fmt"
	"io/ioutil"
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

type Heap struct {
	hslice []State
}

func heap_parent(i int) int {
	return (i - 1) / 2
}

func heap_left(i int) int {
	return 2*i + 1
}

func heap_right(i int) int {
	return 2*i + 2
}

func (heap *Heap) insert(state State) {
	heap.hslice = append(heap.hslice, state)

	i := len(heap.hslice) - 1
	for i != 0 {
		parent := heap_parent(i)
		if heap.hslice[parent].risk > heap.hslice[i].risk {
			heap.hslice[i], heap.hslice[parent] = heap.hslice[parent], heap.hslice[i]
			i = parent
		} else {
			break
		}
	}
}

func (heap *Heap) pop_min() State {
	if len(heap.hslice) == 1 {
		state := heap.hslice[0]
		heap.hslice = []State{}
		return state
	}

	min := heap.hslice[0]
	heap.hslice[0] = heap.hslice[len(heap.hslice)-1]
	heap.hslice = heap.hslice[0 : len(heap.hslice)-1]
	heap.min_heapify(0)
	return min
}

func (heap *Heap) min_heapify(i int) {
	l := heap_left(i)
	r := heap_right(i)
	smallest := i
	if l < len(heap.hslice) && heap.hslice[l].risk < heap.hslice[i].risk {
		smallest = l
	}
	if r < len(heap.hslice) && heap.hslice[r].risk < heap.hslice[smallest].risk {
		smallest = r
	}
	if smallest != i {
		heap.hslice[i], heap.hslice[smallest] = heap.hslice[smallest], heap.hslice[i]
		heap.min_heapify(smallest)
	}
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
	queue := Heap{}
	queue.insert(start_state)

	visited := map[Pos]bool{}
	for len(queue.hslice) > 0 {
		state := queue.pop_min()

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
			queue.insert(new_state)
		}
	}
	return 0
}

func (grid *Grid) expand() Grid {
	new_grid := Grid{height: grid.height * 5, width: grid.width * 5}

	new_grid.squares = make([][]int, new_grid.height)
	for i := range new_grid.squares {
		new_grid.squares[i] = make([]int, new_grid.width)
	}

	// copy initial grid data
	for y := 0; y < grid.height; y++ {
		for x := 0; x < grid.width; x++ {
			new_grid.squares[y][x] = grid.squares[y][x]
		}
	}

	// fill in expanded squares
	for square_y := 0; square_y < 5; square_y++ {
		for square_x := 0; square_x < 5; square_x++ {
			if square_x == 0 && square_y == 0 {
				continue
			}
			for y := 0; y < grid.height; y++ {
				for x := 0; x < grid.width; x++ {
					new_val := grid.squares[y][x] + square_x + square_y
					if new_val > 9 {
						new_val -= 9
					}
					new_grid.squares[square_y*grid.height+y][square_x*grid.width+x] = new_val
				}
			}
		}
	}

	return new_grid
}

func main() {
	grid := get_data()
	risk := grid.search()
	fmt.Println(risk)

	expanded_grid := grid.expand()
	risk = expanded_grid.search()
	fmt.Println(risk)
}
