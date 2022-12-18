package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Cube struct {
	x int
	y int
	z int
}

type Empty struct {
}

type Droplet struct {
	cubes        map[Cube]Empty
	x_min, x_max int
	y_min, y_max int
	z_min, z_max int
}

func load_data(filename string) []string {
	data, _ := ioutil.ReadFile(filename)
	if data[len(data)-1] == '\n' {
		data = data[:len(data)-1]
	}
	return strings.Split(string(data), "\n")
}

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func parse_data(data []string) Droplet {
	droplet := Droplet{
		cubes: map[Cube]Empty{},
		x_min: 99,
		x_max: -99,
		y_min: 99,
		y_max: -99,
		z_min: 99,
		z_max: -99,
	}

	for _, line := range data {
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		z, _ := strconv.Atoi(parts[2])
		droplet.cubes[Cube{x: x, y: y, z: z}] = Empty{}
		droplet.x_min = min(droplet.x_min, x)
		droplet.x_max = max(droplet.x_max, x)
		droplet.y_min = min(droplet.y_min, y)
		droplet.y_max = max(droplet.y_max, y)
		droplet.z_min = min(droplet.z_min, z)
		droplet.z_max = max(droplet.z_max, z)
	}

	return droplet
}

func count_sides(droplet Droplet) int {
	sides := len(droplet.cubes) * 6

	for cube := range droplet.cubes {
		possible := []Cube{
			Cube{x: cube.x - 1, y: cube.y, z: cube.z},
			Cube{x: cube.x + 1, y: cube.y, z: cube.z},
			Cube{x: cube.x, y: cube.y - 1, z: cube.z},
			Cube{x: cube.x, y: cube.y + 1, z: cube.z},
			Cube{x: cube.x, y: cube.y, z: cube.z - 1},
			Cube{x: cube.x, y: cube.y, z: cube.z + 1},
		}

		for _, c := range possible {
			if _, exists := droplet.cubes[c]; exists {
				sides--
			}
		}
	}

	return sides
}

func flood_fill(droplet Droplet, space Cube) (map[Cube]Empty, int) {
	queue := []Cube{space}
	spaces := map[Cube]Empty{}
	spaces[space] = Empty{}

	sides := 0

	went_out := false
	for len(queue) > 0 {
		sp := queue[0]
		queue = queue[1:]

		if sp.x <= droplet.x_min || sp.x >= droplet.x_max || sp.y <= droplet.y_min || sp.y >= droplet.y_max || sp.z <= droplet.z_min || sp.z >= droplet.z_max {
			went_out = true
			break
		}

		surrounding := []Cube{
			Cube{x: sp.x - 1, y: sp.y, z: sp.z},
			Cube{x: sp.x + 1, y: sp.y, z: sp.z},
			Cube{x: sp.x, y: sp.y - 1, z: sp.z},
			Cube{x: sp.x, y: sp.y + 1, z: sp.z},
			Cube{x: sp.x, y: sp.y, z: sp.z - 1},
			Cube{x: sp.x, y: sp.y, z: sp.z + 1},
		}

		for _, c := range surrounding {
			if _, exists := droplet.cubes[c]; !exists {
				if _, seen := spaces[c]; !seen {
					spaces[c] = Empty{}
					queue = append(queue, c)
				}
			} else {
				sides++
			}
		}
	}
	if went_out {
		return map[Cube]Empty{}, 0
	} else {
		return spaces, sides
	}
}

func count_holes(droplet Droplet) int {
	holes := map[Cube]Empty{}

	total_sides := 0
	for x := droplet.x_min; x <= droplet.x_max; x++ {
		for y := droplet.y_min; y <= droplet.y_max; y++ {
			for z := droplet.z_min; z <= droplet.z_max; z++ {
				space := Cube{x: x, y: y, z: z}
				if _, is_cube := droplet.cubes[space]; is_cube {
					continue
				}

				if _, hole := holes[space]; hole {
					continue
				}

				spaces, sides := flood_fill(droplet, space)
				for space := range spaces {
					holes[space] = Empty{}
				}
				total_sides += sides
			}
		}
	}

	return total_sides
}

func main() {
	data := load_data("input.txt")
	droplet := parse_data(data)
	sides := count_sides(droplet)
	fmt.Println(sides)

	inside_sides := count_holes(droplet)
	fmt.Println(sides - inside_sides)
}
