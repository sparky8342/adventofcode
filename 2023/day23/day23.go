package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Pos struct {
	x int
	y int
}

type QueueEntry struct {
	pos   Pos
	node  *Node
	steps int
}

type Node struct {
	pos   Pos
	edges []*Edge
}

type Edge struct {
	distance int
	node     *Node
}

var height, width int

func load_data(filename string) []string {
	data, _ := ioutil.ReadFile(filename)
	if data[len(data)-1] == '\n' {
		data = data[:len(data)-1]
	}
	return strings.Split(string(data), "\n")
}

func find_nodes(data []string) {
	start_pos := Pos{x: 1, y: 1}
	start := &Node{pos: start_pos}
	queue := []QueueEntry{QueueEntry{pos: start_pos, node: start}}

	visited := map[Pos]struct{}{
		Pos{x: 1, y: 0}: struct{}{},
	}

	for len(queue) > 0 {
		qe := queue[0]
		queue = queue[1:]

		pos := qe.pos

		if _, exists := visited[pos]; exists {
			continue
		}
		visited[pos] = struct{}{}

		if pos.y == height-1 && pos.x == width-2 {
			qe.node.edges = append(qe.node.edges, &Edge{distance: qe.steps, node: &Node{pos: pos}})
			continue
		}

		neighbours := []Pos{
			Pos{x: pos.x, y: pos.y + 1},
			Pos{x: pos.x, y: pos.y - 1},
			Pos{x: pos.x + 1, y: pos.y},
			Pos{x: pos.x - 1, y: pos.y},
		}
		valid_neighbours := []Pos{}
		for _, neighbour := range neighbours {
			if _, exists := visited[neighbour]; !exists && data[neighbour.y][neighbour.x] != '#' {
				valid_neighbours = append(valid_neighbours, neighbour)
			}
		}

		if len(valid_neighbours) > 1 {
			node := &Node{pos: pos}
			qe.node.edges = append(qe.node.edges, &Edge{distance: qe.steps, node: node})
			for _, neighbour := range valid_neighbours {
				queue = append(queue, QueueEntry{pos: neighbour, node: node, steps: 0})
			}
		} else if len(valid_neighbours) == 1 {
			queue = append(queue, QueueEntry{pos: valid_neighbours[0], node: qe.node, steps: qe.steps + 1})
		}
	}

	// debug output
	nodes := []*Node{start}
	for len(nodes) > 0 {
		n := nodes[0]
		nodes = nodes[1:]

		for _, edge := range n.edges {
			fmt.Println(n.pos, edge.distance, edge.node.pos)
			nodes = append(nodes, edge.node)
		}
	}

}

func find_path(data []string) {
	height = len(data)
	width = len(data[0])
	find_nodes(data)
}

func main() {
	data := load_data("input.txt")
	find_path(data)
}
