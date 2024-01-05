package main

import (
	"fmt"
	"io/ioutil"
	"math/big"
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

var height, width int

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

func mod(a, b int) int {
	return (a%b + b) % b
}

func plots(data []string, amount int) int {
	height = len(data)
	width = len(data[0])

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
			if data[mod(n.y, height)][mod(n.x, width)] == '#' {
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

func large_steps(data []string, amount int) *big.Int {
	a := plots(data, 65)
	b := plots(data, 65+height)
	c := plots(data, 65+height*2)

	diff1 := b - a
	diff2 := c - b

	amount_big := big.NewInt(int64(amount))
	diff1_big := big.NewInt(int64(diff1))
	diff2_big := big.NewInt(int64(diff2))

	sum := new(big.Int)
	sum.Sub(amount_big, big.NewInt(int64(1)))
	sum.Mul(sum, amount_big)
	sum.Div(sum, big.NewInt(int64(2)))
	sum.Mul(sum, new(big.Int).Sub(diff2_big, diff1_big))
	sum.Add(sum, new(big.Int).Mul(diff1_big, amount_big))
	sum.Add(sum, big.NewInt(int64(a)))

	return sum
}

func main() {
	data := load_data("input.txt")
	fmt.Println(plots(data, 64))
	fmt.Println(large_steps(data, 26501365/height))
}
