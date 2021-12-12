package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Node struct {
	name       string
	neighbours []*Node
	big        bool
}

type State struct {
	node                *Node
	visited             map[string]bool
	small_visited_twice bool
}

func (state *State) copy_state() State {
	cp := State{node: state.node, small_visited_twice: state.small_visited_twice}

	cp.visited = map[string]bool{}
	for k, _ := range state.visited {
		cp.visited[k] = true
	}

	return cp
}

func all_caps(name string) bool {
	for _, ch := range name {
		if ch > 'Z' {
			return false
		}
	}
	return true
}

func get_data() *Node {
	data, _ := ioutil.ReadFile("input.txt")
	data = data[:len(data)-1]
	lines := strings.Split(string(data), "\n")

	nodes := map[string]*Node{}

	pairs := [][2]string{}
	for _, line := range lines {
		pair := strings.Split(line, "-")
		pairs = append(pairs, [2]string{pair[0], pair[1]})
	}

	for _, pair := range pairs {
		for _, name := range pair {
			if _, found := nodes[name]; !found {
				nodes[name] = &Node{name: name, big: all_caps(name)}
			}
		}
	}

	for _, pair := range pairs {
		left := nodes[pair[0]]
		right := nodes[pair[1]]
		left.neighbours = append(left.neighbours, right)
		right.neighbours = append(right.neighbours, left)
	}

	return nodes["start"]
}

func walk(start *Node, twice_visit bool) int {
	queue := []State{State{node: start, visited: map[string]bool{"start": true}}}

	path_count := 0

	for len(queue) > 0 {
		state := queue[0]
		queue = queue[1:]

		if state.node.name == "end" {
			path_count++
			continue
		}

		for _, neighbour := range state.node.neighbours {
			_, found := state.visited[neighbour.name]

			if found && !(twice_visit && neighbour.name != "start" && neighbour.big == false && state.small_visited_twice == false) {
				continue
			}

			new_state := state.copy_state()
			new_state.node = neighbour
			if !neighbour.big {
				new_state.visited[neighbour.name] = true
			}
			if found {
				new_state.small_visited_twice = true
			}
			queue = append(queue, new_state)
		}
	}

	return path_count
}

func main() {
	start := get_data()

	path_count := walk(start, false)
	fmt.Println(path_count)

	path_count = walk(start, true)
	fmt.Println(path_count)
}
