package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type Step struct {
	on                     bool
	x1, x2, y1, y2, z1, z2 int
}

type Cube struct {
	x, y, z int
}

func get_data() []Step {
	data, _ := ioutil.ReadFile("input.txt")
	data = data[:len(data)-1]
	lines := strings.Split(string(data), "\n")

	r := regexp.MustCompile("(on|off).*?([\\-0-9]+).*?([\\-0-9]+).*?([\\-0-9]+).*?([\\-0-9]+).*?([\\-0-9]+).*?([\\-0-9]+)")
	steps := []Step{}
	for _, line := range lines {
		match := r.FindStringSubmatch(line)
		on := true
		if match[1] == "off" {
			on = false
		}
		x1, _ := strconv.Atoi(match[2])
		x2, _ := strconv.Atoi(match[3])
		y1, _ := strconv.Atoi(match[4])
		y2, _ := strconv.Atoi(match[5])
		z1, _ := strconv.Atoi(match[6])
		z2, _ := strconv.Atoi(match[7])
		steps = append(steps, Step{on: on, x1: x1, x2: x2, y1: y1, y2: y2, z1: z1, z2: z2})
	}
	return steps
}

func main() {
	steps := get_data()

	cubes := map[Cube]struct{}{}

	for _, step := range steps {
		if step.x1 > 50 || step.x2 < -50 || step.y1 > 50 || step.y2 < -50 || step.z1 > 50 || step.z2 < -50 {
			continue
		}

		for x := step.x1; x <= step.x2; x++ {
			for y := step.y1; y <= step.y2; y++ {
				for z := step.z1; z <= step.z2; z++ {
					if x >= -50 && x <= 50 && y >= -50 && y <= 50 && z >= -50 && z <= 50 {
						cube := Cube{x: x, y: y, z: z}
						if step.on {
							cubes[cube] = struct{}{}
						} else {
							if _, found := cubes[cube]; found {
								delete(cubes, cube)
							}
						}
					}
				}
			}
		}
	}

	fmt.Println(len(cubes))
}
