package day23

import (
	"fmt"
	"loader"
	"sort"
	"strings"
)

type Node struct {
	name  string
	nodes []*Node
}

func parse_data(data []string) map[string]*Node {
	nodes_map := map[string]*Node{}

	for _, line := range data {
		node1 := line[0:2]
		node2 := line[3:5]

		if _, ok := nodes_map[node1]; !ok {
			nodes_map[node1] = &Node{name: node1}
		}
		if _, ok := nodes_map[node2]; !ok {
			nodes_map[node2] = &Node{name: node2}
		}

		nodes_map[node1].nodes = append(nodes_map[node1].nodes, nodes_map[node2])
		nodes_map[node2].nodes = append(nodes_map[node2].nodes, nodes_map[node1])
	}

	return nodes_map
}

func dfs(source string, node *Node, depth int, path []string, paths map[string]struct{}) {
	if depth == 3 {
		if node.name == source {
			sort.Slice(path, func(i, j int) bool {
				return path[i] < path[j]
			})
			paths[strings.Join(path, "")] = struct{}{}
		}
		return
	}

	for _, linked_node := range node.nodes {
		cpy := make([]string, len(path))
		copy(cpy, path)
		cpy = append(cpy, linked_node.name)
		dfs(source, linked_node, depth+1, cpy, paths)
	}
}

func find_three(nodes map[string]*Node) int {
	paths := map[string]struct{}{}

	for node_name, node := range nodes {
		if node_name[0] == 't' {
			dfs(node_name, node, 0, []string{}, paths)
		}
	}

	return len(paths)
}

func Run() {
	loader.Day = 23
	data := loader.GetStrings()
	nodes := parse_data(data)

	part1 := find_three(nodes)
	part2 := -1

	fmt.Printf("%d %d\n", part1, part2)
}
