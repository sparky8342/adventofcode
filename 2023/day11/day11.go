package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Universe struct {
	grid   [][]byte
	height int
	width  int
}

type Galaxy struct {
	x int
	y int
}

func abs(n int) int {
	if n < 0 {
		return n * -1
	} else {
		return n
	}
}

func load_data(filename string) []string {
	data, _ := ioutil.ReadFile(filename)
	if data[len(data)-1] == '\n' {
		data = data[:len(data)-1]
	}
	return strings.Split(string(data), "\n")
}

func parse_data(data []string) Universe {
	universe := Universe{
		grid:   [][]byte{},
		height: len(data),
		width:  len(data[0]),
	}

	for _, line := range data {
		universe.grid = append(universe.grid, []byte(line))
	}

	return universe
}

func (universe *Universe) expand() {
	// y direction
	for y := universe.height - 1; y >= 0; y-- {
		empty := true
		for x := 0; x < universe.width; x++ {
			if universe.grid[y][x] == '#' {
				empty = false
				break
			}
		}
		if empty {
			row := make([]byte, universe.width)
			for i := 0; i < universe.width; i++ {
				row[i] = '.'
			}
			universe.grid = append(universe.grid[0:y], append([][]byte{row}, universe.grid[y:]...)...)
		}
	}
	universe.height = len(universe.grid)

	// x direction
	for x := universe.width - 1; x >= 0; x-- {
		empty := true
		for y := 0; y < universe.height; y++ {
			if universe.grid[y][x] == '#' {
				empty = false
				break
			}
		}
		if empty {

			for y := 0; y < universe.height; y++ {
				universe.grid[y] = append(universe.grid[y][0:x], append([]byte{'.'}, universe.grid[y][x:]...)...)
			}
		}
	}
	universe.width = len(universe.grid[0])
}

func (universe *Universe) path_sum() int {
	galaxies := []Galaxy{}

	for y := 0; y < universe.height; y++ {
		for x := 0; x < universe.width; x++ {
			if universe.grid[y][x] == '#' {
				galaxies = append(galaxies, Galaxy{x: x, y: y})
			}
		}
	}

	sum := 0
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			sum += abs(galaxies[i].x-galaxies[j].x) + abs(galaxies[i].y-galaxies[j].y)
		}
	}

	return sum
}

func (universe *Universe) pp() {
	for y := 0; y < universe.height; y++ {
		fmt.Println(string(universe.grid[y]))
	}
	fmt.Println()
}

func main() {
	data := load_data("input.txt")
	universe := parse_data(data)
	universe.expand()
	fmt.Println(universe.path_sum())
}
