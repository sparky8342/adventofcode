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

type Blizzard struct {
	pos Pos
	dir rune
}

type Grid struct {
	width     int
	height    int
	blizzards []Blizzard
	lookup    map[Pos]Empty
	start     Pos
	end       Pos
	cycle     int
}

type State struct {
	pos  Pos
	time int
}

type Dir struct {
	dx int
	dy int
}

type Empty struct {
}

var dirs []Dir

func init() {
	dirs = []Dir{
		Dir{dx: 0, dy: 1},
		Dir{dx: 0, dy: -1},
		Dir{dx: 1, dy: 0},
		Dir{dx: -1, dy: 0},
	}
}

func load_data(filename string) []string {
	data, _ := ioutil.ReadFile(filename)
	if data[len(data)-1] == '\n' {
		data = data[:len(data)-1]
	}
	return strings.Split(string(data), "\n")
}

func parse_data(data []string) Grid {
	grid := Grid{height: len(data), width: len(data[0]), lookup: map[Pos]Empty{}}
	for y, line := range data {
		for x, ru := range line {
			if y == 0 && ru == '.' {
				grid.start = Pos{x: x, y: y}
			} else if y == grid.height-1 && ru == '.' {
				grid.end = Pos{x: x, y: y}
			} else if ru == '<' || ru == '>' || ru == '^' || ru == 'v' {
				pos := Pos{x: x, y: y}
				grid.blizzards = append(grid.blizzards, Blizzard{pos: pos, dir: ru})
				grid.lookup[pos] = Empty{}
			}
		}
	}
	return grid
}

func gcd(a int, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func lcm(a int, b int) int {
	return a * b / gcd(a, b)
}

func (g *Grid) next_grid() *Grid {
	grid := new(Grid)
	grid.width = g.width
	grid.height = g.height
	grid.start = g.start
	grid.end = g.end
	grid.lookup = map[Pos]Empty{}

	for _, blizzard := range g.blizzards {
		new_pos := Pos{x: blizzard.pos.x, y: blizzard.pos.y}
		switch blizzard.dir {
		case '>':
			new_pos.x++
			if new_pos.x == g.width-1 {
				new_pos.x = 1
			}
		case '<':
			new_pos.x--
			if new_pos.x == 0 {
				new_pos.x = g.width - 2
			}
		case 'v':
			new_pos.y++
			if new_pos.y == g.height-1 {
				new_pos.y = 1
			}
		case '^':
			new_pos.y--
			if new_pos.y == 0 {
				new_pos.y = g.height - 2
			}
		}
		grid.blizzards = append(grid.blizzards, Blizzard{pos: new_pos, dir: blizzard.dir})
		grid.lookup[new_pos] = Empty{}
	}

	return grid
}

func bfs(grid *Grid) int {
	grids := []*Grid{grid}

	start := State{pos: grid.start, time: 0}
	queue := []State{start}
	visited := map[State]Empty{}
	visited[start] = Empty{}

	for len(queue) > 0 {
		state := queue[0]
		queue = queue[1:]

		if state.time >= len(grids)-1 {
			grids = append(grids, grids[len(grids)-1].next_grid())
		}

		for _, dir := range dirs {
			new_x := state.pos.x + dir.dx
			new_y := state.pos.y + dir.dy
			if new_x == grid.end.x && new_y == grid.end.y {
				return state.time + 1
			}
			if new_x < 1 || new_x > grid.width-2 || new_y < 1 || new_y > grid.height-2 {
				continue
			}
			new_pos := Pos{x: new_x, y: new_y}

			_, exists := grids[state.time+1].lookup[new_pos]
			if !exists {
				new_state := State{pos: new_pos, time: state.time + 1}
				if _, exists := visited[new_state]; exists {
					continue
				}
				visited[new_state] = Empty{}
				queue = append(queue, new_state)
			}
		}
		_, exists := grids[state.time+1].lookup[state.pos]
		if !exists {
			new_state := State{pos: state.pos, time: state.time + 1}
			if _, exists := visited[new_state]; exists {
				continue
			}
			visited[new_state] = Empty{}
			queue = append(queue, new_state)
		}

	}
	return -1
}

func (grid *Grid) print_grid(pos_x int, pos_y int) {
	for y := 0; y < grid.height; y++ {
		for x := 0; x < grid.width; x++ {
			if x == pos_x && y == pos_y {
				fmt.Print("o")
			} else if _, exists := grid.lookup[Pos{x: x, y: y}]; exists {
				count := 0
				var dir rune
				for _, blizzard := range grid.blizzards {
					if blizzard.pos.x == x && blizzard.pos.y == y {
						count++
						dir = blizzard.dir
					}
				}
				if count == 1 {
					fmt.Print(string(dir))
				} else {
					fmt.Print(count)
				}
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func main() {
	data := load_data("input.txt")
	grid := parse_data(data)
	fmt.Println(bfs(&grid))
}
