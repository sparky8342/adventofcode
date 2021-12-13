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

type Dots map[Pos]struct{}

type Fold struct {
	axis string
	val  int
}

func get_data() (Dots, []Fold) {
	data, _ := ioutil.ReadFile("input.txt")
	data = data[:len(data)-1]
	sections := strings.Split(string(data), "\n\n")

	dots := Dots{}
	for _, line := range strings.Split(sections[0], "\n") {
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		dots[Pos{x: x, y: y}] = struct{}{}
	}

	folds := []Fold{}
	for _, line := range strings.Split(sections[1], "\n") {
		parts := strings.Split(line, " ")
		parts2 := strings.Split(parts[2], "=")
		val, _ := strconv.Atoi(parts2[1])
		folds = append(folds, Fold{axis: parts2[0], val: val})
	}

	return dots, folds
}

func fold_dots(dots Dots, fold Fold) Dots {
	new_dots := Dots{}
	for dot, _ := range dots {
		if fold.axis == "y" && dot.y > fold.val {
			new_dots[Pos{x: dot.x, y: dot.y - (dot.y-fold.val)*2}] = struct{}{}
		} else if fold.axis == "x" && dot.x > fold.val {
			new_dots[Pos{x: dot.x - (dot.x-fold.val)*2, y: dot.y}] = struct{}{}
		} else {
			new_dots[dot] = struct{}{}
		}
	}
	return new_dots
}

func print_dots(dots Dots) {
	height := 0
	width := 0
	for dot, _ := range dots {
		if dot.y > height {
			height = dot.y
		}
		if dot.x > width {
			width = dot.x
		}
	}

	for y := 0; y <= height; y++ {
		for x := 0; x <= width; x++ {
			if _, found := dots[Pos{x: x, y: y}]; found {
				fmt.Print("\u2B1C")
			} else {
				fmt.Print("\u2B1B")
			}
		}
		fmt.Println()
	}
}

func main() {
	dots, folds := get_data()

	for i, fold := range folds {
		dots = fold_dots(dots, fold)
		if i == 0 {
			fmt.Println(len(dots))
		}
	}

	print_dots(dots)
}
