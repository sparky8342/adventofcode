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

type Cubes map[Cube]Empty

func load_data(filename string) []string {
	data, _ := ioutil.ReadFile(filename)
	if data[len(data)-1] == '\n' {
		data = data[:len(data)-1]
	}
	return strings.Split(string(data), "\n")
}

func parse_data(data []string) Cubes {
	cubes := Cubes{}
	for _, line := range data {
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		z, _ := strconv.Atoi(parts[2])
		cubes[Cube{x: x, y: y, z: z}] = Empty{}
	}
	return cubes
}

func count_sides(cubes Cubes) int {
	sides := len(cubes) * 6

	for cube := range cubes {
		possible := []Cube{
			Cube{x: cube.x - 1, y: cube.y, z: cube.z},
			Cube{x: cube.x + 1, y: cube.y, z: cube.z},
			Cube{x: cube.x, y: cube.y - 1, z: cube.z},
			Cube{x: cube.x, y: cube.y + 1, z: cube.z},
			Cube{x: cube.x, y: cube.y, z: cube.z - 1},
			Cube{x: cube.x, y: cube.y, z: cube.z + 1},
		}

		for _, c := range possible {
			if _, exists := cubes[c]; exists {
				sides--
			}
		}
	}

	return sides
}

func main() {
	data := load_data("input.txt")
	cubes := parse_data(data)
	sides := count_sides(cubes)
	fmt.Println(sides)
}
