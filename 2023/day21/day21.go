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

type State struct {
	x     int
	y     int
	steps int
}

func load_data(filename string) []string {
	data, _ := ioutil.ReadFile(filename)
	if data[len(data)-1] == '\n' {
		data = data[:len(data)-1]
	}
	return strings.Split(string(data), "\n")
}

func find_start(data []string) State {
	height := len(data)
	width := len(data[0])

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if data[y][x] == 'S' {
				return State{x: x, y: y}
			}
		}
	}

	return State{}
}

func plots(data []string, amount int) int {
	visited := map[State]struct{}{}
	positions := map[Pos]struct{}{}

	queue := []State{}

	start := find_start(data)
	queue = append(queue, start)
	visited[start] = struct{}{}

	for steps := 2; steps <= amount; steps += 2 {
		state := State{x: start.x, y: start.y, steps: steps}
		queue = append(queue, state)
		visited[state] = struct{}{}
	}

	for len(queue) > 0 {
		state := queue[0]
		queue = queue[1:]

		if state.steps == amount {
			positions[Pos{x: state.x, y: state.y}] = struct{}{}
			continue
		}

		neighbours := []State{
			State{x: state.x + 1, y: state.y, steps: state.steps + 1},
			State{x: state.x - 1, y: state.y, steps: state.steps + 1},
			State{x: state.x, y: state.y + 1, steps: state.steps + 1},
			State{x: state.x, y: state.y - 1, steps: state.steps + 1},
		}

		for _, n := range neighbours {
			if data[n.y][n.x] == '#' {
				continue
			}
			for steps := n.steps; steps <= amount; steps += 2 {
				state := State{x: n.x, y: n.y, steps: steps}
				if _, exists := visited[state]; !exists {
					queue = append(queue, state)
					visited[state] = struct{}{}
				}
			}
		}
	}

	return len(positions)
}

func main() {
	data := load_data("input.txt")
	amount := plots(data, 64)
	fmt.Println(amount)
}
