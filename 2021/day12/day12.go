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
	node    *Node
	visited map[string]bool
	path    []string
}

func (state *State) copy_state() State {
	cp := State{node: state.node}

	cp.visited = map[string]bool{}
	for k, _ := range state.visited {
		cp.visited[k] = true
	}

	cp.path = make([]string, len(state.path))
	copy(cp.path, state.path)
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

	for _, line := range lines {
		parts := strings.Split(line, "-")
		for _, part := range parts {
			if _, found := nodes[part]; !found {
				big := false
				if all_caps(part) {
					big = true
				}
				nodes[part] = &Node{name: part, big: big}

			}
		}
	}

	for _, line := range lines {
		parts := strings.Split(line, "-")
		left := nodes[parts[0]]
		right := nodes[parts[1]]
		left.neighbours = append(left.neighbours, right)
		right.neighbours = append(right.neighbours, left)
	}

	return nodes["start"]
}

func walk(start *Node) int {
	queue := []State{State{node: start, visited: map[string]bool{"start": true}, path: []string{"start"}}}

	path_count := 0

	for len(queue) > 0 {
		state := queue[0]
		queue = queue[1:]

		if state.node.name == "end" {
			path_count++
			continue
		}

		for _, neighbour := range state.node.neighbours {
			if _, found := state.visited[neighbour.name]; !found {
				new_state := state.copy_state()
				new_state.node = neighbour
				if !neighbour.big {
					new_state.visited[neighbour.name] = true
				}
				new_state.path = append(new_state.path, neighbour.name)
				queue = append(queue, new_state)
			}
		}
	}

	return path_count
}

func main() {
	start := get_data()

	path_count := walk(start)
	fmt.Println(path_count)
}
