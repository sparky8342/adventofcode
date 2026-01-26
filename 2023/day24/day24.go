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

func find_relative_hailstones(hailstones []Hailstone) []Hailstone {
	// find 3 with same vy
	seen := map[float64][]Hailstone{}
	for _, hailstone := range hailstones {
		seen[hailstone.vy] = append(seen[hailstone.vy], hailstone)
	}
	for _, hailstones := range seen {
		if len(hailstones) >= 3 {
			return hailstones[0:3]
		}
	}
	return []Hailstone{}
}

func find_rock(hailstones []Hailstone) int {
	hail := find_relative_hailstones(hailstones)

	// calculate relative velocities of hail 1 and 2 to hail 0
	// the y component is zero due to selection of hail
	vxr1 := hail[1].vx - hail[0].vx
	vzr1 := hail[1].vz - hail[0].vz
	vxr2 := hail[2].vx - hail[0].vx
	vzr2 := hail[2].vz - hail[0].vz

	// relative initial position of hail 1
	xr1 := hail[1].x - hail[0].x
	yr1 := hail[1].y - hail[0].y
	zr1 := hail[1].z - hail[0].z

	// relative initial position of hail 2
	xr2 := hail[2].x - hail[0].x
	yr2 := hail[2].y - hail[0].y
	zr2 := hail[2].z - hail[0].z

	// Solve set of two linear equations x=x and z=z
	num := (yr2 * xr1 * vzr1) - (vxr1 * yr2 * zr1) + (yr1 * zr2 * vxr1) - (yr1 * xr2 * vzr1)
	den := yr1 * ((vzr1 * vxr2) - (vxr1 * vzr2))
	t2 := num / den

	// Substitute t2 into a t1 equation
	num = (yr1 * xr2) + (yr1 * vxr2 * t2) - (yr2 * xr1)
	den = yr2 * vxr1
	t1 := num / den

	// calculate collision position at t1 and t2 of hail 1 and 2 in normal frame of reference
	cx1 := hail[1].x + (t1 * hail[1].vx)
	cy1 := hail[1].y + (t1 * hail[1].vy)
	cz1 := hail[1].z + (t1 * hail[1].vz)

	cx2 := hail[2].x + (t2 * hail[2].vx)
	cy2 := hail[2].y + (t2 * hail[2].vy)
	cz2 := hail[2].z + (t2 * hail[2].vz)

	// calculate the vector the rock travelled between those two collisions
	xm := (cx2 - cx1) / (t2 - t1)
	ym := (cy2 - cy1) / (t2 - t1)
	zm := (cz2 - cz1) / (t2 - t1)

	// calculate the initial position of the rock based on its vector
	xc := cx1 - (xm * t1)
	yc := cy1 - (ym * t1)
	zc := cz1 - (zm * t1)

	// answer is sometimes wrong, probably because of floating point errors
	// TODO use big.Float variables instead
	return int(xc + yc + zc)
}

func main() {
	data := load_data("input.txt")
	hailstones := parse_data(data)
	fmt.Println(find_intersections(hailstones, 200000000000000, 400000000000000))
	fmt.Println(find_rock(hailstones))
}
