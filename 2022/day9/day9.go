package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Command struct {
	dir   byte
	steps int
}

type Pos struct {
	x int
	y int
}

type Dir struct {
	x int
	y int
}

type Empty struct {
}

var dirs map[byte]Dir

func init() {
	dirs = map[byte]Dir{
		'U': Dir{x: 0, y: -1},
		'D': Dir{x: 0, y: +1},
		'L': Dir{x: -1, y: 0},
		'R': Dir{x: +1, y: 0},
	}
}

func load_data(filename string) []string {
	data, _ := ioutil.ReadFile(filename)
	if data[len(data)-1] == '\n' {
		data = data[:len(data)-1]
	}
	return strings.Split(string(data), "\n")
}

func parse_data(data []string) []Command {
	commands := []Command{}
	for _, line := range data {
		parts := strings.Split(line, " ")
		steps, _ := strconv.Atoi(parts[1])
		commands = append(commands, Command{dir: line[0], steps: steps})
	}
	return commands
}

func abs(n int) int {
	if n < 0 {
		return n * -1
	} else {
		return n
	}
}

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func move_tail(head *Pos, tail *Pos) {
	// close enough
	if abs(head.x-tail.x) <= 1 && abs(head.y-tail.y) <= 1 {
		return
	}

	tail.x += min(max(head.x-tail.x, -1), 1)
	tail.y += min(max(head.y-tail.y, -1), 1)
}

func run_commands(commands []Command, knots int) int {
	visited := map[Pos]Empty{}

	knots++
	rope := make([]Pos, knots)
	for i := 0; i < knots; i++ {
		rope[i] = Pos{x: 0, y: 0}
	}
	visited[rope[len(rope)-1]] = Empty{}

	for _, command := range commands {
		dir := dirs[command.dir]

		for i := 0; i < command.steps; i++ {
			rope[0].x += dir.x
			rope[0].y += dir.y

			for i := 0; i < knots-1; i++ {
				move_tail(&rope[i], &rope[i+1])
			}

			visited[rope[len(rope)-1]] = Empty{}
		}

	}

	return len(visited)
}

func main() {
	data := load_data("input.txt")
	commands := parse_data(data)

	squares := run_commands(commands, 1)
	fmt.Println(squares)

	squares = run_commands(commands, 9)
	fmt.Println(squares)
}
