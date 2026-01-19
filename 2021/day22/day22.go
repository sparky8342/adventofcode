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

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
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

func intersect(a Step, b Step) (bool, Step) {
	if (b.x1 >= a.x1 && b.x1 <= a.x2) || (b.x2 >= a.x1 && b.x2 <= a.x2) || (a.x1 >= b.x1 && a.x1 <= b.x2) || (a.x2 >= b.x1 && a.x2 <= b.x2) {
		if (b.y1 >= a.y1 && b.y1 <= a.y2) || (b.y2 >= a.y1 && b.y2 <= a.y2) || (a.y1 >= b.y1 && a.y1 <= b.y2) || (a.y2 >= b.y1 && a.y2 <= b.y2) {
			if (b.z1 >= a.z1 && b.z1 <= a.z2) || (b.z2 >= a.z1 && b.z2 <= a.z2) || (a.z1 >= b.z1 && a.z1 <= b.z2) || (a.z2 >= b.z1 && a.z2 <= b.z2) {
				intersection := Step{
					x1: max(a.x1, b.x1),
					x2: min(a.x2, b.x2),
					y1: max(a.y1, b.y1),
					y2: min(a.y2, b.y2),
					z1: max(a.z1, b.z1),
					z2: min(a.z2, b.z2),
					on: !b.on,
				}
				return true, intersection
			}
		}
	}
	return false, Step{}
}

func count_on(steps []Step) int {
	cores := []Step{}

	for _, step := range steps {
		to_add := []Step{}
		if step.on {
			to_add = []Step{step}
		}
		for _, core := range cores {
			ok, intersection := intersect(step, core)
			if ok {
				to_add = append(to_add, intersection)
			}
		}
		cores = append(cores, to_add...)
	}

	switched_on := 0
	for _, core := range cores {
		amount := (core.x2 - core.x1 + 1) * (core.y2 - core.y1 + 1) * (core.z2 - core.z1 + 1)
		if core.on {
			switched_on += amount
		} else {
			switched_on -= amount
		}
	}

	return switched_on
}

func main() {
	steps := get_data()

	small_steps := []Step{}
	for _, step := range steps {
		if step.x1 <= 50 && step.x2 >= -50 && step.y1 <= 50 && step.y2 >= -50 && step.z1 <= 50 && step.z2 >= -50 {
			small_steps = append(small_steps, step)
		}
	}

	fmt.Println(count_on(small_steps))
	fmt.Println(count_on(steps))
}
