package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type Hailstone struct {
	x, y, z    float64
	vx, vy, vz float64
	x2, y2, z2 float64
}

type Point struct {
	x float64
	y float64
}

func load_data(filename string) []string {
	data, _ := ioutil.ReadFile(filename)
	if data[len(data)-1] == '\n' {
		data = data[:len(data)-1]
	}
	return strings.Split(string(data), "\n")
}

func parse_data(data []string) []Hailstone {
	hailstones := []Hailstone{}

	re := regexp.MustCompile(",\\s+")

	for _, line := range data {
		parts := strings.Split(line, " @ ")
		coord_parts := strings.Split(parts[0], ", ")
		x, _ := strconv.Atoi(coord_parts[0])
		y, _ := strconv.Atoi(coord_parts[1])
		z, _ := strconv.Atoi(coord_parts[2])

		vel_parts := re.Split(parts[1], -1)
		vx, _ := strconv.Atoi(vel_parts[0])
		vy, _ := strconv.Atoi(vel_parts[1])
		vz, _ := strconv.Atoi(vel_parts[2])

		x2 := x + vx
		y2 := y + vy
		z2 := z + vz

		hailstones = append(hailstones, Hailstone{
			x: float64(x), y: float64(y), z: float64(z),
			vx: float64(vx), vy: float64(vy), vz: float64(vz),
			x2: float64(x2), y2: float64(y2), z2: float64(z2),
		})
	}

	return hailstones
}

func find_intersections(hailstones []Hailstone, min float64, max float64) int {
	intersections := 0
	for i := 0; i < len(hailstones)-1; i++ {
		for j := i + 1; j < len(hailstones); j++ {
			x, y := intersection(
				Point{x: hailstones[i].x, y: hailstones[i].y},
				Point{x: hailstones[i].x2, y: hailstones[i].y2},
				Point{x: hailstones[j].x, y: hailstones[j].y},
				Point{x: hailstones[j].x2, y: hailstones[j].y2},
				min,
				max,
			)
			if x == math.MaxInt32 && y == math.MaxInt32 {
				continue
			}
			if (x < hailstones[i].x && hailstones[i].vx > 0) || (x > hailstones[i].x && hailstones[i].vx < 0) || (y < hailstones[i].y && hailstones[i].vy > 0) || (y > hailstones[i].y && hailstones[i].vy < 0) || (x < hailstones[j].x && hailstones[j].vx > 0) || (x > hailstones[j].x && hailstones[j].vx < 0) || (y < hailstones[j].y && hailstones[j].vy > 0) || (y > hailstones[j].y && hailstones[j].vy < 0) {
				continue
			}
			intersections++
		}
	}
	return intersections
}

func intersection(a Point, b Point, c Point, d Point, min float64, max float64) (float64, float64) {
	a1 := b.y - a.y
	b1 := a.x - b.x
	c1 := a1*a.x + b1*a.y

	a2 := d.y - c.y
	b2 := c.x - d.x
	c2 := a2*c.x + b2*c.y

	determinant := a1*b2 - a2*b1

	if determinant == 0 {
		// parallel
		return math.MaxInt32, math.MaxInt32
	}

	var x, y float64
	x = (b2*c1 - b1*c2) / determinant
	y = (a1*c2 - a2*c1) / determinant
	if x < min || x > max || y < min || y > max {
		return math.MaxInt32, math.MaxInt32
	}
	return x, y
}

func main() {
	data := load_data("input.txt")
	hailstones := parse_data(data)
	fmt.Println(find_intersections(hailstones, 200000000000000, 400000000000000))
}
