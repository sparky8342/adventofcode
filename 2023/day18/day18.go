package main

import (
	"fmt"
	"io/ioutil"
	"math/big"
	"strconv"
	"strings"
)

type Pos struct {
	x int
	y int
}

func load_data(filename string) []string {
	data, _ := ioutil.ReadFile(filename)
	if data[len(data)-1] == '\n' {
		data = data[:len(data)-1]
	}
	return strings.Split(string(data), "\n")
}

func shoelace(vertices []Pos) *big.Int {
	l := len(vertices)

	area := big.NewInt(int64(0))

	j := l - 1
	for i := 0; i < l; i++ {
		i_x := big.NewInt(int64(vertices[i].x))
		i_y := big.NewInt(int64(vertices[i].y))
		j_x := big.NewInt(int64(vertices[j].x))
		j_y := big.NewInt(int64(vertices[j].y))

		var sum_a = new(big.Int)
		sum_a.Add(j_x, i_x)

		var sum_b = new(big.Int)
		sum_b.Sub(j_y, i_y)

		var sum_c = new(big.Int)
		sum_c.Mul(sum_a, sum_b)

		area.Add(area, sum_c)

		j = i
	}

	area.Div(area, big.NewInt(int64(2)))
	area.Abs(area)
	return area
}

func get_size(data []string, use_colours bool) *big.Int {
	x, y := 0, 0

	vertices := []Pos{Pos{x: 0, y: 0}}

	// find all the vertices
	for i := 0; i < len(data)-1; i++ {
		line := data[i]

		var n int
		var direction byte

		parts := strings.Split(line, " ")

		if use_colours {
			num, _ := strconv.ParseInt(parts[2][2:7], 16, 64)
			n = int(num)
			switch parts[2][7] {
			case '0':
				direction = 'R'
			case '1':
				direction = 'D'
			case '2':
				direction = 'L'
			case '3':
				direction = 'U'
			}
		} else {
			direction = parts[0][0]
			n, _ = strconv.Atoi(parts[1])
		}

		switch direction {
		case 'U':
			y = y - n
		case 'R':
			x = x + n
		case 'D':
			y = y + n
		case 'L':
			x = x - n
		}

		vertices = append(vertices, Pos{x: x, y: y})
	}

	// find the vertices actually around the path
	vertices_around := []Pos{}

	vertices = append(vertices, vertices[0])
	vertices = append(vertices, vertices[1])
	for i := 1; i < len(vertices)-1; i++ {
		current := vertices[i]
		prev := vertices[i-1]
		next := vertices[i+1]

		dx, dy := 0, 0
		if prev.x < current.x {
			if current.y < next.y {
				dx = 1
			}
		} else if prev.x > current.x {
			if current.y < next.y {
				dx, dy = 1, 1
			} else {
				dy = 1
			}
		} else if prev.y < current.y {
			if current.x < next.x {
				dx = 1
			} else {
				dx, dy = 1, 1
			}
		} else if prev.y > current.y {
			if current.x > next.x {
				dy = 1
			}
		}

		vertices_around = append(vertices_around, Pos{x: current.x + dx, y: current.y + dy})
	}

	// shoelace formula for area of a polygon given it's vertices (in order)
	return shoelace(vertices_around)
}

func main() {
	data := load_data("input.txt")
	fmt.Println(get_size(data, false))
	fmt.Println(get_size(data, true))
}
