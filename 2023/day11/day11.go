package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Universe struct {
	galaxies      []Galaxy
	height        int
	width         int
	empty_rows    map[int]Empty
	empty_columns map[int]Empty
}

type Galaxy struct {
	x int
	y int
}

type Empty struct{}

func load_data(filename string) []string {
	data, _ := ioutil.ReadFile(filename)
	if data[len(data)-1] == '\n' {
		data = data[:len(data)-1]
	}
	return strings.Split(string(data), "\n")
}

func parse_data(data []string) Universe {
	universe := Universe{
		galaxies:      []Galaxy{},
		height:        len(data),
		width:         len(data[0]),
		empty_rows:    map[int]Empty{},
		empty_columns: map[int]Empty{},
	}

	for y := 0; y < universe.height; y++ {
		for x := 0; x < universe.width; x++ {
			if data[y][x] == '#' {
				universe.galaxies = append(universe.galaxies, Galaxy{x: x, y: y})
			}
		}
	}

	for y := 0; y < universe.height; y++ {
		empty_row := true
		for x := 0; x < universe.width; x++ {
			if data[y][x] == '#' {
				empty_row = false
				break
			}
		}
		if empty_row {
			universe.empty_rows[y] = Empty{}
		}
	}

	for x := 0; x < universe.width; x++ {
		empty_column := true
		for y := 0; y < universe.height; y++ {
			if data[y][x] == '#' {
				empty_column = false
				break
			}
		}
		if empty_column {
			universe.empty_columns[x] = Empty{}
		}
	}

	return universe
}

func (universe *Universe) path_sum(expand int) int {
	sum := 0
	for i := 0; i < len(universe.galaxies); i++ {
		for j := i + 1; j < len(universe.galaxies); j++ {

			add_x := 0
			x_start := universe.galaxies[i].x
			x_end := universe.galaxies[j].x
			if x_start > x_end {
				x_start, x_end = x_end, x_start
			}
			for x := x_start + 1; x < x_end; x++ {
				if _, exists := universe.empty_columns[x]; exists {
					add_x += expand - 1
				}
			}

			add_y := 0
			y_start := universe.galaxies[i].y
			y_end := universe.galaxies[j].y
			if y_start > y_end {
				y_start, y_end = y_end, y_start
			}
			for y := y_start + 1; y < y_end; y++ {
				if _, exists := universe.empty_rows[y]; exists {
					add_y += expand - 1
				}
			}

			sum += x_end - x_start + add_x + y_end - y_start + add_y
		}
	}

	return sum
}

func main() {
	data := load_data("input.txt")
	universe := parse_data(data)
	fmt.Println(universe.path_sum(2))
	fmt.Println(universe.path_sum(1000000))
}
