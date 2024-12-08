package day8

import (
	"fmt"
	"loader"
)

type Pos struct {
	x int
	y int
}

func antinodes(grid []string, spread bool) int {
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
				if spread {
					antinodes[positions[i]] = struct{}{}
					antinodes[positions[j]] = struct{}{}
				}

				dx := positions[j].x - positions[i].x
				dy := positions[j].y - positions[i].y

				antinode := Pos{x: positions[j].x + dx, y: positions[j].y + dy}
				for antinode.x >= 0 && antinode.x < width && antinode.y >= 0 && antinode.y < height {
					antinodes[antinode] = struct{}{}
					antinode.x += dx
					antinode.y += dy
					if !spread {
						break
					}
				}

				antinode = Pos{x: positions[i].x - dx, y: positions[i].y - dy}
				for antinode.x >= 0 && antinode.x < width && antinode.y >= 0 && antinode.y < height {
					antinodes[antinode] = struct{}{}
					antinode.x -= dx
					antinode.y -= dy
					if !spread {
						break
					}
				}
			}
		}
	}

	return len(antinodes)
}

func Run() {
	loader.Day = 8
	grid := loader.GetStrings()

	part1 := antinodes(grid, false)
	part2 := antinodes(grid, true)

	fmt.Printf("%d %d\n", part1, part2)
}
