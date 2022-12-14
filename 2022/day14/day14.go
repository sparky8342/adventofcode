package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Pos struct {
	x int
	y int
}

type Empty struct {
}

type Cave struct {
	points map[Pos]Empty
	max_y  int
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

func NewCave() *Cave {
	c := new(Cave)
	c.points = map[Pos]Empty{}
	return c
}

func (cave *Cave) set_in_between(pos_a Pos, pos_b Pos) {
	x := pos_a.x
	y := pos_a.y

	inc_x := min(max(pos_b.x-pos_a.x, -1), 1)
	inc_y := min(max(pos_b.y-pos_a.y, -1), 1)

	for !(x == pos_b.x && y == pos_b.y) {
		cave.points[Pos{x: x, y: y}] = Empty{}
		x += inc_x
		y += inc_y
	}
}

func (cave *Cave) draw_points(data []string) {
	max_y := 0

	for _, line := range data {
		parts := strings.Split(line, " -> ")
		points := []Pos{}
		for _, part := range parts {
			coords := strings.Split(part, ",")
			x, _ := strconv.Atoi(coords[0])
			y, _ := strconv.Atoi(coords[1])
			points = append(points, Pos{x: x, y: y})

			if y > max_y {
				max_y = y
			}
		}

		for i := 0; i < len(points)-1; i++ {
			cave.set_in_between(points[i], points[i+1])
		}
		cave.points[points[len(points)-1]] = Empty{}
	}

	cave.max_y = max_y
}

func (cave *Cave) drop_one_sand(fall bool) bool {
	x := 500
	y := 0

	for {
		if _, exists := cave.points[Pos{x: x, y: y + 1}]; !exists && y < cave.max_y+1 {
			y++
		} else if _, exists := cave.points[Pos{x: x - 1, y: y + 1}]; !exists && y < cave.max_y+1 {
			x--
			y++
		} else if _, exists := cave.points[Pos{x: x + 1, y: y + 1}]; !exists && y < cave.max_y+1 {
			x++
			y++
		} else if _, exists := cave.points[Pos{x: x, y: y}]; !exists {
			cave.points[Pos{x: x, y: y}] = Empty{}
			return true
		} else {
			return false
		}

		if fall && y > cave.max_y {
			return false
		}
	}
}

func (cave *Cave) drop_sand(fall bool) int {
	sand := 0
	for cave.drop_one_sand(fall) {
		sand++
	}
	return sand
}

func load_data(filename string) []string {
	data, _ := ioutil.ReadFile(filename)
	if data[len(data)-1] == '\n' {
		data = data[:len(data)-1]
	}
	return strings.Split(string(data), "\n")
}

func main() {
	data := load_data("input.txt")
	cave := NewCave()
	cave.draw_points(data)
	sand := cave.drop_sand(true)
	fmt.Println(sand)

	cave = NewCave()
	cave.draw_points(data)
	sand = cave.drop_sand(false)
	fmt.Println(sand)
}
