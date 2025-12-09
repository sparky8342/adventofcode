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
	x1, y1 int
	x2, y2 int
}

type Floor struct {
	tiles []Pos
	walls []Wall
}

func order_pair(a, b int) (int, int) {
	if a > b {
		return b, a
	} else {
		return a, b
	}
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
	for i := 0; i < len(tiles); i++ {
		wall := Wall{
			x1: tiles[i].x,
			y1: tiles[i].y,
			x2: tiles[(i+1)%len(tiles)].x,
			y2: tiles[(i+1)%len(tiles)].y,
		}
		wall.x1, wall.x2 = order_pair(wall.x1, wall.x2)
		wall.y1, wall.y2 = order_pair(wall.y1, wall.y2)
		walls = append(walls, wall)
	}

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
		from_x, to_x = order_pair(from_x, to_x)
		from_y, to_y = order_pair(from_y, to_y)

		for _, wall := range floor.walls {
			for y := wall.y1; y <= wall.y2; y++ {
				for x := wall.x1; x <= wall.x2; x++ {
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
