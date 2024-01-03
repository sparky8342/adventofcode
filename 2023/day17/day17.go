package main

import (
	"container/heap"
	"fmt"
	"io/ioutil"
	"strings"
)

type PriorityQueue []*QueueEntry

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].distance < pq[j].distance
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*QueueEntry)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return item
}

var dirs [4]Dir

type Pos struct {
	x int
	y int
}

type State struct {
	pos   Pos
	north int
	east  int
	south int
	west  int
}

type QueueEntry struct {
	state    State
	distance int
}

type Dir struct {
	dx   int
	dy   int
	name byte
}

func init() {
	dirs = [4]Dir{
		Dir{dx: 1, dy: 0, name: 'E'},
		Dir{dx: -1, dy: 0, name: 'W'},
		Dir{dx: 0, dy: 1, name: 'S'},
		Dir{dx: 0, dy: -1, name: 'N'},
	}
}

func load_data(filename string) []string {
	data, _ := ioutil.ReadFile(filename)
	if data[len(data)-1] == '\n' {
		data = data[:len(data)-1]
	}
	return strings.Split(string(data), "\n")
}

func find_path(data []string, ultra bool) int {
	height := len(data)
	width := len(data[0])

	min, max := 0, 3
	if ultra {
		min, max = 4, 10
	}

	queue := make(PriorityQueue, 1)
	queue[0] = &QueueEntry{state: State{pos: Pos{x: 0, y: 0}}}
	heap.Init(&queue)

	visited := map[State]struct{}{}

	for queue.Len() > 0 {
		qe := heap.Pop(&queue).(*QueueEntry)
		state := qe.state

		if _, exists := visited[state]; exists {
			continue
		}
		visited[state] = struct{}{}

		if state.pos.x == width-1 && state.pos.y == height-1 {
			return qe.distance
		}

		for _, dir := range dirs {
			x := state.pos.x + dir.dx
			y := state.pos.y + dir.dy
			if x < 0 || x >= width || y < 0 || y >= height {
				continue
			}

			new_pos := Pos{x: x, y: y}
			new_state := State{pos: new_pos}

			switch dir.name {
			case 'N':
				if (state.east > 0 && state.east < min) || (state.west > 0 && state.west < min) || state.north == max || state.south > 0 {
					continue
				}
				new_state.north = state.north + 1
			case 'E':
				if (state.north > 0 && state.north < min) || (state.south > 0 && state.south < min) || state.east == max || state.west > 0 {
					continue
				}
				new_state.east = state.east + 1
			case 'S':
				if (state.east > 0 && state.east < min) || (state.west > 0 && state.west < min) || state.south == max || state.north > 0 {
					continue
				}
				new_state.south = state.south + 1
			case 'W':
				if (state.north > 0 && state.north < min) || (state.south > 0 && state.south < min) || state.west == max || state.east > 0 {
					continue
				}
				new_state.west = state.west + 1
			}

			if _, exists := visited[new_state]; exists {
				continue
			}

			heap.Push(&queue, &QueueEntry{state: new_state, distance: qe.distance + int(data[y][x]-'0')})
		}
	}

	return -1
}

func main() {
	data := load_data("input.txt")
	fmt.Println(find_path(data, false))
	fmt.Println(find_path(data, true))
}
