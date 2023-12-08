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

func find_cycle(position string, instructions string, nodes map[string]Node) (int, int) {
	seen := map[string]int{}
	step_count := 0

	for {
		for _, ins := range instructions {
			if ins == 'L' {
				position = nodes[position].left
			} else if ins == 'R' {
				position = nodes[position].right
			}
			step_count++
			if position[2] == 'Z' {
				if val, exists := seen[position]; exists {
					return step_count, step_count - val
				} else {
					seen[position] = step_count
				}
			}
		}
	}
}

func all_same(nums []int) bool {
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[0] {
			return false
		}
	}
	return true
}

func ghost_steps(instructions string, nodes map[string]Node) int {
	positions := []string{}
	for name, _ := range nodes {
		if name[2] == 'A' {
			positions = append(positions, name)
		}
	}

	steps := make([]int, len(positions))
	cycles := make([]int, len(positions))

	for i, position := range positions {
		steps[i], cycles[i] = find_cycle(position, instructions, nodes)
	}

	for {
		smallest := 0
		largest := 0
		for i := 1; i < len(steps); i++ {
			if steps[i] < steps[smallest] {
				smallest = i
			} else if steps[i] > steps[largest] {
				largest = i
			}
		}

		diff := steps[largest] - steps[smallest]
		if diff%cycles[smallest] == 0 {
			steps[smallest] = steps[largest]
		} else {
			steps[smallest] += ((diff / cycles[smallest]) + 1) * cycles[smallest]
		}

		if all_same(steps) {
			return steps[0]
		}
	}
}

func main() {
	data := load_data("input.txt")
	instructions, nodes := parse_data(data)
	steps := follow_instructions(instructions, nodes)
	fmt.Println(steps)
	steps = ghost_steps(instructions, nodes)
	fmt.Println(steps)
}
