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

type Dir struct {
	x int
	y int
}

type Node struct {
	elevation rune
	distance  int
}

type Empty struct {
}

var dirs []Dir

func init() {
	dirs = []Dir{
		Dir{x: 0, y: -1},
		Dir{x: 0, y: +1},
		Dir{x: -1, y: 0},
		Dir{x: +1, y: 0},
	}
}

func load_data(filename string) []string {
	data, _ := ioutil.ReadFile(filename)
	if data[len(data)-1] == '\n' {
		data = data[:len(data)-1]
	}
	return strings.Split(string(data), "\n")
}

func parse_data(data []string) (map[Pos]*Node, Pos, Pos) {
	nodes := map[Pos]*Node{}
	var start, end Pos

	for y, line := range data {
		for x, ru := range line {
			pos := Pos{x: x, y: y}
			node := Node{elevation: ru}
			if ru == 'S' {
				node.distance = 0
				node.elevation = 'a'
				start = pos
			} else if ru == 'E' {
				node.elevation = 'z'
				end = pos
			}

			nodes[pos] = &node
		}
	}

	return nodes, start, end
}

func get_neighbours(pos Pos) []Pos {
	neighbours := []Pos{}
	for _, dir := range dirs {
		neighbours = append(neighbours, Pos{pos.x + dir.x, pos.y + dir.y})
	}
	return neighbours
}

func bfs(nodes map[Pos]*Node, start Pos, end Pos) int {
	queue := []Pos{start}
	visited := map[Pos]Empty{}

	for len(queue) > 0 {
		pos := queue[0]
		queue = queue[1:]

		if pos == end {
			return nodes[pos].distance
		}

		if _, seen := visited[pos]; seen {
			continue
		}
		visited[pos] = Empty{}

		current := nodes[pos]
		neighbours := get_neighbours(pos)

		for _, neighbour := range neighbours {
			if _, exists := nodes[neighbour]; exists {
				if _, visited := visited[neighbour]; visited {
					continue
				}
				elevation := nodes[neighbour].elevation
				if current.elevation >= elevation-1 {
					queue = append(queue, neighbour)
					nodes[neighbour].distance = current.distance + 1
				}
			}
		}
	}

	return -1
}

func main() {
	data := load_data("input.txt")

	nodes, start, end := parse_data(data)
	distance := bfs(nodes, start, end)
	fmt.Println(distance)

}
