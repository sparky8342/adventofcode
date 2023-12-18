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

func load_data(filename string) []string {
	data, _ := ioutil.ReadFile(filename)
	if data[len(data)-1] == '\n' {
		data = data[:len(data)-1]
	}
	return strings.Split(string(data), "\n")
}

func get_size(data []string) int {
	x, y := 0, 0

	holes := map[Pos]struct{}{}
	holes[Pos{x: 0, y: 0}] = struct{}{}

	x_min := 0
	x_max := 0
	y_min := 0
	y_max := 0

	for _, line := range data {
		parts := strings.Split(line, " ")
		n, _ := strconv.Atoi(parts[1])
		switch parts[0][0] {
		case 'U':
			next_y := y - n
			for ; y > next_y; y-- {
				holes[Pos{x: x, y: y}] = struct{}{}
			}
		case 'R':
			next_x := x + n
			for ; x < next_x; x++ {
				holes[Pos{x: x, y: y}] = struct{}{}
			}
		case 'D':
			next_y := y + n
			for ; y < next_y; y++ {
				holes[Pos{x: x, y: y}] = struct{}{}
			}
		case 'L':
			next_x := x - n
			for ; x > next_x; x-- {
				holes[Pos{x: x, y: y}] = struct{}{}
			}
		}

		if x < x_min {
			x_min = x
		} else if x > x_max {
			x_max = x
		}
		if y < y_min {
			y_min = y
		} else if y > y_max {
			y_max = y
		}
	}

	// fill
	start_y := y_min + 1
	start_x := 0
	for x := x_min; x <= x_max; x++ {
		if _, exists := holes[Pos{x: x, y: start_y}]; exists {
			start_x = x + 1
			break
		}
	}

	queue := []Pos{Pos{x: start_x, y: start_y}}

	for len(queue) > 0 {
		pos := queue[0]
		queue = queue[1:]

		if _, exists := holes[pos]; exists {
			continue
		}

		holes[pos] = struct{}{}

		queue = append(queue, Pos{x: pos.x + 1, y: pos.y})
		queue = append(queue, Pos{x: pos.x - 1, y: pos.y})
		queue = append(queue, Pos{x: pos.x, y: pos.y + 1})
		queue = append(queue, Pos{x: pos.x, y: pos.y - 1})
	}

	/*
		for y := y_min; y <= y_max; y++ {
			for x := x_min; x <= x_max; x++ {
				if _, exists := holes[Pos{x : x, y : y}]; exists {
					fmt.Print("#")
				} else {
					fmt.Print(" ")
				}
			}
			fmt.Println()
		}
		fmt.Println()
	*/

	return len(holes)
}

func main() {
	data := load_data("input.txt")
	fmt.Println(get_size(data))
}
