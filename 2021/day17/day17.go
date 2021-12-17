package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type TargetArea struct {
	x1, x2, y1, y2 int
}

type Probe struct {
	x, y, x_vel, y_vel int
}

func get_data() TargetArea {
	data, _ := ioutil.ReadFile("input.txt")
	line := strings.TrimSuffix(string(data), "\n")

	r := regexp.MustCompile("([0-9]+).*?([0-9]+).*?([\\-0-9]+).*?([\\-0-9]+)")
	match := r.FindStringSubmatch(line)
	x1, _ := strconv.Atoi(match[1])
	x2, _ := strconv.Atoi(match[2])
	y1, _ := strconv.Atoi(match[3])
	y2, _ := strconv.Atoi(match[4])

	return TargetArea{x1: x1, x2: x2, y1: y1, y2: y2}
}

func (probe *Probe) follow_path(target_area TargetArea) (bool, int) {
	hit := false
	max_y := 0
	for probe.y >= target_area.y1 {
		probe.x += probe.x_vel
		probe.y += probe.y_vel

		if probe.y > max_y {
			max_y = probe.y
		}

		if probe.x_vel > 0 {
			probe.x_vel--
		}
		probe.y_vel--

		if probe.x >= target_area.x1 && probe.x <= target_area.x2 && probe.y >= target_area.y1 && probe.y <= target_area.y2 {
			hit = true
			break
		}
	}
	return hit, max_y
}

func main() {
	target_area := get_data()

	best_y := 0
	hits := 0
	// just brute force lots of combinations
	// there should be a better way
	for x_vel := 0; x_vel <= 1000; x_vel++ {
		for y_vel := -1000; y_vel <= 1000; y_vel++ {
			probe := Probe{x: 0, y: 0, x_vel: x_vel, y_vel: y_vel}
			hit, max_y := probe.follow_path(target_area)
			if hit {
				hits++
				if max_y > best_y {
					best_y = max_y
				}
			}
		}
	}

	fmt.Println(best_y)
	fmt.Println(hits)
}
