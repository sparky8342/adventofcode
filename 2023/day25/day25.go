package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Node struct {
	name  string
	edges []Edge
	label byte
}

type Edge struct {
	left  *Node
	right *Node
}

type Gain struct {
	a        string
	b        string
	max_gain int
}

func load_data(filename string) []string {
	data, _ := ioutil.ReadFile(filename)
	if data[len(data)-1] == '\n' {
		data = data[:len(data)-1]
	}
	return strings.Split(string(data), "\n")
}

func parse_data(data []string) []*Node {
	node_map := map[string]*Node{}

	for _, line := range data {
		parts := strings.Split(line, ": ")
		name := parts[0]

		neighbours := strings.Split(parts[1], " ")

		for _, n := range append([]string{name}, neighbours...) {
			if _, exists := node_map[n]; !exists {
				node_map[n] = &Node{name: n}
			}
		}

		for _, n := range neighbours {
			edge := Edge{left: node_map[name], right: node_map[n]}
			edge2 := Edge{left: node_map[n], right: node_map[name]}
			node_map[name].edges = append(node_map[name].edges, edge)
			node_map[n].edges = append(node_map[n].edges, edge2)
		}
	}

	nodes := []*Node{}
	for _, node := range node_map {
		nodes = append(nodes, node)
	}

	return nodes
}

func get_d_value(node *Node) int {
	d := 0

	for _, edge := range node.edges {
		if edge.right.label != node.label {
			d++
		} else {
			d--
		}
	}

	return d
}

func find_split(nodes []*Node) int {
	// a simplified and hacky version of Kernighan-Lin

	l := len(nodes)

	group_a := map[string]*Node{}
	group_b := map[string]*Node{}

	for i := 0; i < l/2; i++ {
		group_a[nodes[i].name] = nodes[i]
		nodes[i].label = 'a'
	}
	for i := l / 2; i < l; i++ {
		group_b[nodes[i].name] = nodes[i]
		nodes[i].label = 'b'
	}

	for {
		max_d := -1000
		var max_node *Node
		if len(group_a) > int(float64(len(group_b))/float64(1.2)) {
			for _, a := range group_a {
				d := get_d_value(a)
				if d > max_d {
					max_d = d
					max_node = a
				}
			}
		}

		if len(group_b) > int(float64(len(group_a))/float64(1.2)) {
			for _, b := range group_b {
				d := get_d_value(b)
				if d > max_d {
					max_d = d
					max_node = b
				}
			}
		}

		if max_d < 0 {
			return len(group_a) * len(group_b)
		}

		if max_node.label == 'a' {
			delete(group_a, max_node.name)
			max_node.label = 'b'
			group_b[max_node.name] = max_node
		} else {
			delete(group_b, max_node.name)
			max_node.label = 'a'
			group_a[max_node.name] = max_node
		}
	}
}

func main() {
	data := load_data("input.txt")
	nodes := parse_data(data)
	fmt.Println(find_split(nodes))
}
