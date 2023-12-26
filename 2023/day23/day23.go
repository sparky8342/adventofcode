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
	steps int
}

type Node struct {
	pos   Pos
	dist  int
	edges []*Edge
	end   bool
}

type Edge struct {
	distance int
	node     *Node
	blocked  bool
}

type Dir struct {
	dx    int
	dy    int
	arrow byte
}

var height, width int

var dirs []Dir

func init() {
	dirs = []Dir{
		Dir{dx: 0, dy: 1, arrow: 'v'},
		Dir{dx: 0, dy: -1, arrow: '^'},
		Dir{dx: 1, dy: 0, arrow: '>'},
		Dir{dx: -1, dy: 0, arrow: '<'},
	}
}

func load_data(filename string) []string {
	data, _ := ioutil.ReadFile(filename)
	if data[len(data)-1] == '\n' {
		data = data[:len(data)-1]
	}
	return strings.Split(string(data), "\n")
}

func find_nodes(data []string) *Node {
	//TODO find all the nodes and connections in one pass

	// find nodes (points where the path branches)
	start_pos := Pos{x: 1, y: 0}
	start_node := &Node{pos: start_pos}
	end_pos := Pos{x: width - 2, y: height - 1}
	end_node := &Node{pos: end_pos, end: true}
	nodes := []*Node{start_node}
	nodes_map := map[Pos]*Node{start_pos: start_node, end_pos: end_node}

	for y := 1; y < height-1; y++ {
		for x := 1; x < width-1; x++ {

			if data[y][x] == '#' {
				continue
			}

			neighbours := 0
			if data[y+1][x] != '#' {
				neighbours++
			}
			if data[y-1][x] != '#' {
				neighbours++
			}
			if data[y][x+1] != '#' {
				neighbours++
			}
			if data[y][x-1] != '#' {
				neighbours++
			}

			if neighbours >= 3 {
				pos := Pos{x: x, y: y}
				node := &Node{pos: pos}
				nodes = append(nodes, node)
				nodes_map[pos] = node
			}
		}
	}

	// link them together
	for _, node := range nodes {
		visited := map[Pos]struct{}{
			Pos{x: 1, y: 0}: struct{}{},
		}

		queue := []QueueEntry{QueueEntry{pos: node.pos}}

		for len(queue) > 0 {
			qe := queue[0]
			queue = queue[1:]

			pos := qe.pos

			if n, exists := nodes_map[pos]; exists && !(pos.x == node.pos.x && pos.y == node.pos.y) {
				node.edges = append(node.edges, &Edge{distance: qe.steps, node: n})
				n.edges = append(n.edges, &Edge{distance: qe.steps, node: node, blocked: true})
				continue
			}

			for _, dir := range dirs {
				next_pos := Pos{x: pos.x + dir.dx, y: pos.y + dir.dy}
				if _, exists := visited[next_pos]; !exists {
					if next_pos.y >= 0 && (data[next_pos.y][next_pos.x] == '.' || data[next_pos.y][next_pos.x] == dir.arrow) {
						queue = append(queue, QueueEntry{pos: next_pos, steps: qe.steps + 1})
						visited[next_pos] = struct{}{}
					}
				}
			}
		}
	}

	return start_node
}

func pp(node *Node) {
	nodes := []*Node{node}
	for len(nodes) > 0 {
		n := nodes[0]
		nodes = nodes[1:]

		for _, edge := range n.edges {
			fmt.Println(n.pos, n.dist, edge.distance, edge.node.pos)
			nodes = append(nodes, edge.node)
		}
	}
}

func dfs(node *Node, visited map[*Node]struct{}, distance int, max_distance *int, ignore_blocked bool) {
	if node.end {
		if distance > *max_distance {
			*max_distance = distance
		}
		return
	}

	for _, edge := range node.edges {
		if ignore_blocked && edge.blocked {
			continue
		}
		next := edge.node
		if _, exists := visited[next]; !exists {
			visited[next] = struct{}{}
			dfs(edge.node, visited, distance+edge.distance, max_distance, ignore_blocked)
			delete(visited, next)
		}
	}
}

func find_path(data []string) (int, int) {
	height = len(data)
	width = len(data[0])
	start_node := find_nodes(data)

	visited := map[*Node]struct{}{}
	longest := 0
	dfs(start_node, visited, 0, &longest, true)

	longest_any := 0
	dfs(start_node, visited, 0, &longest_any, false)

	return longest, longest_any
}

func main() {
	data := load_data("input.txt")
	fmt.Println(find_path(data))
}
