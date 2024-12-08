package day8

import (
	"fmt"
	"loader"
)

type Pos struct {
	x int
	y int
}

func antinodes(grid []string) int {
	height := len(grid)
	width := len(grid[0])

	antennas := map[byte][]Pos{}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			space := grid[y][x]
			if space != '.' {
				antennas[space] = append(antennas[space], Pos{x: x, y: y})
			}
		}
	}

	antinodes := map[Pos]struct{}{}

	for _, positions := range antennas {
		for i := 0; i < len(positions); i++ {
			for j := i + 1; j < len(positions); j++ {
				antinode := Pos{
					x: positions[j].x + (positions[j].x - positions[i].x),
					y: positions[j].y + (positions[j].y - positions[i].y),
				}
				if antinode.x >= 0 && antinode.x < width && antinode.y >= 0 && antinode.y < height {
					antinodes[antinode] = struct{}{}
				}
				antinode = Pos{
					x: positions[i].x - (positions[j].x - positions[i].x),
					y: positions[i].y - (positions[j].y - positions[i].y),
				}
				if antinode.x >= 0 && antinode.x < width && antinode.y >= 0 && antinode.y < height {
					antinodes[antinode] = struct{}{}
				}
			}
		}
	}

	return len(antinodes)
}

func Run() {
	loader.Day = 8
	grid := loader.GetStrings()
	part1 := antinodes(grid)

	part2 := -1
	fmt.Printf("%d %d\n", part1, part2)
}
