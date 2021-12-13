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

type Grid struct {
	dots map[Pos]bool
}

type Fold struct {
	axis string
	val  int
}

func get_data() (Grid, []Fold) {
	data, _ := ioutil.ReadFile("input.txt")
	data = data[:len(data)-1]
	sections := strings.Split(string(data), "\n\n")

	grid := Grid{dots: map[Pos]bool{}}
	for _, line := range strings.Split(sections[0], "\n") {
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		grid.dots[Pos{x: x, y: y}] = true
	}

	folds := []Fold{}
	for _, line := range strings.Split(sections[1], "\n") {
		parts := strings.Split(line, " ")
		parts2 := strings.Split(parts[2], "=")
		val, _ := strconv.Atoi(parts2[1])
		folds = append(folds, Fold{axis: parts2[0], val: val})
	}

	return grid, folds
}

func (grid *Grid) fold_grid(fold Fold) {
	new_dots := map[Pos]bool{}
	for dot, _ := range grid.dots {
		if fold.axis == "y" && dot.y > fold.val {
			new_dots[Pos{x: dot.x, y: dot.y - (dot.y-fold.val)*2}] = true
		} else if fold.axis == "x" && dot.x > fold.val {
			new_dots[Pos{x: dot.x - (dot.x-fold.val)*2, y: dot.y}] = true
		} else {
			new_dots[dot] = true
		}
	}
	grid.dots = new_dots
}

func (grid *Grid) print_grid() {
	height := 0
	width := 0
	for dot, _ := range grid.dots {
		if dot.y > height {
			height = dot.y
		}
		if dot.x > width {
			width = dot.x
		}
	}

	for y := 0; y <= height; y++ {
		for x := 0; x <= width; x++ {
			if _, found := grid.dots[Pos{x: x, y: y}]; found {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func main() {
	grid, folds := get_data()

	for i, fold := range folds {
		grid.fold_grid(fold)
		if i == 0 {
			fmt.Println(len(grid.dots))
		}
	}

	grid.print_grid()
}
