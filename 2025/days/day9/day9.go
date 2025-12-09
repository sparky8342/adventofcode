package day9

import (
	"fmt"
	"loader"
	"math"
	"strconv"
	"strings"
	"utils"
)

type Pos struct {
	x int
	y int
}

func parse_data(data []string) []Pos {
	tiles := make([]Pos, len(data))
	for i, line := range data {
		parts := strings.Split(line, ",")
		n1, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		n2, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		tiles[i] = Pos{x: n1, y: n2}
	}
	return tiles
}

func largest_rectangle_attempt1(tiles []Pos) int {
	tl := Pos{x: math.MaxInt64, y: math.MaxInt64}
	tr := Pos{x: math.MinInt64, y: math.MaxInt64}
	bl := Pos{x: math.MaxInt64, y: math.MinInt64}
	br := Pos{x: math.MinInt64, y: math.MinInt64}

	for _, tile := range tiles {
		if tile.x <= tl.x && tile.y <= tl.y {
			tl = tile
		}
		if tile.x >= tr.x && tile.y <= tr.y {
			tr = tile
		}
		if tile.x <= bl.x && tile.y >= bl.y {
			bl = tile
		}
		if tile.x >= br.x && tile.y >= br.y {
			br = tile
		}
	}

	fmt.Println(tl, tr, bl, br)

	rect1 := (br.x - tl.x + 1) * (br.y - tl.y + 1)
	rect2 := (tr.x - bl.x + 1) * (bl.y - tr.y + 1)

	fmt.Println(rect1, rect2)

	if rect1 > rect2 {
		return rect1
	} else {
		return rect2
	}
}

func largest_rectangle(tiles []Pos) int {
	max := 0
	for i := 0; i < len(tiles); i++ {
		for j := i + 1; j < len(tiles); j++ {
			size := (utils.Abs(tiles[i].x-tiles[j].x) + 1) * (utils.Abs(tiles[i].y-tiles[j].y) + 1)
			if size > max {
				max = size
			}
		}
	}
	return max
}

func Run() {
	loader.Day = 9
	data := loader.GetStrings()
	tiles := parse_data(data)
	part1 := largest_rectangle(tiles)

	fmt.Printf("%d %d\n", part1, 0)
}
