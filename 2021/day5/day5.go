package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type Line struct {
	x1, y1, x2, y2 int
}

type Pos struct {
	x, y int
}

func get_data() []Line {
	data, _ := ioutil.ReadFile("input.txt")
	data = data[:len(data)-1]
	data_lines := strings.Split(string(data), "\n")

	lines := []Line{}
	r := regexp.MustCompile("([0-9]+),([0-9]+) -> ([0-9]+),([0-9]+)")
	for _, data_line := range data_lines {
		match := r.FindStringSubmatch(data_line)
		x1, _ := strconv.Atoi(match[1])
		y1, _ := strconv.Atoi(match[2])
		x2, _ := strconv.Atoi(match[3])
		y2, _ := strconv.Atoi(match[4])
		if y1 == y2 && x1 > x2 {
			x1, x2 = x2, x1
		}
		if x1 == x2 && y1 > y2 {
			y1, y2 = y2, y1
		}
		lines = append(lines, Line{x1: x1, y1: y1, x2: x2, y2: y2})
	}

	return lines
}

func update_visited(visited map[Pos]int, pos *Pos) {
	if _, ok := visited[*pos]; ok {
		visited[*pos]++
	} else {
		visited[*pos] = 1
	}
}

func get_overlaps(lines []Line, use_diagonals bool) int {
	visited := make(map[Pos]int)

	for _, line := range lines {
		if line.x1 == line.x2 {
			for y := line.y1; y <= line.y2; y++ {
				pos := Pos{x: line.x1, y: y}
				update_visited(visited, &pos)
			}
		} else if line.y1 == line.y2 {
			for x := line.x1; x <= line.x2; x++ {
				pos := Pos{x: x, y: line.y1}
				update_visited(visited, &pos)
			}
		} else if use_diagonals {
			var x_inc, y_inc int
			if line.x1 < line.x2 {
				x_inc = 1
			} else {
				x_inc = -1
			}
			if line.y1 < line.y2 {
				y_inc = 1
			} else {
				y_inc = -1
			}

			x := line.x1
			y := line.y1

			pos := Pos{x: x, y: y}
			update_visited(visited, &pos)
			for x != line.x2 {
				x += x_inc
				y += y_inc
				pos := Pos{x: x, y: y}
				update_visited(visited, &pos)
			}
		}
	}

	overlaps := 0
	for _, v := range visited {
		if v > 1 {
			overlaps++
		}
	}
	return overlaps
}

func main() {
	lines := get_data()
	overlaps := get_overlaps(lines, false)
	fmt.Println(overlaps)
	overlaps = get_overlaps(lines, true)
	fmt.Println(overlaps)
}
