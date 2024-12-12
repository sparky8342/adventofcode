package day12

import (
	"fmt"
	"loader"
	"sort"
)

var size int

type Pos struct {
	x int
	y int
}

type Edge struct {
	x float64
	y float64
}

func find_region(grid [][]byte, start_pos Pos) (int, int) {
	queue := []Pos{start_pos}
	visited := map[Pos]struct{}{}
	visited[start_pos] = struct{}{}

	plant := grid[start_pos.y][start_pos.x]
	region_size := 0
	perimeter := 0

	vertical_edges := []Edge{}
	horizontal_edges := []Edge{}

	for len(queue) > 0 {
		pos := queue[0]
		queue = queue[1:]

		region_size++

		neighbours := []Pos{}
		// left
		if pos.x == 0 || grid[pos.y][pos.x-1] != plant {
			vertical_edges = append(vertical_edges, Edge{x: float64(pos.x) - 0.1, y: float64(pos.y)})
			perimeter++
		} else {
			next_pos := Pos{x: pos.x - 1, y: pos.y}
			neighbours = append(neighbours, next_pos)
		}

		// right
		if pos.x == size-1 || grid[pos.y][pos.x+1] != plant {
			vertical_edges = append(vertical_edges, Edge{x: float64(pos.x) + 0.1, y: float64(pos.y)})
			perimeter++
		} else {
			next_pos := Pos{x: pos.x + 1, y: pos.y}
			neighbours = append(neighbours, next_pos)
		}

		// up
		if pos.y == 0 || grid[pos.y-1][pos.x] != plant {
			horizontal_edges = append(horizontal_edges, Edge{x: float64(pos.x), y: float64(pos.y) - 0.1})
			perimeter++
		} else {
			next_pos := Pos{x: pos.x, y: pos.y - 1}
			neighbours = append(neighbours, next_pos)
		}

		// down
		if pos.y == size-1 || grid[pos.y+1][pos.x] != plant {
			horizontal_edges = append(horizontal_edges, Edge{x: float64(pos.x), y: float64(pos.y) + 0.1})
			perimeter++
		} else {
			next_pos := Pos{x: pos.x, y: pos.y + 1}
			neighbours = append(neighbours, next_pos)
		}

		for _, neighbour := range neighbours {
			if _, ok := visited[neighbour]; ok {
				continue
			}
			queue = append(queue, neighbour)
			visited[neighbour] = struct{}{}
		}
	}

	for pos := range visited {
		grid[pos.y][pos.x] = '.'
	}

	sort.Slice(vertical_edges, func(i, j int) bool {
		if vertical_edges[i].x == vertical_edges[j].x {
			return vertical_edges[i].y < vertical_edges[j].y
		} else {
			return vertical_edges[i].x < vertical_edges[j].x
		}
	})
	vertical_sides := 1
	for i := 0; i < len(vertical_edges)-1; i++ {
		if vertical_edges[i].x != vertical_edges[i+1].x {
			vertical_sides++
		} else if vertical_edges[i].y+1 != vertical_edges[i+1].y {
			vertical_sides++
		}
	}

	sort.Slice(horizontal_edges, func(i, j int) bool {
		if horizontal_edges[i].y == horizontal_edges[j].y {
			return horizontal_edges[i].x < horizontal_edges[j].x
		} else {
			return horizontal_edges[i].y < horizontal_edges[j].y
		}
	})
	horizontal_sides := 1
	for i := 0; i < len(horizontal_edges)-1; i++ {
		if horizontal_edges[i].y != horizontal_edges[i+1].y {
			horizontal_sides++
		} else if horizontal_edges[i].x+1 != horizontal_edges[i+1].x {
			horizontal_sides++
		}
	}

	return region_size * perimeter, region_size * (vertical_sides + horizontal_sides)
}

func price(data []string) (int, int) {
	grid := make([][]byte, len(data))
	for i, line := range data {
		grid[i] = []byte(line)
	}
	size = len(grid)

	total := 0
	discount_total := 0
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			if grid[y][x] != '.' {
				t, d := find_region(grid, Pos{x: x, y: y})
				total += t
				discount_total += d
			}
		}
	}
	return total, discount_total
}

func Run() {
	loader.Day = 12
	data := loader.GetStrings()

	part1, part2 := price(data)

	fmt.Printf("%d %d\n", part1, part2)
}
