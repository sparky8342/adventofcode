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

func parse_data(data []string) ([]*Grid, int, Pos, Pos) {
	grid := Grid{height: len(data), width: len(data[0]), lookup: map[Pos]Empty{}}
	var start_pos, end_pos Pos
	for y, line := range data {
		for x, ru := range line {
			if y == 0 && ru == '.' {
				start_pos = Pos{x: x, y: y}
			} else if y == grid.height-1 && ru == '.' {
				end_pos = Pos{x: x, y: y}
			} else if ru == '<' || ru == '>' || ru == '^' || ru == 'v' {
				pos := Pos{x: x, y: y}
				grid.blizzards = append(grid.blizzards, Blizzard{pos: pos, dir: ru})
				grid.lookup[pos] = Empty{}
			}
		}
	}
	cycle := lcm(grid.height-2, grid.width-2)

	grids := []*Grid{&grid}
	g := &grid
	for i := 0; i < cycle; i++ {
		g = g.next_grid()
		grids = append(grids, g)
	}

	return grids, cycle, start_pos, end_pos
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

func bfs(grids []*Grid, cycle int, start_pos Pos, end_pos Pos, start_time int) int {
	start := State{pos: start_pos, time: start_time}
	queue := []State{start}
	visited := map[State]Empty{}
	visited[start] = Empty{}

	for len(queue) > 0 {
		state := queue[0]
		queue = queue[1:]

		for _, dir := range dirs {
			new_x := state.pos.x + dir.dx
			new_y := state.pos.y + dir.dy
			if new_x == end_pos.x && new_y == end_pos.y {
				return state.time + 1
			}
			if new_x < 1 || new_x > grids[0].width-2 || new_y < 1 || new_y > grids[0].height-2 {
				continue
			}
			new_pos := Pos{x: new_x, y: new_y}

			_, exists := grids[(state.time+1)%cycle].lookup[new_pos]
			if !exists {
				new_state := State{pos: new_pos, time: state.time + 1}
				if _, exists := visited[new_state]; exists {
					continue
				}
				visited[new_state] = Empty{}
				queue = append(queue, new_state)
			}
		}
		_, exists := grids[(state.time+1)%cycle].lookup[state.pos]
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
	grids, cycle, start, end := parse_data(data)

	time := bfs(grids, cycle, start, end, 0)
	fmt.Println(time)

	time2 := bfs(grids, cycle, end, start, time)
	time3 := bfs(grids, cycle, start, end, time2)
	fmt.Println(time3)
}
