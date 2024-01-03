package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

var height, width int
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

func find_path(data []string) int {
	height = len(data)
	width = len(data[0])

	start_pos := Pos{x: 0, y: 0}
	start_state := State{pos: start_pos}
	queue := []QueueEntry{QueueEntry{state: start_state}}

	visited := map[State]struct{}{}

	for len(queue) > 0 {
		qe := queue[0]
		queue = queue[1:]

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
				if state.north == 3 || state.south > 0 {
					continue
				}
				new_state.north = state.north + 1
			case 'E':
				if state.east == 3 || state.west > 0 {
					continue
				}
				new_state.east = state.east + 1
			case 'S':
				if state.south == 3 || state.north > 0 {
					continue
				}
				new_state.south = state.south + 1
			case 'W':
				if state.west == 3 || state.east > 0 {
					continue
				}
				new_state.west = state.west + 1
			}

			queue = append(queue, QueueEntry{state: new_state, distance: qe.distance + int(data[y][x]-'0')})
		}

		// TODO use heap
		sort.Slice(queue, func(i, j int) bool {
			return queue[i].distance < queue[j].distance
		})
	}

	return -1
}

func main() {
	data := load_data("input.txt")
	fmt.Println(find_path(data))
}
