package day12

import (
	"fmt"
	"loader"
	"strconv"
	"strings"
)

type Shape struct {
	id         int
	variations [][]string
	squares    int
}

type Region struct {
	width     int
	height    int
	shape_ids []int
}

func parse_data(data []string) ([]Shape, []Region) {
	shapes := []Shape{}
	regions := []Region{}

	var shape_id int
	var shape []string

	for _, line := range data {
		if line == "" {
			variations := [][]string{shape}
			for i := 0; i < 3; i++ {
				shape = rotate_right(shape)
				variations = append(variations, shape)
			}
			shape = rotate_right(shape)
			shape = flip(shape)
			variations = append(variations, shape)
			for i := 0; i < 3; i++ {
				shape = rotate_right(shape)
				variations = append(variations, shape)
			}

			squares := 0
			for y := 0; y < 3; y++ {
				for x := 0; x < 3; x++ {
					if shape[y][x] == '#' {
						squares++
					}
				}
			}

			shapes = append(shapes, Shape{
				id:         shape_id,
				variations: variations,
				squares:    squares,
			})
			shape = []string{}
		} else if line[1] == ':' {
			shape_id = int(line[0] - '0')
		} else if len(line) == 3 {
			shape = append(shape, line)
		} else {
			line = strings.Replace(line, ":", "", 1)
			parts := strings.Split(line, " ")
			dimensions := strings.Split(parts[0], "x")
			width, err := strconv.Atoi(dimensions[0])
			if err != nil {
				panic(err)
			}
			height, err := strconv.Atoi(dimensions[1])
			if err != nil {
				panic(err)
			}
			shape_ids := []int{}
			for i := 1; i < len(parts); i++ {
				n, err := strconv.Atoi(parts[i])
				if err != nil {
					panic(err)
				}
				shape_ids = append(shape_ids, n)
			}
			regions = append(regions, Region{
				width:     width,
				height:    height,
				shape_ids: shape_ids,
			})
		}
	}

	return shapes, regions
}

func rotate_right(shape []string) []string {
	r := make([][]byte, 3)
	for i := range r {
		r[i] = make([]byte, 3)
	}

	for y := 0; y < 3; y++ {
		for x := 0; x < 3; x++ {
			r[y][x] = shape[3-(x+1)][y]
		}
	}

	new_shape := make([]string, 3)
	for i := range r {
		new_shape[i] = string(r[i])
	}

	return new_shape
}

func flip(shape []string) []string {
	f := make([][]byte, 3)
	for i := range f {
		f[i] = make([]byte, 3)
	}

	for y := 0; y < 3; y++ {
		for x := 0; x < 3; x++ {
			f[2-y][x] = shape[y][x]
		}
	}

	new_shape := make([]string, 3)
	for i := range f {
		new_shape[i] = string(f[i])
	}

	return new_shape
}

func shape_fits(grid [][]byte, shape []string, x int, y int) bool {
	height := len(grid)
	width := len(grid[0])
	if height-y < 3 || width-x < 3 {
		return false
	}

	for shape_y := 0; shape_y < 3; shape_y++ {
		for shape_x := 0; shape_x < 3; shape_x++ {
			if shape[shape_y][shape_x] == '#' {
				if grid[shape_y+y][shape_x+x] == '#' {
					return false
				}
			}
		}
	}
	return true
}

func place_shape(grid [][]byte, shape []string, x int, y int) [][]byte {
	for shape_y := 0; shape_y < 3; shape_y++ {
		for shape_x := 0; shape_x < 3; shape_x++ {
			if shape[shape_y][shape_x] == '#' {
				grid[shape_y+y][shape_x+x] = '#'
			}
		}
	}
	return grid
}

func remove_shape(grid [][]byte, shape []string, x int, y int) [][]byte {
	for shape_y := 0; shape_y < 3; shape_y++ {
		for shape_x := 0; shape_x < 3; shape_x++ {
			if shape[shape_y][shape_x] == '#' {
				grid[shape_y+y][shape_x+x] = ' '
			}
		}
	}
	return grid
}

func search(shapes []Shape, region Region, grid [][]byte, shape_i int, shape_count int) bool {
	if shape_count == region.shape_ids[shape_i] {
		shape_i++
		if shape_i == len(region.shape_ids) {
			return true
		}
		shape_count = 0
	}

	for region.shape_ids[shape_i] == 0 {
		shape_i++
		if shape_i == len(region.shape_ids) {
			return true
		}
	}

	shape := shapes[shape_i]

	for _, variation := range shape.variations {
		for y := 0; y < region.height; y++ {
			for x := 0; x < region.width; x++ {
				if shape_fits(grid, variation, x, y) {
					grid = place_shape(grid, variation, x, y)
					if search(shapes, region, grid, shape_i, shape_count+1) {
						return true
					}
					grid = remove_shape(grid, variation, x, y)
				}
			}
		}
	}

	return false
}

func place_shapes(shapes []Shape, regions []Region) int {
	total := 0

	for _, region := range regions {
		region_size := region.height * region.width
		spaces_needed := 0
		for i := 0; i < len(region.shape_ids); i++ {
			spaces_needed += region.shape_ids[i] * shapes[i].squares
		}
		if spaces_needed > region_size {
			continue
		}

		grid := make([][]byte, region.height)
		for i := range grid {
			grid[i] = make([]byte, region.width)
			for j := 0; j < region.width; j++ {
				grid[i][j] = ' '
			}
		}

		if search(shapes, region, grid, 0, 0) {
			total++
		}
	}

	return total
}

func Run() {
	loader.Day = 12
	data := loader.GetStrings()
	shapes, regions := parse_data(data)
	answer := place_shapes(shapes, regions)

	fmt.Printf("%d\n", answer)
}
