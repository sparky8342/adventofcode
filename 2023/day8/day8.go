package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Node struct {
	left  string
	right string
}

func load_data(filename string) []string {
	data, _ := ioutil.ReadFile(filename)
	if data[len(data)-1] == '\n' {
		data = data[:len(data)-1]
	}
	return strings.Split(string(data), "\n")
}

func parse_data(data []string) (string, map[string]Node) {
	instructions := data[0]

	nodes := map[string]Node{}

	for i := 2; i < len(data); i++ {
		name := data[i][0:3]
		left := data[i][7:10]
		right := data[i][12:15]
		nodes[name] = Node{left: left, right: right}
	}

	return instructions, nodes
}

func follow_instructions(instructions string, nodes map[string]Node) int {
	pos := "AAA"

	steps := 0

	for {
		for _, ins := range instructions {
			if ins == 'L' {
				pos = nodes[pos].left
			} else if ins == 'R' {
				pos = nodes[pos].right
			}
			steps++
			if pos == "ZZZ" {
				return steps
			}
		}
	}
}

func main() {
	data := load_data("input.txt")
	instructions, nodes := parse_data(data)
	steps := follow_instructions(instructions, nodes)
	fmt.Println(steps)
}
