package day9

import (
	"fmt"
	"loader"
	"sort"
	"strconv"
	"strings"
	"utils"
)

type Pos struct {
	x int
	y int
}

type Wall struct {
	tile1 Pos
	tile2 Pos
}

type Floor struct {
	tiles []Pos
	walls []Wall
}

func parse_data(data []string) Floor {
	tiles := make([]Pos, len(data))
	for i, line := range data {
		parts := strings.Split(line, ",")
		x, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		y, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		tiles[i] = Pos{x: x, y: y}
	}

	walls := []Wall{}
	for i := 0; i < len(tiles)-1; i++ {
		walls = append(walls, Wall{tile1: tiles[i], tile2: tiles[i+1]})
	}
	walls = append(walls, Wall{tile1: tiles[len(tiles)-1], tile2: tiles[0]})

	return Floor{
		tiles: tiles,
		walls: walls,
	}
}

func largest_rectangle(floor Floor) int {
	max := 0
	for i := 0; i < len(floor.tiles); i++ {
		for j := i + 1; j < len(floor.tiles); j++ {
			size := (utils.Abs(floor.tiles[i].x-floor.tiles[j].x) + 1) * (utils.Abs(floor.tiles[i].y-floor.tiles[j].y) + 1)
			if size > max {
				max = size
			}
		}
	}
	return max
}

func largest_rectangle_inside(floor Floor) int {
	rectangles := [][]int{}

	for i := 0; i < len(floor.tiles); i++ {
		for j := i + 1; j < len(floor.tiles); j++ {
			size := (utils.Abs(floor.tiles[i].x-floor.tiles[j].x) + 1) * (utils.Abs(floor.tiles[i].y-floor.tiles[j].y) + 1)
			rectangles = append(rectangles, []int{size, i, j})
		}
	}

	sort.Slice(rectangles, func(i, j int) bool {
		return rectangles[i][0] > rectangles[j][0]
	})

outer:
	for _, rectangle := range rectangles {
		tile1 := floor.tiles[rectangle[1]]
		tile2 := floor.tiles[rectangle[2]]

		from_x, to_x, from_y, to_y := tile1.x, tile2.x, tile1.y, tile2.y
		if from_x > to_x {
			from_x, to_x = to_x, from_x
		}
		if from_y > to_y {
			from_y, to_y = to_y, from_y
		}

		for _, wall := range floor.walls {
			x_start, x_end := wall.tile1.x, wall.tile2.x
			if x_start > x_end {
				x_start, x_end = x_end, x_start
			}
			y_start, y_end := wall.tile1.y, wall.tile2.y
			if y_start > y_end {
				y_start, y_end = y_end, y_start
			}

			for y := y_start; y <= y_end; y++ {
				for x := x_start; x <= x_end; x++ {
					if x > from_x && x < to_x && y > from_y && y < to_y {
						continue outer
					}
				}
			}
		}

		return rectangle[0]
	}

	return -1
}

func Run() {
	loader.Day = 9
	data := loader.GetStrings()
	floor := parse_data(data)
	part1 := largest_rectangle(floor)
	part2 := largest_rectangle_inside(floor)

	fmt.Printf("%d %d\n", part1, part2)
}
