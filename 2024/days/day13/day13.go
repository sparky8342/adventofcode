package day13

import (
	"fmt"
	"loader"
	"regexp"
	"strconv"
	"utils"
)

type Machine struct {
	a_x    int
	a_y    int
	b_x    int
	b_y    int
	goal_x int
	goal_y int
}

func parse_line(line string) (int, int) {
	r := regexp.MustCompile(".*?(\\d+).*?(\\d+)")
	matches := r.FindStringSubmatch(line)
	n1, err := strconv.Atoi(matches[1])
	if err != nil {
		panic(err)
	}
	n2, err := strconv.Atoi(matches[2])
	if err != nil {
		panic(err)
	}
	return n1, n2
}

func parse_data(data [][]string) []Machine {
	machines := []Machine{}
	for _, group := range data {
		a_x, a_y := parse_line(group[0])
		b_x, b_y := parse_line(group[1])
		goal_x, goal_y := parse_line(group[2])
		machine := Machine{
			a_x:    a_x,
			a_y:    a_y,
			b_x:    b_x,
			b_y:    b_y,
			goal_x: goal_x,
			goal_y: goal_y,
		}
		machines = append(machines, machine)
	}
	return machines
}

func calc(machine Machine) int {
	// solve simultaneous equations
	// e.g.
	// ap * 94 + bp * 22 = 8400
	// ap * 34 + bp * 67 = 5400

	// multiply both to get the same ap value
	l := utils.Lcm(machine.a_x, machine.a_y)
	mult1 := l / machine.a_x
	mult2 := l / machine.a_y

	bp1 := machine.b_x * mult1
	bp2 := machine.b_y * mult2

	gx := machine.goal_x * mult1
	gy := machine.goal_y * mult2

	// subtract the equations, giving us a bp value and goal value
	var bp, g int
	if bp1 > bp2 {
		bp = bp1 - bp2
		g = gx - gy
	} else {
		bp = bp2 - bp1
		g = gy - gx
	}

	// divide to get the value of b
	b := g / bp

	// replace to get the value of a
	a := (machine.goal_x - b*machine.b_x) / machine.a_x

	// check a and b work for the goals
	if a*machine.a_x+b*machine.b_x == machine.goal_x &&
		a*machine.a_y+b*machine.b_y == machine.goal_y {
		return a*3 + b
	} else {
		return 0
	}
}

func tokens(machines []Machine, add int) int {
	total := 0
	for _, machine := range machines {
		machine.goal_x += add
		machine.goal_y += add
		total += calc(machine)
	}
	return total
}

func Run() {
	loader.Day = 13
	data := loader.GetStringGroups()
	machines := parse_data(data)

	part1 := tokens(machines, 0)
	part2 := tokens(machines, 10000000000000)

	fmt.Printf("%d %d\n", part1, part2)
}
